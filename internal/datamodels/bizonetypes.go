package datamodels

// VerifiedBiZoneAlert основная структура
type VerifiedBiZoneAlert struct {
	AdditionalInformation
	Snapshots          []BiZoneSnapshot `json:"snapshots"`
	Tags               []BiZoneTag      `json:"tags"`
	SpecialUUID        string           `json:"@special_uuid"`
	UUID               string           `json:"uuid"`
	Title              string           `json:"title"`
	Severity           string           `json:"severity"`
	ExternalID         string           `json:"external_id"`
	Confidence         string           `json:"confidence"`
	Description        string           `json:"description"`
	CreatedTime        string           `json:"created_time"` //дата создания (формат RFC3339)
	UpdatedTime        string           `json:"updated_time"` //дата обновления (формат RFC3339)
	PlatformType       string           `json:"platform_type"`
	EventEndTime       string           `json:"event_end_time"` //дата конца события (формат RFC3339)
	DetectionRule      string           `json:"detection_rule"`
	CustomerSystem     string           `json:"customer_system"`
	EventStartTime     string           `json:"event_start_time"` //дата начала события (формат RFC3339)
	Recommendations    string           `json:"recommendations"`
	PlatformHostname   string           `json:"platform_hostname"`
	LastDetectionTime  string           `json:"last_detection_time"`  //дата последнего обнаружения (формат RFC3339)
	FirstDetectionTime string           `json:"first_detection_time"` //дата первого обнаружения (формат RFC3339)
	AffectedLogSources string           `json:"affected_log_sources"`
	IDNum              uint64           `json:"id"`
	Data               BiZoneData       `json:"data"`
}

// BiZoneData структура для вложенного объекта data
type BiZoneData struct {
	AllIPHome    []string `json:"all__ip_home"`
	SnortSid     []uint64 `json:"snort_sid"`
	AllSensors   []uint64 `json:"all_sensors"`
	ID           string   `json:"_id"`
	Title        string   `json:"title"`
	IPHome       string   `json:"ip_home"`
	URLFTP       string   `json:"url____ftp"`
	IPExter      string   `json:"ip_exter"`
	URLHTTP      string   `json:"url___http"`
	EventType    string   `json:"event_type"`
	URLArkime    string   `json:"url_arkime"`
	Sensor       uint64   `json:"sensor"`
	AllIPExt     uint64   `json:"all____ip_ext"`
	ResponseTeam uint64   `json:"response_team"`
}

type verifiedTags []*BiZoneTag

// BiZoneTag структура для тегов
type BiZoneTag struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	Created   string `json:"created"` //дата создания (формат RFC3339)
	CreatedBy struct {
		Username string `json:"username"`
		ID       uint64 `json:"id"`
	} `json:"created_by"`
}

type verifiedSnapshots []*BiZoneSnapshot

// BiZoneSnapshot структура для снимков
type BiZoneSnapshot struct {
	IPAddresses  []string `json:"ip_addresses"`
	MACAddresses []string `json:"mac_addresses"`
	OS           string   `json:"os"`
	Fqdn         string   `json:"fqdn"`
	Domain       string   `json:"domain"`
	CMDBID       string   `json:"cmdb_id"`
	OSType       string   `json:"os_type"`
	Hostname     string   `json:"hostname"`
	UserCMDBName string   `json:"user_cmdb_name"`
}
