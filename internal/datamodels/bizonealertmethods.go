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
func (e *VerifiedBiZoneAlert) SetAnySeverity(i any) {
	e.Severity = fmt.Sprint(i)
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
func (e *VerifiedBiZoneAlert) SetAnyUpdatedTime(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
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
func (e *VerifiedBiZoneAlert) SetAnyEventEndTime(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
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
func (e *VerifiedBiZoneAlert) SetAnyPlatformHostname(i any) {
	e.PlatformHostname = fmt.Sprint(i)
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
func (e *VerifiedBiZoneAlert) SetAnyTitle(i any) {
	e.Title = fmt.Sprint(i)
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
func (e *VerifiedBiZoneAlert) SetAnyRecommendations(i any) {
	e.Recommendations = fmt.Sprint(i)
}

// GetData для поля Data
func (e *VerifiedBiZoneAlert) GetData() BiZoneData {
	return e.Data
}

// SetData для поля Data
func (e *VerifiedBiZoneAlert) SetData(data BiZoneData) {
	e.Data = data
}
