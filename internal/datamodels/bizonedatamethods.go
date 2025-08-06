package datamodels

import (
	"fmt"
	"time"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ************* Структура Data ************
// *****************************************

// NewBiZoneData новый объект
func NewBiZoneData() *BiZoneData {
	return &BiZoneData{
		CreatedTime:        time.Now().Format(time.RFC3339),
		EventStartTime:     time.Now().Format(time.RFC3339),
		LastDetectionTime:  time.Now().Format(time.RFC3339),
		FirstDetectionTime: time.Now().Format(time.RFC3339),
		SnortSid:           []uint64(nil),
		AllIPHome:          []string(nil),
		AllSensors:         []uint64(nil),
		Tags:               []BiZoneTag(nil),
		Snapshots:          []BiZoneSnapshot(nil),
	}
}

func (d *BiZoneData) Get() *BiZoneData {
	return d
}

// GetID для поля ID
func (d *BiZoneData) GetID() string {
	return d.ID
}

// SetID для поля ID
func (d *BiZoneData) SetID(id string) {
	d.ID = id
}

// SetAnyID для поля ID
func (d *BiZoneData) SetAnyID(a any) {
	d.ID = fmt.Sprint(a)
}

// GetTitle для поля Title
func (d *BiZoneData) GetTitle() string {
	return d.Title
}

// SetTitle для поля Title
func (d *BiZoneData) SetTitle(title string) {
	d.Title = title
}

// SetAnyTitle для поля Title
func (d *BiZoneData) SetAnyTitle(a any) {
	d.Title = fmt.Sprint(a)
}

// GetCustomerSystem для поля CustomerSystem
func (d *BiZoneData) GetCustomerSystem() string {
	return d.CustomerSystem
}

// SetCustomerSystem для поля CustomerSystem
func (d *BiZoneData) SetCustomerSystem(customerSystem string) {
	d.CustomerSystem = customerSystem
}

// SetAnyCustomerSystem для поля CustomerSystem
func (d *BiZoneData) SetAnyCustomerSystem(a any) {
	d.CustomerSystem = fmt.Sprint(a)
}

// GetExternalID для поля ExternalID
func (d *BiZoneData) GetExternalID() string {
	return d.ExternalID
}

// SetExternalID для поля ExternalID
func (d *BiZoneData) SetExternalID(externalID string) {
	d.ExternalID = externalID
}

// SetAnyExternalID для поля ExternalID
func (d *BiZoneData) SetAnyExternalID(a any) {
	d.ExternalID = fmt.Sprint(a)
}

// GetPlatformType для поля PlatformType
func (d *BiZoneData) GetPlatformType() string {
	return d.PlatformType
}

// SetPlatformType для поля PlatformType
func (d *BiZoneData) SetPlatformType(platformType string) {
	d.PlatformType = platformType
}

// SetAnyPlatformType для поля PlatformType
func (d *BiZoneData) SetAnyPlatformType(a any) {
	d.PlatformType = fmt.Sprint(a)
}

// GetIPHome для поля IPHome
func (d *BiZoneData) GetIPHome() string {
	return d.IPHome
}

// SetIPHome для поля IPHome
func (d *BiZoneData) SetIPHome(ipHome string) {
	d.IPHome = ipHome
}

// SetAnyIPHome для поля IPHome
func (d *BiZoneData) SetAnyIPHome(a any) {
	d.IPHome = fmt.Sprint(a)
}

// GetIPExter для поля IPExter
func (d *BiZoneData) GetIPExter() string {
	return d.IPExter
}

// SetIPExter для поля IPExter
func (d *BiZoneData) SetIPExter(ipExter string) {
	d.IPExter = ipExter
}

// SetAnyIPExter для поля IPExter
func (d *BiZoneData) SetAnyIPExter(a any) {
	d.IPExter = fmt.Sprint(a)
}

// GetURLHTTP для поля URLHTTP
func (d *BiZoneData) GetURLHTTP() string {
	return d.URLHTTP
}

// SetURLHTTP для поля URLHTTP
func (d *BiZoneData) SetURLHTTP(urlHTTP string) {
	d.URLHTTP = urlHTTP
}

// SetAnyURLHTTP для поля URLHTTP
func (d *BiZoneData) SetAnyURLHTTP(a any) {
	d.URLHTTP = fmt.Sprint(a)
}

// GetEventType для поля EventType
func (d *BiZoneData) GetEventType() string {
	return d.EventType
}

// SetEventType для поля EventType
func (d *BiZoneData) SetEventType(eventType string) {
	d.EventType = eventType
}

// SetAnyEventType для поля EventType
func (d *BiZoneData) SetAnyEventType(a any) {
	d.EventType = fmt.Sprint(a)
}

// GetURLArkime для поля URLArkime
func (d *BiZoneData) GetURLArkime() string {
	return d.URLArkime
}

// SetURLArkime для поля URLArkime
func (d *BiZoneData) SetURLArkime(urlArkime string) {
	d.URLArkime = urlArkime
}

// SetAnyURLArkime для поля URLArkime
func (d *BiZoneData) SetAnyURLArkime(a any) {
	d.URLArkime = fmt.Sprint(a)
}

// GetURLFTP для поля URLFTP
func (d *BiZoneData) GetURLFTP() string {
	return d.URLFTP
}

// SetURLFTP для поля URLFTP
func (d *BiZoneData) SetURLFTP(urlFTP string) {
	d.URLFTP = urlFTP
}

// SetAnyURLFTP для поля URLFTP
func (d *BiZoneData) SetAnyURLFTP(a any) {
	d.URLFTP = fmt.Sprint(a)
}

// GetAffectedLogSources для поля AffectedLogSources
func (d *BiZoneData) GetAffectedLogSources() string {
	return d.AffectedLogSources
}

// SetAffectedLogSources для поля AffectedLogSources
func (d *BiZoneData) SetAffectedLogSources(affectedLogSources string) {
	d.AffectedLogSources = affectedLogSources
}

// SetAnyAffectedLogSources для поля AffectedLogSources
func (d *BiZoneData) SetAnyAffectedLogSources(a any) {
	d.AffectedLogSources = fmt.Sprint(a)
}

// GetConfidence для поля Confidence
func (d *BiZoneData) GetConfidence() string {
	return d.Confidence
}

// SetConfidence для поля Confidence
func (d *BiZoneData) SetConfidence(confidence string) {
	d.Confidence = confidence
}

// SetAnyConfidence для поля Confidence
func (d *BiZoneData) SetAnyConfidence(a any) {
	d.Confidence = fmt.Sprint(a)
}

// GetDescription для поля Description
func (d *BiZoneData) GetDescription() string {
	return d.Description
}

// SetDescription для поля Description
func (d *BiZoneData) SetDescription(description string) {
	d.Description = description
}

// SetAnyDescription для поля Description
func (d *BiZoneData) SetAnyDescription(a any) {
	d.Description = fmt.Sprint(a)
}

// GetUUID для поля UUID
func (d *BiZoneData) GetUUID() string {
	return d.UUID
}

// SetUUID для поля UUID
func (d *BiZoneData) SetUUID(UUID string) {
	d.UUID = UUID
}

// SetAnyUUID для поля UUID
func (d *BiZoneData) SetAnyUUID(a any) {
	d.UUID = fmt.Sprint(a)
}

// GetDetectionRule для поля DetectionRule
func (d *BiZoneData) GetDetectionRule() string {
	return d.DetectionRule
}

// SetDetectionRule для поля DetectionRule
func (d *BiZoneData) SetDetectionRule(detectionRule string) {
	d.DetectionRule = detectionRule
}

// SetAnyDetectionRule для поля DetectionRule
func (d *BiZoneData) SetAnyDetectionRule(a any) {
	d.DetectionRule = fmt.Sprint(a)
}

// GetCreatedTime для поля CreatedTime (формат RFC3339)
func (d *BiZoneData) GetCreatedTime() string {
	return d.CreatedTime
}

// SetCreatedTime для поля CreatedTime (формат RFC3339)
func (d *BiZoneData) SetCreatedTime(createdTime string) {
	d.CreatedTime = createdTime
}

// SetAnyCreatedTime для поля CreatedTime (формат RFC3339)
func (d *BiZoneData) SetAnyCreatedTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	d.CreatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetEventStartTime для поля EventStartTime (формат RFC3339)
func (d *BiZoneData) GetEventStartTime() string {
	return d.EventStartTime
}

// SetEventStartTime для поля EventStartTime (формат RFC3339)
func (d *BiZoneData) SetEventStartTime(eventStartTime string) {
	d.EventStartTime = eventStartTime
}

// SetAnyEventStartTime для поля EventStartTime (формат RFC3339)
func (d *BiZoneData) SetAnyEventStartTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	d.EventStartTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (d *BiZoneData) GetLastDetectionTime() string {
	return d.LastDetectionTime
}

// SetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (d *BiZoneData) SetLastDetectionTime(lastDetectionTime string) {
	d.LastDetectionTime = lastDetectionTime
}

// SetAnyLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (d *BiZoneData) SetAnyLastDetectionTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	d.LastDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (d *BiZoneData) GetFirstDetectionTime() string {
	return d.FirstDetectionTime
}

// SetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (d *BiZoneData) SetFirstDetectionTime(firstDetectionTime string) {
	d.FirstDetectionTime = firstDetectionTime
}

// SetAnyFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (d *BiZoneData) SetAnyFirstDetectionTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	d.FirstDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetSensor для поля Sensor
func (d *BiZoneData) GetSensor() uint64 {
	return d.Sensor
}

// SetSensor для поля Sensor
func (d *BiZoneData) SetSensor(sensor uint64) {
	d.Sensor = sensor
}

// SetAnySensor для поля Sensor
func (d *BiZoneData) SetAnySensor(a any) {
	if v, ok := a.(int); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.Sensor = v
	}
}

// GetAllIPExt для поля AllIPExt
func (d *BiZoneData) GetAllIPExt() uint64 {
	return d.AllIPExt
}

// SetAllIPExt для поля AllIPExt
func (d *BiZoneData) SetAllIPExt(allIPExt uint64) {
	d.AllIPExt = allIPExt
}

// SetAnyAllIPExt для поля AllIPExt
func (d *BiZoneData) SetAnyAllIPExt(a any) {
	if v, ok := a.(int); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.AllIPExt = v
	}
}

// GetIDNum для поля IDNum
func (d *BiZoneData) GetIDNum() uint64 {
	return d.IDNum
}

// SetIDNum для поля IDNum
func (d *BiZoneData) SetIDNum(idNum uint64) {
	d.IDNum = idNum
}

// SetAnyIDNum для поля IDNum
func (d *BiZoneData) SetAnyIDNum(a any) {
	if v, ok := a.(int); ok {
		d.IDNum = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.IDNum = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.IDNum = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.IDNum = v
	}
}

// GetResponseTeam для поля ResponseTeam
func (d *BiZoneData) GetResponseTeam() uint64 {
	return d.ResponseTeam
}

// SetResponseTeam для поля ResponseTeam
func (d *BiZoneData) SetResponseTeam(responseTeam uint64) {
	d.ResponseTeam = responseTeam
}

// SetAnyResponseTeam для поля ResponseTeam
func (d *BiZoneData) SetAnyResponseTeam(a any) {
	if v, ok := a.(int); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		d.ResponseTeam = v
	}
}

// GetSnapshots для поля Snapshots
func (d *BiZoneData) GetSnapshots() []BiZoneSnapshot {
	return d.Snapshots
}

// SetSnapshots для поля Snapshots
func (d *BiZoneData) SetSnapshots(snapshots []BiZoneSnapshot) {
	d.Snapshots = snapshots
}

// GetTags для поля Tags
func (d *BiZoneData) GetTags() []BiZoneTag {
	return d.Tags
}

// SetTags для поля Tags
func (d *BiZoneData) SetTags(tags []BiZoneTag) {
	d.Tags = tags
}

// GetSnortSids список sid snort
func (d *BiZoneData) GetSnortSids() []uint64 {
	return d.SnortSid
}

// SetSnortSid список sid snort
func (d *BiZoneData) SetSnortSids(sids []uint64) {
	d.SnortSid = sids
}

// SetSnortSid записывает значение в список SnortSid
func (d *BiZoneData) SetSnortSid(sid uint64) {
	if d.SnortSid == nil {
		d.SnortSid = []uint64(nil)
	}

	d.SnortSid = append(d.SnortSid, sid)
}

// SetAnySnortSid записывает значение в список SnortSid
func (d *BiZoneData) SetAnySnortSid(a any) {
	var sid uint64
	if d.SnortSid == nil {
		d.SnortSid = []uint64(nil)
	}

	if v, ok := a.(int); ok {
		sid = uint64(v)
	}

	if v, ok := a.(float32); ok {
		sid = uint64(v)
	}

	if v, ok := a.(float64); ok {
		sid = uint64(v)
	}

	if v, ok := a.(uint64); ok {
		sid = v
	}

	d.SnortSid = append(d.SnortSid, sid)
}

// GetAllSensors для поля AllSensors
func (d *BiZoneData) GetAllSensors() []uint64 {
	return d.AllSensors
}

// SetAllSensors для поля AllSensors
func (d *BiZoneData) SetAllSensors(allSensors []uint64) {
	d.AllSensors = allSensors
}

// SetAllSensor записывает значение в список AllSensors
func (d *BiZoneData) SetAllSensor(sensor uint64) {
	if d.AllSensors == nil {
		d.AllSensors = []uint64(nil)
	}

	d.SnortSid = append(d.SnortSid, sensor)
}

// SetAnyAllSensor записывает значение в список AllSensors
func (d *BiZoneData) SetAnyAllSensor(a any) {
	var sensor uint64
	if d.AllSensors == nil {
		d.AllSensors = []uint64(nil)
	}

	if v, ok := a.(int); ok {
		sensor = uint64(v)
	}

	if v, ok := a.(float32); ok {
		sensor = uint64(v)
	}

	if v, ok := a.(float64); ok {
		sensor = uint64(v)
	}

	if v, ok := a.(uint64); ok {
		sensor = v
	}

	d.AllSensors = append(d.AllSensors, sensor)
}

// GetAllIPHomes для поля AllIPHome
func (d *BiZoneData) GetAllIPHomes() []string {
	return d.AllIPHome
}

// SetAllIPHomes для поля AllIPHomes
func (d *BiZoneData) SetAllIPHomes(allIPHome []string) {
	d.AllIPHome = allIPHome
}

// SetAllIPHome добавляет значение AllIPHome в список
func (d *BiZoneData) SetAllIPHome(ipHome string) {
	if d.AllIPHome == nil {
		d.AllIPHome = []string(nil)
	}

	d.AllIPHome = append(d.AllIPHome, ipHome)
}

// SetAnyAllIPHome добавляет значение AllIPHome  в список
func (d *BiZoneData) SetAnyAllIPHome(i any) {
	if d.AllIPHome == nil {
		d.AllIPHome = []string(nil)
	}

	d.AllIPHome = append(d.AllIPHome, fmt.Sprint(i))
}
