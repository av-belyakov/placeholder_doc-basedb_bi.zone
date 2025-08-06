package datamodels

import (
	"fmt"
	"time"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ******* Основная структура Alert ********
// *****************************************

// NewVerifiedBiZoneAlert новый объект
func NewVerifiedBiZoneAlert() *VerifiedBiZoneAlert {
	return &VerifiedBiZoneAlert{
		UpdatedTime:  time.Now().Format(time.RFC3339),
		EventEndTime: time.Now().Format(time.RFC3339),
	}
}

func (va *VerifiedBiZoneAlert) Get() *VerifiedBiZoneAlert {
	return va
}

// GetID для дополнительного поля @id
func (va *VerifiedBiZoneAlert) GetID() string {
	return va.ID
}

// SetID для дополнительного поля @id
func (va *VerifiedBiZoneAlert) SetID(id string) {
	va.ID = id
}

// SetAnyID для дополнительного поля @id
func (va *VerifiedBiZoneAlert) SetAnyID(a any) {
	va.ID = fmt.Sprint(a)
}

// GetUUID для дополнительного поля @uuid
func (va *VerifiedBiZoneAlert) GetUUID() string {
	return va.UUID
}

// SetUUID для дополнительного поля @uuid
func (va *VerifiedBiZoneAlert) SetUUID(uuid string) {
	va.UUID = uuid
}

// SetAnyUUID для дополнительного поля @uuid
func (va *VerifiedBiZoneAlert) SetAnyUUID(a any) {
	va.UUID = fmt.Sprint(a)
}

// GetSeverity для поля Severity
func (va *VerifiedBiZoneAlert) GetSeverity() string {
	return va.Severity
}

// SetSeverity для поля Severity
func (va *VerifiedBiZoneAlert) SetSeverity(severity string) {
	va.Severity = severity
}

// SetAnySeverity для поля Severity
func (va *VerifiedBiZoneAlert) SetAnySeverity(a any) {
	va.Severity = fmt.Sprint(a)
}

// GetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) GetUpdatedTime() string {
	return va.UpdatedTime
}

// SetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetUpdatedTime(updatedTime string) {
	va.UpdatedTime = updatedTime
}

// SetAnyUpdatedTime для поля UpdatedTime (строка)
func (va *VerifiedBiZoneAlert) SetAnyUpdatedTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	va.UpdatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetVerifiedBiZoneAlertEndTime для поля EventEndTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) GetEventEndTime() string {
	return va.EventEndTime
}

// SetEventEndTime для поля EventEndTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetEventEndTime(eventEndTime string) {
	va.EventEndTime = eventEndTime
}

// SetAnyEventEndTime для поля EventEndTime (строка)
func (va *VerifiedBiZoneAlert) SetAnyEventEndTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	va.EventEndTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneAlert) GetPlatformHostname() string {
	return va.PlatformHostname
}

// SetPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneAlert) SetPlatformHostname(platformHostname string) {
	va.PlatformHostname = platformHostname
}

// SetAnyPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneAlert) SetAnyPlatformHostname(a any) {
	va.PlatformHostname = fmt.Sprint(a)
}

// GetTitle для поля Title
func (va *VerifiedBiZoneAlert) GetTitle() string {
	return va.Title
}

// SetTitle для поля Title
func (va *VerifiedBiZoneAlert) SetTitle(title string) {
	va.Title = title
}

// SetAnyTitle для поля Title
func (va *VerifiedBiZoneAlert) SetAnyTitle(a any) {
	va.Title = fmt.Sprint(a)
}

// GetRecommendations для поля Recommendations
func (va *VerifiedBiZoneAlert) GetRecommendations() string {
	return va.Recommendations
}

// SetRecommendations для поля Recommendations
func (va *VerifiedBiZoneAlert) SetRecommendations(recommendations string) {
	va.Recommendations = recommendations
}

// SetAnyRecommendations для поля Recommendations
func (va *VerifiedBiZoneAlert) SetAnyRecommendations(a any) {
	va.Recommendations = fmt.Sprint(a)
}

// GetData для поля Data
func (va *VerifiedBiZoneAlert) GetData() BiZoneData {
	return va.Data
}

// SetData для поля Data
func (va *VerifiedBiZoneAlert) SetData(data BiZoneData) {
	va.Data = data
}
