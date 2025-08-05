package datamodels

import "slices"

var (
	fieldsFroTagsRepresentedAsList []string = []string{
		"tags.name",
		"tags.color",
		"tags.created",
		"tags.created_by.id",
		"tags.created_by.username",
		"tags.is_visible_for_customer",
	}

	fieldsForSnapshotsRepresentedAsList []string = []string{
		"os",
		"fqdn",
		"domain",
		"cmdb_id",
		"os_type",
		"hostname",
		"ip_addresses",
		"mac_addresses",
		"user_cmdb_name",
	}
)

// ************* Структура SupportingStructureForTags *************
// *** используется как временный накопитель информации о тегах ***
// ****************************************************************

// NewSupportingStructureForTags формирует вспомогательный объект для обработки
// объектов типа tag находящихся в data.tags
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
func (stags *SupportingStructureForTags) GetTagTmp() *BiZoneTag {
	return &stags.tagTmp
}

// HandlerValue обрабатывает значения, добавляет поле в список listAcceptedFields
func (stags *SupportingStructureForTags) HandlerValue(fieldBranch string, a any, f func(any)) {
	//если поле повторяется то считается что это уже новый объект
	isExist := isExistFieldsRepresentedAsList(fieldBranch, fieldsFroTagsRepresentedAsList)

	if !isExist && stags.isExistFieldBranch(fieldBranch) {
		stags.listAcceptedFields = []string(nil)
		stags.tags = append(stags.tags, stags.tagTmp)
		stags.tagTmp = *NewBiZoneTag()
	}

	stags.listAcceptedFields = append(stags.listAcceptedFields, fieldBranch)

	f(a)
}

// isExistFieldBranch проверяет есть ли в списке listAcceptedFields принятое значение
func (stags *SupportingStructureForTags) isExistFieldBranch(v string) bool {
	return slices.Contains(stags.listAcceptedFields, v)
}

// ************* Структура SupportingStructureForSnapshots ***************
// ****** используется как временный накопитель информации о снимках *****
// ***********************************************************************

// NewSupportingStructureForSnapshots формирует вспомогательный объект для
// обработки объектов типа snapshot находящихся в data.snapshots
func NewSupportingStructureForSnapshots() *SupportingStructureForSnapshots {
	return &SupportingStructureForSnapshots{
		listAcceptedFields: []string(nil),
		snapshotTmp:        *NewBiZneSnapshot(),
		snapshots:          []BiZoneSnapshot(nil),
	}
}

// GetSnapshots возвращает []BiZoneSnapshots, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из ssnap.snapshotTmp в
// список ssnap.snapshots, так как snapshots автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются ssnap.snapshotTmp
func (ssnap *SupportingStructureForSnapshots) GetSnapshots() []BiZoneSnapshot {
	ssnap.listAcceptedFields = []string(nil)
	if ssnap.snapshotTmp.CMDBID != "" {
		ssnap.snapshots = append(ssnap.snapshots, ssnap.snapshotTmp)
	}

	return ssnap.snapshots
}

// GetSnapshotTmp возвращает временный объект snapshotTmp
func (ssnap *SupportingStructureForSnapshots) GetSnapshotTmp() *BiZoneSnapshot {
	return &ssnap.snapshotTmp
}

// HandlerValue обрабатывает значения, добавляет поле в список listAcceptedFields
func (ssnap *SupportingStructureForSnapshots) HandlerValue(fieldBranch string, a any, f func(any)) {
	//если поле повторяется то считается что это уже новый объект
	isExist := isExistFieldsRepresentedAsList(fieldBranch, fieldsFroTagsRepresentedAsList)

	if !isExist && ssnap.isExistFieldBranch(fieldBranch) {
		ssnap.listAcceptedFields = []string(nil)
		ssnap.snapshots = append(ssnap.snapshots, ssnap.snapshotTmp)
		ssnap.snapshotTmp = *NewBiZneSnapshot()
	}

	ssnap.listAcceptedFields = append(ssnap.listAcceptedFields, fieldBranch)

	f(a)
}

// isExistFieldBranch проверяет есть ли в списке listAcceptedFields принятое значение
func (ssnap *SupportingStructureForSnapshots) isExistFieldBranch(v string) bool {
	return slices.Contains(ssnap.listAcceptedFields, v)
}

func isExistFieldsRepresentedAsList(field string, list []string) bool {
	return slices.Contains(list, field)
}
