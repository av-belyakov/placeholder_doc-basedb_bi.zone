package documentgenerator

import "time"

// ******* Основная структура Event ********
// *****************************************

// GetSeverity для поля Severity
func (e *VerifiedBiZoneAlert) GetSeverity() string {
	return e.Severity
}

// SetSeverity для поля Severity
func (e *VerifiedBiZoneAlert) SetSeverity(severity string) {
	e.Severity = severity
}

// GetUpdatedTime для поля UpdatedTime
func (e *VerifiedBiZoneAlert) GetUpdatedTime() time.Time {
	return e.UpdatedTime
}

// SetUpdatedTime для поля UpdatedTime
func (e *VerifiedBiZoneAlert) SetUpdatedTime(updatedTime time.Time) {
	e.UpdatedTime = updatedTime
}

// GetVerifiedBiZoneAlertEndTime для поля EventEndTime
func (e *VerifiedBiZoneAlert) GetEventEndTime() time.Time {
	return e.EventEndTime
}

// SetEventEndTime для поля EventEndTime
func (e *VerifiedBiZoneAlert) SetEventEndTime(eventEndTime time.Time) {
	e.EventEndTime = eventEndTime
}

// GetPlatformHostname для поля PlatformHostname
func (e *VerifiedBiZoneAlert) GetPlatformHostname() string {
	return e.PlatformHostname
}

// SetPlatformHostname для поля PlatformHostname
func (e *VerifiedBiZoneAlert) SetPlatformHostname(platformHostname string) {
	e.PlatformHostname = platformHostname
}

// GetTitle для поля Title
func (e *VerifiedBiZoneAlert) GetTitle() string {
	return e.Title
}

// SetTitle для поля Title
func (e *VerifiedBiZoneAlert) SetTitle(title string) {
	e.Title = title
}

// GetRecommendations для поля Recommendations
func (e *VerifiedBiZoneAlert) GetRecommendations() string {
	return e.Recommendations
}

// SetRecommendations для поля Recommendations
func (e *VerifiedBiZoneAlert) SetRecommendations(recommendations string) {
	e.Recommendations = recommendations
}

// GetData для поля Data
func (e *VerifiedBiZoneAlert) GetData() Data {
	return e.Data
}

// SetData для поля Data
func (e *VerifiedBiZoneAlert) SetData(data Data) {
	e.Data = data
}

// ************* Структура Data ************
// *****************************************

// GetID для поля ID
func (d *Data) GetID() string {
	return d.ID
}

// SetID для поля ID
func (d *Data) SetID(id string) {
	d.ID = id
}

// GetTitle для поля Title
func (d *Data) GetTitle() string {
	return d.Title
}

// SetTitle для поля Title
func (d *Data) SetTitle(title string) {
	d.Title = title
}

// GetSensor для поля Sensor
func (d *Data) GetSensor() int {
	return d.Sensor
}

// SetSensor для поля Sensor
func (d *Data) SetSensor(sensor int) {
	d.Sensor = sensor
}

// GetIPHome для поля IPHome
func (d *Data) GetIPHome() string {
	return d.IPHome
}

// SetIPHome для поля IPHome
func (d *Data) SetIPHome(ipHome string) {
	d.IPHome = ipHome
}

// GetIPExter для поля IPExter
func (d *Data) GetIPExter() string {
	return d.IPExter
}

// SetIPExter для поля IPExter
func (d *Data) SetIPExter(ipExter string) {
	d.IPExter = ipExter
}

// GetSnortSid для поля SnortSid
func (d *Data) GetSnortSid() []int {
	return d.SnortSid
}

// SetSnortSid для поля SnortSid
func (d *Data) SetSnortSid(snortSid []int) {
	d.SnortSid = snortSid
}

// GetAllSensors для поля AllSensors
func (d *Data) GetAllSensors() []int {
	return d.AllSensors
}

// SetAllSensors для поля AllSensors
func (d *Data) SetAllSensors(allSensors []int) {
	d.AllSensors = allSensors
}

// GetAllSensors для поля AllIPHome
func (d *Data) GetAllIPHome() []string {
	return d.AllIPHome
}

// SetAllSensors для поля AllIPHome
func (d *Data) SetAllIPHome(allIPHome []string) {
	d.AllIPHome = allIPHome
}

// GetURLHTTP для поля URLHTTP
func (d *Data) GetURLHTTP() string {
	return d.URLHTTP
}

// SetURLHTTP для поля URLHTTP
func (d *Data) SetURLHTTP(urlHTTP string) {
	d.URLHTTP = urlHTTP
}

// GetEventType для поля EventType
func (d *Data) GetEventType() string {
	return d.EventType
}

// SetEventType для поля EventType
func (d *Data) SetEventType(eventType string) {
	d.EventType = eventType
}

// GetURLArkime для поля URLArkime
func (d *Data) GetURLArkime() string {
	return d.URLArkime
}

// SetURLArkime для поля URLArkime
func (d *Data) SetURLArkime(urlArkime string) {
	d.URLArkime = urlArkime
}

// GetURLFTP для поля URLFTP
func (d *Data) GetURLFTP() string {
	return d.URLFTP
}

// SetURLFTP для поля URLFTP
func (d *Data) SetURLFTP(urlFTP string) {
	d.URLFTP = urlFTP
}

// GetAllIPExt для поля AllIPExt
func (d *Data) GetAllIPExt() int {
	return d.AllIPExt
}

// SetAllIPExt для поля AllIPExt
func (d *Data) SetAllIPExt(allIPExt int) {
	d.AllIPExt = allIPExt
}

// GetAffectedLogSources для поля AffectedLogSources
func (d *Data) GetAffectedLogSources() string {
	return d.AffectedLogSources
}

// SetAffectedLogSources для поля AffectedLogSources
func (d *Data) SetAffectedLogSources(affectedLogSources string) {
	d.AffectedLogSources = affectedLogSources
}

// GetConfidence для поля Confidence
func (d *Data) GetConfidence() string {
	return d.Confidence
}

// SetConfidence для поля Confidence
func (d *Data) SetConfidence(confidence string) {
	d.Confidence = confidence
}

// GetCreatedTime для поля CreatedTime
func (d *Data) GetCreatedTime() time.Time {
	return d.CreatedTime
}

// SetCreatedTime для поля CreatedTime
func (d *Data) SetCreatedTime(createdTime time.Time) {
	d.CreatedTime = createdTime
}

// GetTags для поля Tags
func (d *Data) GetTags() []Tag {
	return d.Tags
}

// SetTags для поля Tags
func (d *Data) SetTags(tags []Tag) {
	d.Tags = tags
}

// GetDescription для поля Description
func (d *Data) GetDescription() string {
	return d.Description
}

// SetDescription для поля Description
func (d *Data) SetDescription(description string) {
	d.Description = description
}

// GetEventStartTime для поля EventStartTime
func (d *Data) GetEventStartTime() time.Time {
	return d.EventStartTime
}

// SetEventStartTime для поля EventStartTime
func (d *Data) SetEventStartTime(eventStartTime time.Time) {
	d.EventStartTime = eventStartTime
}

// GetLastDetectionTime для поля LastDetectionTime
func (d *Data) GetLastDetectionTime() time.Time {
	return d.LastDetectionTime
}

// SetLastDetectionTime для поля LastDetectionTime
func (d *Data) SetLastDetectionTime(lastDetectionTime time.Time) {
	d.LastDetectionTime = lastDetectionTime
}

// GetIDNum для поля IDNum
func (d *Data) GetIDNum() int {
	return d.IDNum
}

// SetIDNum для поля IDNum
func (d *Data) SetIDNum(idNum int) {
	d.IDNum = idNum
}

// GetDetectionRule для поля DetectionRule
func (d *Data) GetDetectionRule() string {
	return d.DetectionRule
}

// SetDetectionRule для поля DetectionRule
func (d *Data) SetDetectionRule(detectionRule string) {
	d.DetectionRule = detectionRule
}

// GetSnapshots для поля Snapshots
func (d *Data) GetSnapshots() []Snapshot {
	return d.Snapshots
}

// SetSnapshots для поля Snapshots
func (d *Data) SetSnapshots(snapshots []Snapshot) {
	d.Snapshots = snapshots
}

// GetUUID для поля UUID
func (d *Data) GetUUID() string {
	return d.UUID
}

// SetUUID для поля UUID
func (d *Data) SetUUID(uuid string) {
	d.UUID = uuid
}

// GetResponseTeam для поля ResponseTeam
func (d *Data) GetResponseTeam() int {
	return d.ResponseTeam
}

// SetResponseTeam для поля ResponseTeam
func (d *Data) SetResponseTeam(responseTeam int) {
	d.ResponseTeam = responseTeam
}

// GetCustomerSystem для поля CustomerSystem
func (d *Data) GetCustomerSystem() string {
	return d.CustomerSystem
}

// SetCustomerSystem для поля CustomerSystem
func (d *Data) SetCustomerSystem(customerSystem string) {
	d.CustomerSystem = customerSystem
}

// GetExternalID для поля ExternalID
func (d *Data) GetExternalID() string {
	return d.ExternalID
}

// SetExternalID для поля ExternalID
func (d *Data) SetExternalID(externalID string) {
	d.ExternalID = externalID
}

// GetFirstDetectionTime для поля FirstDetectionTime
func (d *Data) GetFirstDetectionTime() time.Time {
	return d.FirstDetectionTime
}

// SetFirstDetectionTime для поля FirstDetectionTime
func (d *Data) SetFirstDetectionTime(firstDetectionTime time.Time) {
	d.FirstDetectionTime = firstDetectionTime
}

// GetPlatformType для поля PlatformType
func (d *Data) GetPlatformType() string {
	return d.PlatformType
}

// SetPlatformType для поля PlatformType
func (d *Data) SetPlatformType(platformType string) {
	d.PlatformType = platformType
}

// ************* Структура Tag *************
// *****************************************

// GetCreated для поля Created
func (t *Tag) GetCreated() time.Time {
	return t.Created
}

// SetCreated для поля Created
func (t *Tag) Set(created time.Time) {
	t.Created = created
}

// GetName для поля Name
func (t *Tag) GetName() string {
	return t.Name
}

// SetName для поля Name
func (t *Tag) SetName(name string) {
	t.Name = name
}

// GetColor для поля Color
func (t *Tag) GetColor() string {
	return t.Color
}

// SetColor для поля Color
func (t *Tag) SetColor(color string) {
	t.Color = color
}

// GetCreatedByUsername для поля CreatedBy.Username
func (t *Tag) GetCreatedByUsername() string {
	return t.CreatedBy.Username
}

// SetCreatedByUsername для поля CreatedBy.Username
func (t *Tag) SetCreatedByUsername(username string) {
	t.CreatedBy.Username = username
}

// GetCreatedByID для поля CreatedBy.ID
func (t *Tag) GetCreatedByID() int {
	return t.CreatedBy.ID
}

// SetCreatedByID для поля CreatedBy.ID
func (t *Tag) SetCreatedByID(id int) {
	t.CreatedBy.ID = id
}

// ********** Структура Snapshot ***********
// *****************************************

// GetIPAddresses для поля IPAddresses
func (s *Snapshot) GetIPAddresses() []string {
	return s.IPAddresses
}

// SetIPAddresses для поля IPAddresses
func (s *Snapshot) Set(ipAddress []string) {
	s.IPAddresses = ipAddress
}

// GetMACAddresses для поля MACAddresses
func (s *Snapshot) GetMACAddresses() []string {
	return s.IPAddresses
}

// SetMACAddresses для поля MACAddresses
func (s *Snapshot) SetMACAddresses(macAddress []string) {
	s.MACAddresses = macAddress
}

// GetDomain для поля Domain
func (s *Snapshot) GetDomain() string {
	return s.Domain
}

// SetDomain для поля Domain
func (s *Snapshot) SetDomain(domain string) {
	s.Domain = domain
}

// GetFqdn для поля Fqdn
func (s *Snapshot) GetFqdn() string {
	return s.Fqdn
}

// SetFqdn для поля Fqdn
func (s *Snapshot) SetFqdn(fqdn string) {
	s.Fqdn = fqdn
}

// GetOS для поля OS
func (s *Snapshot) GetOS() string {
	return s.OS
}

// SetOS для поля OS
func (s *Snapshot) SetOS(value string) {
	s.OS = value
}

// GetUserCMDBName для поля UserCMDBName
func (s *Snapshot) GetUserCMDBName() string {
	return s.UserCMDBName
}

// SetUserCMDBName для поля UserCMDBName
func (s *Snapshot) SetUserCMDBName(value string) {
	s.UserCMDBName = value
}

// GetCMDBID для поля CMDBID
func (s *Snapshot) GetCMDBID() string {
	return s.CMDBID
}

// SetCMDBID для поля CMDBID
func (s *Snapshot) SetCMDBID(value string) {
	s.CMDBID = value
}

// GetHostname для поля Hostname
func (s *Snapshot) GetHostname() string {
	return s.Hostname
}

// SetHostname для поля Hostname
func (s *Snapshot) SetHostname(hostname string) {
	s.Hostname = hostname
}

// GetOSType для поля OSType
func (s *Snapshot) GetOSType() string {
	return s.OSType
}

// SetOSType для поля OSType
func (s *Snapshot) SetOSType(osType string) {
	s.OSType = osType
}
