package helpers

import (
	"context"
	"fmt"

	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"
)

// *** для счётчика ***
type Counting struct {
	ch chan struct {
		Message string
		Count   int
	}
}

func NewCounting() *Counting {
	return &Counting{ch: make(chan struct {
		Message string
		Count   int
	})}
}

func (c *Counting) SendMessage(msg string, count int) {
	c.ch <- struct {
		Message string
		Count   int
	}{Message: msg, Count: count}
}

func (c *Counting) GetChan() <-chan struct {
	Message string
	Count   int
} {
	return c.ch
}

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
	Ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{
		Ch: make(chan interfaces.Messager),
	}
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.Ch
}

func (l *Logging) Send(msgType, msgData string) {
	l.Ch <- &Message{Type: msgType, Message: msgData}
}

func ConnectionInitialization(
	ctx context.Context,
	logging interfaces.Logger,
	counting interfaces.Counter,
	opts ...databasestorageapi.DatabaseStorageOptions,
) (*databasestorageapi.DatabaseStorage, error) {
	var (
		apiDBS *databasestorageapi.DatabaseStorage
		err    error
	)

	apiDBS, err = databasestorageapi.New(counting, logging, opts...)
	if err != nil {
		return apiDBS, err
	}

	if err := apiDBS.Start(ctx); err != nil {
		return apiDBS, err
	}

	go func(ctx context.Context, ch <-chan databasestorageapi.SettingsChanOutput) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-ch:
				fmt.Println("Chan output message:", msg)

			}
		}
	}(ctx, apiDBS.GetChanDataFromModule())

	return apiDBS, err
}
