package databasestorageapi

import (
	"github.com/elastic/go-elasticsearch/v8"

	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"
)

type DatabaseStorage struct {
	counter  interfaces.Counter
	logger   interfaces.Logger
	client   *elasticsearch.Client
	settings settingsDatabaseStorage
	chInput  chan SettingsChanInput  //канал для передачи данных в модуль
	chOutput chan SettingsChanOutput //канал для приёма данных из модуля
}

type settingsDatabaseStorage struct {
	storages map[string]string
	namedb   string
	user     string
	passwd   string
	host     string
	port     int
}

// SettingsChanInput параметры канала для передачи данных в модуль
type SettingsChanInput struct {
	Data    any
	Section string
	Command string
}

// SettingsChanOutput параметры канала для приёма данных из модуля
type SettingsChanOutput struct {
	Data    []byte
	RootId  string
	CaseId  string
	Command string
}

type DatabaseStorageOptions func(*DatabaseStorage) error

//****** для объектов типа VerifiedBiZoneAlerts *******

// ResponseVerifiedBiZoneAlerts ответ от базы данных
type ResponseVerifiedBiZoneAlerts struct {
	Options ResponseVerifiedBiZoneAlertsOptions `json:"hits"`
}

// ResponseVerifiedBiZoneAlertsOptions опции ответа
type ResponseVerifiedBiZoneAlertsOptions struct {
	Total    OptionsTotal                 `json:"total"`
	Hits     []PatternVerifiedBiZoneAlert `json:"hits"`
	MaxScore float64                      `json:"max_score"`
}

// PatternVerifiedBiZoneAlert шаблон
type PatternVerifiedBiZoneAlert struct {
	Source datamodels.VerifiedBiZoneAlert `json:"_source"`
	ServiseOption
}

// OptionsTotal опции в результате поиска
type OptionsTotal struct {
	Relation string `json:"relation"` //отношение (==, >, <)
	Value    int    `json:"value"`    //количество найденных значений
}

// ServiseOption сервисные опции
type ServiseOption struct {
	ID    string `json:"_id"`
	Index string `json:"_index"`
}

/*
* 		!!!!!!
* всё что ниже можно удалять
* !!!!!!
 */

// ******* для объектов типа 'alert' *******
// AlertDBResponse информация о кейсах
type AlertDBResponse struct {
	Options AlertDBResponseOptions `json:"hits"`
}

// AlertDBResponseOptions информация о кейсах
type AlertDBResponseOptions struct {
	Total    OptionsTotal           `json:"total"`
	Hits     []PatternVerifiedAlert `json:"hits"`
	MaxScore float64                `json:"max_score"`
}

// PatternVerifiedAlert шаблон
type PatternVerifiedAlert struct {
	Source documentgenerator.VerifiedAlert `json:"_source"`
	ServiseOption
}

// ******* для объектов типа 'case' *******
// CaseDBResponse информация о кейсах
type CaseDBResponse struct {
	Options CaseDBResponseOptions `json:"hits"`
}

// CaseDBResponseOptions информация о кейсах
type CaseDBResponseOptions struct {
	Total    OptionsTotal          `json:"total"`
	Hits     []PatternVerifiedCase `json:"hits"`
	MaxScore float64               `json:"max_score"`
}

// PatternVerifiedCase шаблон
type PatternVerifiedCase struct {
	Source documentgenerator.VerifiedCase `json:"_source"`
	ServiseOption
}

// AdditionalInformationIpAddress дополнительная информация добавляемая к информации по кейсам
type AdditionalInformationIpAddress struct {
	IpAddresses []IpAddressesInformation `json:"@ipAddressAdditionalInformation"`
}

// AdditionalInformationSensors дополнительная информация добавляемая к информации по кейсам
type AdditionalInformationSensors struct {
	Sensors []SensorInformation `json:"@sensorAdditionalInformation"`
}

// SensorInformation содержит дополнительную информацию о сенсоре
type SensorInformation struct {
	SensorId    string `json:"sensorId" bson:"sensorId"`       //идентификатор сенсора
	HostId      string `json:"hostId" bson:"hostId"`           //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	GeoCode     string `json:"geoCode" bson:"geoCode"`         //географический код страны
	ObjectArea  string `json:"objectArea" bson:"objectArea"`   //сфера деятельности объекта
	SubjectRF   string `json:"subjectRF" bson:"subjectRF"`     //субъект Российской Федерации
	INN         string `json:"inn" bson:"inn"`                 //налоговый идентификатор
	HomeNet     string `json:"homeNet" bson:"homeNet"`         //перечень домашних сетей
	OrgName     string `json:"orgName" bson:"orgName"`         //наименование организации
	FullOrgName string `json:"fullOrgName" bson:"fullOrgName"` //полное наименование организации
}

// IpAddressesInformation дополнительная информация об ip адресе
type IpAddressesInformation struct {
	Ip          string `json:"ip"`          //ip адрес по которому выполнялся поиск
	City        string `json:"city"`        //город
	Country     string `json:"country"`     //страна
	CountryCode string `json:"countryCode"` //код страны
}

type listSensorId struct {
	sensors []string
}
