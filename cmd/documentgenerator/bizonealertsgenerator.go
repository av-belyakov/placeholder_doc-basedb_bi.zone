package documentgenerator

import (
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

		verifiedTags      []*datamodels.BiZoneTag      = []*datamodels.BiZoneTag(nil)
		verifiedSnapshots []*datamodels.BiZoneSnapshot = []*datamodels.BiZoneSnapshot(nil)
	)

	//финальный объект
	verifiedAlert := datamodels.NewVerifiedBiZoneAlert()

	verifiedData := datamodels.NewBiZoneData()

	//********* основные обработчики **********

	for msg := range chInput {

	}

	return externalId, verifiedAlert, listRawFields
}
