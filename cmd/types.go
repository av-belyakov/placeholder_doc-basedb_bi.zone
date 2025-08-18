package main

import (
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/databasestorageapi"
	modulekafkaapi "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/kafkaapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/cmd/natsapi"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// ApplicationRouter модуль маршрутизации
type ApplicationRouter struct {
	logger         interfaces.Logger
	counter        interfaces.Counter
	chToNatsApi    chan<- natsapi.SettingsChanInput
	chFromNatsApi  <-chan natsapi.SettingsChanOutput
	chToKafkaApi   chan<- modulekafkaapi.SettingsChanInput
	chFromKafkaApi <-chan modulekafkaapi.SettingsChanOutput
	chToDBSApi     chan<- databasestorageapi.SettingsChanInput
	chFromDBSApi   <-chan databasestorageapi.SettingsChanOutput
}

// ApplicationRouterSettings настройки модуля маршрутизации
type ApplicationRouterSettings struct {
	ChanToNats    chan<- natsapi.SettingsChanInput
	ChanFromNats  <-chan natsapi.SettingsChanOutput
	ChanToKafka   chan<- modulekafkaapi.SettingsChanInput
	ChanFromKafka <-chan modulekafkaapi.SettingsChanOutput
	ChanToDBS     chan<- databasestorageapi.SettingsChanInput
	ChanFromDBS   <-chan databasestorageapi.SettingsChanOutput
}
