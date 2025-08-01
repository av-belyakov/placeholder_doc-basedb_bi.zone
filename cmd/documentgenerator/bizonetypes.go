package documentgenerator

import "time"

// VerifiedBiZoneAlert основная структура
type VerifiedBiZoneAlert struct {
	Severity         string    `json:"severity"`
	PlatformHostname string    `json:"platform_hostname"`
	Title            string    `json:"title"`
	Recommendations  string    `json:"recommendations"`
	UpdatedTime      time.Time `json:"updated_time"`
	EventEndTime     time.Time `json:"event_end_time"`
	Data             Data      `json:"data"`
}

// Data структура для вложенного объекта data
type Data struct {
	Snapshots          []Snapshot `json:"snapshots"`
	Tags               []Tag      `json:"tags"`
	AllIPHome          []string   `json:"all__ip_home"`
	SnortSid           []int      `json:"snort_sid"`
	AllSensors         []int      `json:"all_sensors"`
	CreatedTime        time.Time  `json:"created_time"`
	EventStartTime     time.Time  `json:"event_start_time"`
	LastDetectionTime  time.Time  `json:"last_detection_time"`
	FirstDetectionTime time.Time  `json:"first_detection_time"`
	ID                 string     `json:"_id"`
	Title              string     `json:"title"`
	IPHome             string     `json:"ip_home"`
	IPExter            string     `json:"ip_exter"`
	URLHTTP            string     `json:"url___http"`
	EventType          string     `json:"event_type"`
	URLArkime          string     `json:"url_arkime"`
	URLFTP             string     `json:"url____ftp"`
	AffectedLogSources string     `json:"affected_log_sources"`
	Confidence         string     `json:"confidence"`
	Description        string     `json:"description"`
	DetectionRule      string     `json:"detection_rule"`
	UUID               string     `json:"uuid"`
	CustomerSystem     string     `json:"customer_system"`
	ExternalID         string     `json:"external_id"`
	PlatformType       string     `json:"platform_type"`
	AllIPExt           int        `json:"all____ip_ext"`
	IDNum              int        `json:"id"`
	ResponseTeam       int        `json:"response_team"`
	Sensor             int        `json:"sensor"`
}

// Tag структура для тегов
type Tag struct {
	Created   time.Time `json:"created"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedBy struct {
		Username string `json:"username"`
		ID       int    `json:"id"`
	} `json:"created_by"`
}

// Snapshot структура для снимков
type Snapshot struct {
	IPAddresses  []string `json:"ip_addresses"`
	MACAddresses []string `json:"mac_addresses"`
	Domain       string   `json:"domain"`
	Fqdn         string   `json:"fqdn"`
	OS           string   `json:"os"`
	UserCMDBName string   `json:"user_cmdb_name"`
	CMDBID       string   `json:"cmdb_id"`
	Hostname     string   `json:"hostname"`
	OSType       string   `json:"os_type"`
}
