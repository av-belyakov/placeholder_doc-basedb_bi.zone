package main

import (
	"github.com/isems-development/placeholder_doc-basedb_bi.zone/cmd/databasestorageapi"
	"github.com/isems-development/placeholder_doc-basedb_bi.zone/cmd/kafkaapi"
	"github.com/isems-development/placeholder_doc-basedb_bi.zone/cmd/natsapi"
	"github.com/isems-development/placeholder_doc-basedb_bi.zone/interfaces"
)

// ApplicationRouter модуль маршрутизации
type ApplicationRouter struct {
	logger         interfaces.Logger
	counter        interfaces.Counter
	chToNatsApi    chan<- natsapi.SettingsChanInput
	chFromNatsApi  <-chan natsapi.SettingsChanOutput
	chToKafkaApi   chan<- kafkaapi.SettingsChanInput
	chFromKafkaApi <-chan kafkaapi.SettingsChanOutput
	chToDBSApi     chan<- databasestorageapi.SettingsChanInput
	chFromDBSApi   <-chan databasestorageapi.SettingsChanOutput
}

// ApplicationRouterSettings настройки модуля маршрутизации
type ApplicationRouterSettings struct {
	ChanToNats    chan<- natsapi.SettingsChanInput
	ChanFromNats  <-chan natsapi.SettingsChanOutput
	ChanToKafka   chan<- kafkaapi.SettingsChanInput
	ChanFromKafka <-chan kafkaapi.SettingsChanOutput
	ChanToDBS     chan<- databasestorageapi.SettingsChanInput
	ChanFromDBS   <-chan databasestorageapi.SettingsChanOutput
}
