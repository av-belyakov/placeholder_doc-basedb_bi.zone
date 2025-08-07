package documentgenerator

import (
	alert "github.com/av-belyakov/objectsthehiveformat/alert"
	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
	eventalert "github.com/av-belyakov/objectsthehiveformat/eventalert"
	eventcase "github.com/av-belyakov/objectsthehiveformat/eventcase"
)

// VerifiedAlert объект представляет собой верифицированный тип 'alert'
type VerifiedAlert struct {
	Event           eventalert.TypeEventForAlert `json:"event" bson:"event"`
	Alert           alert.TypeAlert              `json:"alert,omitzero" bson:"alert"`
	ID              string                       `json:"@id" bson:"@id"`
	Source          string                       `json:"source" bson:"source"`
	CreateTimestamp string                       `json:"@timestamp" bson:"@timestamp"`
	ElasticsearchID string                       `json:"@es_id" bson:"@es_id"`
}

// VerifiedCase объект представляет собой верифицированный тип 'case'
type VerifiedCase struct {
	Event eventcase.TypeEventForCase `json:"event" bson:"event"`
	caseobservables.Observables
	casettps.Ttps
	AdditionalInformation
	ID              string `json:"@id" bson:"@id"`
	ElasticsearchID string `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string `json:"@timestamp" bson:"@timestamp"`
	Source          string `json:"source" bson:"source"`
}

type listSensorId struct {
	sensors []string
}

type listIpAddresses struct {
	ip []string
}
