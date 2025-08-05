package handlers

import "github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"

// NewListBiZoneHandlerSnapshots обработчик для значений типа 'data.snapshots.*' основного объекта
func NewListBiZoneHandlerSnapshots(sst *datamodels.SupportingStructureForSnapshots) map[string][]func(any) {
	return map[string][]func(any){
		//--- os ---
		"data.snapshots.os": {func(a any) {
			sst.HandlerValue("data.snapshots.os", a, sst.GetSnapshotTmp().SetAnyOS)
		}},
		//--- fqdn ---
		"data.snapshots.fqdn": {func(a any) {
			sst.HandlerValue("data.snapshots.fqdn", a, sst.GetSnapshotTmp().SetAnyFqdn)
		}},
		//--- cmdb_id ---
		"data.snapshots.cmdb_id": {func(a any) {
			sst.HandlerValue("data.snapshots.cmdb_id", a, sst.GetSnapshotTmp().SetAnyCMDBID)
		}},
		//--- domain ---
		"data.snapshots.domain": {func(a any) {
			sst.HandlerValue("data.snapshots.domain", a, sst.GetSnapshotTmp().SetAnyDomain)
		}},
		//--- os_type ---
		"data.snapshots.os_type": {func(a any) {
			sst.HandlerValue("data.snapshots.os_type", a, sst.GetSnapshotTmp().SetAnyOSType)
		}},
		//--- hostname ---
		"data.snapshots.hostname": {func(a any) {
			sst.HandlerValue("data.snapshots.hostname", a, sst.GetSnapshotTmp().SetAnyHostname)
		}},
		//--- user_cmdb_name ---
		"data.snapshots.user_cmdb_name": {func(a any) {
			sst.HandlerValue("data.snapshots.user_cmdb_name", a, sst.GetSnapshotTmp().SetAnyUserCMDBName)
		}},
		//--- ip_addresses ---
		"data.snapshots.ip_addresses": {func(a any) {
			sst.HandlerValue("data.snapshots.ip_addresses", a, sst.GetSnapshotTmp().SetAnyIPAddresse)
		}},
		//--- mac_addresses ---
		"data.snapshots.mac_addresses": {func(a any) {
			sst.HandlerValue("data.snapshots.mac_addresses", a, sst.GetSnapshotTmp().SetAnyMACAddresse)
		}},
	}
}
