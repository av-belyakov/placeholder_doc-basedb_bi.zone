package datamodels

// SupportingStructureForTags вспомогательный тип используемый для хранения
// объектов типа tags
type SupportingStructureForTags struct {
	listAcceptedFields []string
	tagTmp             BiZoneTag
	tags               []BiZoneTag
}

// SupportingStructureForSnapshots вспомогательный тип используемый для хранения
// объектов типа snapshots
type SupportingStructureForSnapshots struct {
	listAcceptedFields []string
	snapshotTmp        BiZoneSnapshot
	snapshots          []BiZoneSnapshot
}
