package documentgenerator

import "github.com/av-belyakov/placeholder-doc-basedb-bi-zone/interfaces"

// BiZoneAlertsGenerator генерирует верифицированный объект типа 'alerts'. Вернет первым элементом основной ID,
// вторым проверенный объект типа 'alerts', третьим список полей которые не были обработаны
func BiZoneAlertsGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *VerifiedBiZoneAlert, map[string]string) {
	var (
		verifiedBiZoneAlert *VerifiedBiZoneAlert
	)
}
