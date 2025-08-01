package datamodels

import "slices"

var fieldsRepresentedAsList []string = []string{
	"tags.name",
	"tags.color",
	"tags.created",
	"tags.is_visible_for_customer",
	"tags.created_by.id",
	"tags.created_by.username",
}

// ************* Структура SupportingStructureForTags *************
// ****************************************************************

// NewSupportingStructureForTags формирует вспомогательный объект для обработки
// объектов типа data.tag
func NewSupportingStructureForTags() *SupportingStructureForTags {
	return &SupportingStructureForTags{
		listAcceptedFields: []string(nil),
		tagTmp:             *NewBiZoneTag(),
		tags:               []BiZoneTag(nil),
	}
}

// GetTags возвращает []BiZoneTag, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из stags.tagTmp в
// список stag.tags, так как ttps автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются stags.tagTmp
func (stags *SupportingStructureForTags) GetTags() []BiZoneTag {
	stags.listAcceptedFields = []string(nil)
	if stags.tagTmp.Name != "" {
		stags.tags = append(stags.tags, stags.tagTmp)
	}

	return stags.tags
}

// GetTagTmp возвращает временный объект tagTmp
func (stags *SupportingStructureForTags) GetTtpTmp() *BiZoneTag {
	return &stags.tagTmp
}

func (stags *SupportingStructureForTags) HandlerValue(fieldBranch string, i any, f func(any)) {
	//если поле повторяется то считается что это уже новый объект
	isExist := isExistFieldsRepresentedAsList(fieldBranch, fieldsRepresentedAsList)

	if !isExist && stags.isExistFieldBranch(fieldBranch) {
		stags.listAcceptedFields = []string(nil)
		stags.tags = append(stags.tags, stags.tagTmp)
		stags.tagTmp = *NewBiZoneTag()
	}

	stags.listAcceptedFields = append(stags.listAcceptedFields, fieldBranch)

	f(i)
}

func (stags *SupportingStructureForTags) isExistFieldBranch(v string) bool {
	return slices.Contains(stags.listAcceptedFields, v)
}

func isExistFieldsRepresentedAsList(field string, list []string) bool {
	return slices.Contains(list, field)
}
