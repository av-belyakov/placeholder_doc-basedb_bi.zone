package documentgenerator

import (
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/cmd/handlers"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"
)

// BiZoneAlertsGenerator генерирует верифицированный объект типа 'alerts'. Вернет первым элементом основной ID,
// вторым проверенный объект типа 'alerts', третьим список полей которые не были обработаны
func BiZoneAlertsGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *datamodels.VerifiedBiZoneAlert, map[string]string) {
	var (
		externalId string //external_id

		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//verifiedTags      []*datamodels.BiZoneTag      = []*datamodels.BiZoneTag(nil)
		//verifiedSnapshots []*datamodels.BiZoneSnapshot = []*datamodels.BiZoneSnapshot(nil)
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

	}

	return externalId, verifiedMainObject, listRawFields
}
