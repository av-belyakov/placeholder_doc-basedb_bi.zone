package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/supportingfunctions"
)

// addBiZoneAlerts добавление объекта типа 'alerts'
func (dbs *DatabaseStorage) addBiZoneAlerts(ctx context.Context, data any) {
	t := time.Now()

	newDocument, ok := data.(*datamodels.VerifiedBiZoneAlert)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	newDocumentBinary, err := json.Marshal(newDocument.Get())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["alerts"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//получаем существующие индексы
	existingIndexes, err := dbs.GetExistingIndexes(ctx, indexName)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	defer func(document *datamodels.VerifiedBiZoneAlert, getChan func() chan SettingsChanOutput, logger interfaces.Logger) {
		rootId := newDocument.GetUUID()

		//обогащение кейса дополнительной информацией о локальном место положении ip адресов
		listIp := documentgenerator.GetListIPAddr(document.GetAdditionalInformation().GetIpAddressesInformation())
		ok, err := sendGeoIpRequest(rootId, listIp, getChan)
		if err != nil {
			logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
		if ok {
			logger.Send("info", fmt.Sprintf("we are sending a request to search for information on the following list of ip addresses: %#v", listIp))
		}

		//обогащение кейса дополнительной информацией по расположению сенсоров
		listSensor := documentgenerator.GetListSensorId(newDocument.GetAdditionalInformation().GetSensorsInformation())
		ok, err = sendSensorInformationRequest(rootId, listSensor, getChan)
		if err != nil {
			logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
		if ok {
			logger.Send("info", fmt.Sprintf("we are sending a request to search for information on the following list of sensors: %#v", listSensor))
		}
	}(newDocument, dbs.GetChanDataFromModule, dbs.logger)

	//формируем наименование индекса
	currentIndex := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), int(t.Month()))

	//*****************************
	//*** здесь установка тегов ***
	//пока не знаю нужно ли это в BiZone
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//caseId := fmt.Sprint(newDocument.GetEvent().GetObject().CaseId)
	//reqSetTag := fmt.Appendf(nil, `{
	//					  "service": "placeholder_doc-basedb-bi-zone",
	//					  "command": "add_case_tag",
	//					  "root_id": "%s",
	//					  "case_id": "%s",
	//					  "value": "Webhook: send=\"ElasticsearchDB"
	//					}`,
	//	newDocument.GetEvent().GetRootId(),
	//	caseId)
	//*****************************

	//будет выполнятся поиск по индексам только в текущем году так как при
	//накоплении большого количества индексов, поиск по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range existingIndexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}

	// если похожих индексов нет
	if len(indexesOnlyCurrentYear) == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		existingIndexes = append(existingIndexes, currentIndex)
		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject alerts to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document to alerts id:'%s', uuid:'%s', status code:'%d'", newDocument.GetIDNum(), newDocument.GetUUID(), statusCode))

		//*****************************
		//*** здесь установка тегов ***
		//пока не знаю нужно ли это в BiZone
		//запрос на отправку сообщения для установки тега
		//dbs.GetChanDataFromModule() <- SettingsChanOutput{
		//	Command: "set_tag",
		//	CaseId:  caseId,
		//	RootId:  newDocument.GetEvent().GetRootId(),
		//	Data:    reqSetTag,
		//}
		//*****************************

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	currentQuery := strings.NewReader(
		fmt.Sprintf(
			"{\"query\": {\"bool\": {\"must\": [{\"match\": {\"uuid\": \"%s\"}}]}}}",
			newDocument.GetUUID()))
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(context.Background()),
		dbs.client.Search.WithIndex(indexesOnlyCurrentYear...),
		dbs.client.Search.WithBody(currentQuery),
	)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}
	defer res.Body.Close()

	response := CaseDBResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//вставка выполняется только когда не найден искомый документ
	if response.Options.Total.Value == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject alerts to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document to alerts id:'%s', uuid:'%s', status code:'%d'", newDocument.GetIDNum(), newDocument.GetUUID(), statusCode))

		//*****************************
		//*** здесь установка тегов ***
		//пока не знаю нужно ли это в BiZone
		//запрос на отправку сообщения для установки тега
		//dbs.GetChanDataFromModule() <- SettingsChanOutput{
		//	Command: "set_tag",
		//	CaseId:  caseId,
		//	RootId:  newDocument.GetEvent().GetRootId(),
		//	Data:    reqSetTag,
		//}
		//*****************************

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	var countReplacingFields int
	listDeleting := []ServiseOption(nil)
	updateVerified := documentgenerator.NewVerifiedCase()
	for _, v := range response.Options.Hits {
		count, err := updateVerified.Event.ReplacingOldValues(*v.Source.GetEvent())
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		} else {
			countReplacingFields += count
		}

		countReplacingFields += updateVerified.GetObservables().ReplacingOldValues(v.Source.Observables)
		countReplacingFields += updateVerified.GetTtps().ReplacingOldValues(v.Source.Ttps)

		listDeleting = append(listDeleting, ServiseOption{
			ID:    v.ID,
			Index: v.Index,
		})

		//устанавливаем время создания первой записи о кейсе
		updateVerified.SetCreateTimestamp(v.Source.CreateTimestamp)

		//заполняем новый объект информацией о сенсорах и геопринадлежности ip адресов
		updateVerified.SetAdditionalInformation(*v.Source.GetAdditionalInformation())
	}

	//выполняем обновление объекта типа Event
	updateVerified.SetSource(newDocument.GetSource())
	num, err := updateVerified.Event.ReplacingOldValues(*newDocument.GetEvent())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	} else {
		countReplacingFields += num
	}

	countReplacingFields += updateVerified.GetObservables().ReplacingOldValues(*newDocument.GetObservables())
	countReplacingFields += updateVerified.GetTtps().ReplacingOldValues(*newDocument.GetTtps())

	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//обновление уже существующего документа
	statusCode, countDel, err := dbs.UpdateDocument(ctx, currentIndex, listDeleting, nvbyte)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("uuid '%s' '%s'", newDocument.GetUUID(), err.Error())).Error())

		return
	}

	dbs.counter.SendMessage("update count insert subject alerts to db", 1)
	dbs.logger.Send("info", fmt.Sprintf("update document 'alerts' type, count delete:'%d', count replacing fields:'%d' for alerts with uuid:'%s', status code:'%d'", countDel, countReplacingFields, newDocument.GetUUID(), statusCode))
}
