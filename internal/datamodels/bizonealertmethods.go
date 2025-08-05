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

// GetSeverity для поля Severity
func (e *VerifiedBiZoneAlert) GetSeverity() string {
	return e.Severity
}

// SetSeverity для поля Severity
func (e *VerifiedBiZoneAlert) SetSeverity(severity string) {
	e.Severity = severity
}

// SetAnySeverity для поля Severity
func (e *VerifiedBiZoneAlert) SetAnySeverity(a any) {
	e.Severity = fmt.Sprint(a)
}

// GetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (e *VerifiedBiZoneAlert) GetUpdatedTime() string {
	return e.UpdatedTime
}

// SetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (e *VerifiedBiZoneAlert) SetUpdatedTime(updatedTime string) {
	e.UpdatedTime = updatedTime
}

// SetAnyUpdatedTime для поля UpdatedTime (строка)
func (e *VerifiedBiZoneAlert) SetAnyUpdatedTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	e.UpdatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetVerifiedBiZoneAlertEndTime для поля EventEndTime (формат RFC3339)
func (e *VerifiedBiZoneAlert) GetEventEndTime() string {
	return e.EventEndTime
}

// SetEventEndTime для поля EventEndTime (формат RFC3339)
func (e *VerifiedBiZoneAlert) SetEventEndTime(eventEndTime string) {
	e.EventEndTime = eventEndTime
}

// SetAnyEventEndTime для поля EventEndTime (строка)
func (e *VerifiedBiZoneAlert) SetAnyEventEndTime(a any) {
	tmp := supportingfunctions.ConversionAnyToInt(a)
	e.EventEndTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetPlatformHostname для поля PlatformHostname
func (e *VerifiedBiZoneAlert) GetPlatformHostname() string {
	return e.PlatformHostname
}

// SetPlatformHostname для поля PlatformHostname
func (e *VerifiedBiZoneAlert) SetPlatformHostname(platformHostname string) {
	e.PlatformHostname = platformHostname
}

// SetAnyPlatformHostname для поля PlatformHostname
func (e *VerifiedBiZoneAlert) SetAnyPlatformHostname(a any) {
	e.PlatformHostname = fmt.Sprint(a)
}

// GetTitle для поля Title
func (e *VerifiedBiZoneAlert) GetTitle() string {
	return e.Title
}

// SetTitle для поля Title
func (e *VerifiedBiZoneAlert) SetTitle(title string) {
	e.Title = title
}

// SetAnyTitle для поля Title
func (e *VerifiedBiZoneAlert) SetAnyTitle(a any) {
	e.Title = fmt.Sprint(a)
}

// GetRecommendations для поля Recommendations
func (e *VerifiedBiZoneAlert) GetRecommendations() string {
	return e.Recommendations
}

// SetRecommendations для поля Recommendations
func (e *VerifiedBiZoneAlert) SetRecommendations(recommendations string) {
	e.Recommendations = recommendations
}

// SetAnyRecommendations для поля Recommendations
func (e *VerifiedBiZoneAlert) SetAnyRecommendations(a any) {
	e.Recommendations = fmt.Sprint(a)
}

// GetData для поля Data
func (e *VerifiedBiZoneAlert) GetData() BiZoneData {
	return e.Data
}

// SetData для поля Data
func (e *VerifiedBiZoneAlert) SetData(data BiZoneData) {
	e.Data = data
}
