package supportingfunctions

import "maps"

// GetListValues возвращает список значений из карты
func GetListValues[K comparable, V any](incomingMap map[K]V) []V {
	var list []V

	incomingValues := maps.Values(incomingMap)

	for v := range incomingValues {
		list = append(list, v)
	}

	return list
}

// SearchValue поиск значения в карте
func SearchValue[K, V comparable](incomingMap map[K]V, targetValue V) (K, bool) {
	for k, v := range incomingMap {
		if v == targetValue {
			return k, true
		}
	}

	var zero K

	return zero, false
}
