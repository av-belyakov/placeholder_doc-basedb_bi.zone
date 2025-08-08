package datamodels

// ReplacingOldBiZoneTags заменяет или добавляет значения свойства VerifiedBiZoneAlert.Tags
func (va *VerifiedBiZoneAlert) ReplacingOldBiZoneTags(incomingTypes []BiZoneTag) int {
	var countReplacingFields int

	for _, incomingType := range incomingTypes {
		index, isExist := compareTag(incomingType, va.Tags)
		if isExist {
			countReplacingFields += va.Tags[index].ReplacingOldBiZoneTag(incomingType)

			continue
		}

		va.Tags = append(va.Tags, incomingType)
		countReplacingFields++
	}

	return countReplacingFields
}

func compareTag(currentValue BiZoneTag, incomingValue []BiZoneTag) (int, bool) {
	for k, v := range incomingValue {
		if v.Name == currentValue.Name {
			return k, true
		}
	}

	return -1, false
}
