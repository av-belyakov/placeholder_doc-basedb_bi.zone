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
	}
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
func (d *BiZoneData) SetAnyID(i any) {
	d.ID = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyTitle(i any) {
	d.Title = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnySensor(i any) {
	if v, ok := i.(float32); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		d.Sensor = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		d.Sensor = v
	}
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
func (d *BiZoneData) SetAnyIPHome(i any) {
	d.IPHome = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyIPExter(i any) {
	d.IPExter = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyURLHTTP(i any) {
	d.URLHTTP = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyEventType(i any) {
	d.EventType = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyURLArkime(i any) {
	d.URLArkime = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyURLFTP(i any) {
	d.URLFTP = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyAllIPExt(i any) {
	if v, ok := i.(float32); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		d.AllIPExt = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		d.AllIPExt = v
	}
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
func (d *BiZoneData) SetAnyAffectedLogSources(i any) {
	d.AffectedLogSources = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyConfidence(i any) {
	d.Confidence = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyDescription(i any) {
	d.Description = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyCreatedTime(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
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
func (d *BiZoneData) SetAnyEventStartTime(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
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
func (d *BiZoneData) SetAnyLastDetectionTime(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
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
func (d *BiZoneData) SetAnyFirstDetectionTime(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	d.FirstDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
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
func (d *BiZoneData) SetAnyIDNum(i any) {
	if v, ok := i.(float32); ok {
		d.IDNum = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		d.IDNum = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		d.IDNum = v
	}
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
func (d *BiZoneData) SetAnyDetectionRule(i any) {
	d.DetectionRule = fmt.Sprint(i)
}

// GetUUID для поля UUID
func (d *BiZoneData) GetUUID() string {
	return d.UUID
}

// SetUUID для поля UUID
func (d *BiZoneData) SetUUID(uuid string) {
	d.UUID = uuid
}

// SetAnyUUID для поля UUID
func (d *BiZoneData) SetAnyUUID(i any) {
	d.UUID = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyResponseTeam(i any) {
	if v, ok := i.(float32); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		d.ResponseTeam = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		d.ResponseTeam = v
	}
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
func (d *BiZoneData) SetAnyCustomerSystem(i any) {
	d.CustomerSystem = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyExternalID(i any) {
	d.ExternalID = fmt.Sprint(i)
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
func (d *BiZoneData) SetAnyPlatformType(i any) {
	d.PlatformType = fmt.Sprint(i)
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

// GetSnortSid для поля SnortSid
func (d *BiZoneData) GetSnortSid() []uint64 {
	return d.SnortSid
}

// SetSnortSid для поля SnortSid
func (d *BiZoneData) SetSnortSid(snortSid []uint64) {
	d.SnortSid = snortSid
}

// GetAllSensors для поля AllSensors
func (d *BiZoneData) GetAllSensors() []uint64 {
	return d.AllSensors
}

// SetAllSensors для поля AllSensors
func (d *BiZoneData) SetAllSensors(allSensors []uint64) {
	d.AllSensors = allSensors
}

// GetAllSensors для поля AllIPHome
func (d *BiZoneData) GetAllIPHome() []string {
	return d.AllIPHome
}

// SetAllSensors для поля AllIPHome
func (d *BiZoneData) SetAllIPHome(allIPHome []string) {
	d.AllIPHome = allIPHome
}
