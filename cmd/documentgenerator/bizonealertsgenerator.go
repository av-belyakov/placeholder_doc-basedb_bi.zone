package documentgenerator

import (
	"fmt"

	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/cmd/handlers"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"
)

// BiZoneAlertsGenerator генерирует верифицированный объект типа 'alerts'. Вернет первым элементом основной ID,
// вторым проверенный объект типа 'alerts', третьим список полей которые не были обработаны
func BiZoneAlertsGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *datamodels.VerifiedBiZoneAlert, map[string]string) {
	var (
		externalId string

		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)
	)

	verifiedMainObject := datamodels.NewVerifiedBiZoneAlert()
	verifiedData := datamodels.NewBiZoneData()

	//********* основные обработчики **********
	//--- alerts ---
	listHandlerAlerts := handlers.NewListBiZoneHandlerAlerts(verifiedMainObject)

	//--- data ---
	listHandlerData := handlers.NewListBiZoneHandlerData(verifiedData)

	//******** вспомогательные объекты ********
	supportObjectTags := datamodels.NewSupportingStructureForTags()
	supportObjectSnapshot := datamodels.NewSupportingStructureForSnapshots()

	// ********* обработчики для вспомогательных объектов ***********
	listHandlerTags := handlers.NewListBiZoneHandlerTags(supportObjectTags)
	listHandlerSnapshots := handlers.NewListBiZoneHandlerSnapshots(supportObjectSnapshot)

	for msg := range chInput {
		var handlerIsExist bool

		if msg.GetFieldBranch() == "data.id" {
			verifiedMainObject.SetAnyID(msg.GetValue())
		}

		if msg.GetFieldBranch() == "data.uuid" {
			verifiedMainObject.SetAnyUUID(msg.GetValue())
			externalId = verifiedMainObject.GetUUID()
		}

		//*** обработчик для объекта alerts ***
		if lf, ok := listHandlerAlerts[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}

		//*** обработчик для под объекта alerts.data ***
		if lf, ok := listHandlerData[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}

		//***** обработчики для вспомогательных объектов *****
		//****************************************************
		//объект tags
		if lf, ok := listHandlerTags[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}
		//объект snapshots
		if lf, ok := listHandlerSnapshots[msg.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(msg.GetValue())
			}

			continue
		}

		//записываем в лог-файл поля, которые не были обработаны
		if !handlerIsExist {
			listRawFields[msg.GetFieldBranch()] = fmt.Sprint(msg.GetValue())
		}
	}

	//собираем все объекты в один
	verifiedData.SetTags(supportObjectTags.GetTags())
	verifiedData.SetSnapshots(supportObjectSnapshot.GetSnapshots())
	verifiedMainObject.SetData(*verifiedData.Get())

	return externalId, verifiedMainObject, listRawFields
}
