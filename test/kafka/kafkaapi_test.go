package kafka_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/cmd/kafkaapi"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/supportingfunctions"
)

// *** для счётчика ***
type Counting struct{}

func (c *Counting) SendMessage(msg string, count int) {}

// *** для сообщений ***
type Message struct {
	Type, Message string
}

func (m *Message) GetType() string {
	return m.Type
}

func (m *Message) SetType(v string) {
	m.Type = v
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) SetMessage(v string) {
	m.Message = v
}

// *** для логирования ***
type Logging struct {
	ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{
		ch: make(chan interfaces.Messager),
	}
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.ch
}

func (l *Logging) Send(msgType, msgData string) {
	l.ch <- &Message{Type: msgType, Message: msgData}
}

func TestKafkaApi(t *testing.T) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
	}()

	logging := NewLogging()
	counting := &Counting{}

	go func(ctx context.Context, l *Logging) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging)

	//инициализация файла для записи входящих сообщений
	f, err := os.OpenFile(
		fmt.Sprintf("incoming_messages_%s.log", time.Now().Format(time.RFC3339)),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//инициализация модуля
	kafkaApiModule, err := kafkaapi.New(
		counting,
		logging,
		kafkaapi.WithNameRegionalObject("GCM-TEST"),
		kafkaapi.WithHost("192.168.9.71"),
		kafkaapi.WithPort(31792),
		kafkaapi.WithTopicsSubscription(map[string]string{
			"alerts":      "alertmgr-alerts",
			"soar-alerts": "soar-alertmgr-alerts-default-topic",
		}),
		kafkaapi.WithCacheTTL(3600),
	)
	if err != nil {
		log.Fatal(err)
	}

	//запуск модуля
	if err := kafkaApiModule.Start(ctx); err != nil {
		log.Fatal(err)
	}

	for msg := range kafkaApiModule.GetChanDataFromModule() {
		//t.Logf("Received message:%s\n\n", string(msg.Data))
		str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data)
		if err != nil {
			t.Logf("Error: %+v\n", err)

			continue
		}

		//t.Logf("Received message:\n%s\n", str)
		if _, err = fmt.Fprintf(f, "--- TOPIC - '%s' %s\n%s\n\n", msg.SubjectType, time.Now().Format(time.RFC3339), str); err != nil {
			t.Logf("Error: %+v\n", err)
		}
	}
}
