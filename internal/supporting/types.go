package supporting

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"

//вспомогательные пользовательские типы которые можно использовать в тестах

// *** для счётчика ***
type Counting struct{}

// *** для сообщений ***
type Message struct {
	Type, Message string
}

// *** для логирования ***
type Logging struct {
	ch chan interfaces.Messager
}
