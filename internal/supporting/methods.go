package supporting

import "github.com/isems-development/placeholder_doc-basedb_bi.zone/interfaces"

// *** для счётчика ***
func (c *Counting) SendMessage(msg string, count int) {}

// *** для сообщений ***
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
