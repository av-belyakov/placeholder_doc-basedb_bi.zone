package datamodels

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ******* Основная структура Alert ********
// *****************************************

// NewVerifiedBiZoneAlert новый объект
func NewVerifiedBiZoneAlert() *VerifiedBiZoneAlert {
	oldTime := time.Date(1970, time.January, 1, 0, 0, 0, 1, time.UTC).Format(time.RFC3339)

	return &VerifiedBiZoneAlert{
		SpecialUUID:        uuid.NewString(),
		CreatedTime:        oldTime,
		UpdatedTime:        oldTime,
		EventEndTime:       oldTime,
		EventStartTime:     oldTime,
		LastDetectionTime:  oldTime,
		FirstDetectionTime: oldTime,
		Tags:               []BiZoneTag(nil),
		Snapshots:          []BiZoneSnapshot(nil),
	}
}

func (va *VerifiedBiZoneAlert) Get() *VerifiedBiZoneAlert {
	return va
}

// GetSpecialUUID униакальное значение поля SpecialUUID
func (va *VerifiedBiZoneAlert) GetSpecialUUID() string {
	return va.SpecialUUID
}

// GetIDNum для поля IDNum
func (va *VerifiedBiZoneAlert) GetIDNum() uint64 {
	return va.IDNum
}

// SetIDNum для поля IDNum
func (va *VerifiedBiZoneAlert) SetIDNum(idNum uint64) {
	va.IDNum = idNum
}

// SetAnyIDNum для поля IDNum
func (va *VerifiedBiZoneAlert) SetAnyIDNum(a any) {
	if v, ok := a.(int); ok {
		va.IDNum = uint64(v)

		return
	}

	if v, ok := a.(float32); ok {
		va.IDNum = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		va.IDNum = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		va.IDNum = v
	}
}

// GetUUID для поля UUID
func (va *VerifiedBiZoneAlert) GetUUID() string {
	return va.UUID
}

// SetUUID для поля UUID
func (va *VerifiedBiZoneAlert) SetUUID(UUID string) {
	va.UUID = UUID
}

// SetAnyUUID для поля UUID
func (va *VerifiedBiZoneAlert) SetAnyUUID(a any) {
	va.UUID = fmt.Sprint(a)
}

// GetExternalID для поля ExternalID
func (va *VerifiedBiZoneAlert) GetExternalID() string {
	return va.ExternalID
}

// SetExternalID для поля ExternalID
func (va *VerifiedBiZoneAlert) SetExternalID(externalID string) {
	va.ExternalID = externalID
}

// SetAnyExternalID для поля ExternalID
func (va *VerifiedBiZoneAlert) SetAnyExternalID(a any) {
	va.ExternalID = fmt.Sprint(a)
}

// GetCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneAlert) GetCustomerSystem() string {
	return va.CustomerSystem
}

// SetCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneAlert) SetCustomerSystem(customerSystem string) {
	va.CustomerSystem = customerSystem
}

// SetAnyCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneAlert) SetAnyCustomerSystem(a any) {
	va.CustomerSystem = fmt.Sprint(a)
}

// GetPlatformType для поля PlatformType
func (va *VerifiedBiZoneAlert) GetPlatformType() string {
	return va.PlatformType
}

// SetPlatformType для поля PlatformType
func (va *VerifiedBiZoneAlert) SetPlatformType(platformType string) {
	va.PlatformType = platformType
}

// SetAnyPlatformType для поля PlatformType
func (va *VerifiedBiZoneAlert) SetAnyPlatformType(a any) {
	va.PlatformType = fmt.Sprint(a)
}

// GetAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneAlert) GetAffectedLogSources() string {
	return va.AffectedLogSources
}

// SetAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneAlert) SetAffectedLogSources(affectedLogSources string) {
	va.AffectedLogSources = affectedLogSources
}

// SetAnyAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneAlert) SetAnyAffectedLogSources(a any) {
	va.AffectedLogSources = fmt.Sprint(a)
}

// GetConfidence для поля Confidence
func (va *VerifiedBiZoneAlert) GetConfidence() string {
	return va.Confidence
}

// SetConfidence для поля Confidence
func (va *VerifiedBiZoneAlert) SetConfidence(confidence string) {
	va.Confidence = confidence
}

// SetAnyConfidence для поля Confidence
func (va *VerifiedBiZoneAlert) SetAnyConfidence(a any) {
	va.Confidence = fmt.Sprint(a)
}

// GetDescription для поля Description
func (va *VerifiedBiZoneAlert) GetDescription() string {
	return va.Description
}

// SetDescription для поля Description
func (va *VerifiedBiZoneAlert) SetDescription(description string) {
	description = strings.ReplaceAll(description, "\t", "")
	description = strings.ReplaceAll(description, "\n", "")

	va.Description = description
}

// SetAnyDescription для поля Description
func (va *VerifiedBiZoneAlert) SetAnyDescription(a any) {
	va.SetDescription(fmt.Sprint(a))
}

// GetDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneAlert) GetDetectionRule() string {
	return va.DetectionRule
}

// SetDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneAlert) SetDetectionRule(detectionRule string) {
	va.DetectionRule = detectionRule
}

// SetAnyDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneAlert) SetAnyDetectionRule(a any) {
	va.DetectionRule = fmt.Sprint(a)
}

// GetCreatedTime для поля CreatedTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) GetCreatedTime() string {
	return va.CreatedTime
}

// SetCreatedTime для поля CreatedTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetCreatedTime(createdTime string) {
	va.CreatedTime = createdTime
}

// SetAnyCreatedTime для поля CreatedTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetAnyCreatedTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.CreatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.CreatedTime = fmt.Sprint(a)
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
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.UpdatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.UpdatedTime = fmt.Sprint(a)
}

// GetEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) GetEventStartTime() string {
	return va.EventStartTime
}

// SetEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetEventStartTime(eventStartTime string) {
	va.EventStartTime = eventStartTime
}

// SetAnyEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetAnyEventStartTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.EventStartTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.EventStartTime = fmt.Sprint(a)
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
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.EventEndTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.EventEndTime = fmt.Sprint(a)
}

// GetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) GetFirstDetectionTime() string {
	return va.FirstDetectionTime
}

// SetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetFirstDetectionTime(firstDetectionTime string) {
	va.FirstDetectionTime = firstDetectionTime
}

// SetAnyFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetAnyFirstDetectionTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.FirstDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.FirstDetectionTime = fmt.Sprint(a)
}

// GetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) GetLastDetectionTime() string {
	return va.LastDetectionTime
}

// SetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetLastDetectionTime(lastDetectionTime string) {
	va.LastDetectionTime = lastDetectionTime
}

// SetAnyLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneAlert) SetAnyLastDetectionTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.LastDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.LastDetectionTime = fmt.Sprint(a)
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
func (va *VerifiedBiZoneAlert) GetData() *BiZoneData {
	return &va.Data
}

// SetData для поля Data
func (va *VerifiedBiZoneAlert) SetData(data BiZoneData) {
	va.Data = data
}

// GetSnapshots для поля Snapshots
func (va *VerifiedBiZoneAlert) GetSnapshots() []BiZoneSnapshot {
	return va.Snapshots
}

// SetSnapshots для поля Snapshots
func (va *VerifiedBiZoneAlert) SetSnapshots(snapshots []BiZoneSnapshot) {
	va.Snapshots = snapshots
}

// GetTags для поля Tags
func (va *VerifiedBiZoneAlert) GetTags() []BiZoneTag {
	return va.Tags
}

// SetTags для поля Tags
func (va *VerifiedBiZoneAlert) SetTags(tags []BiZoneTag) {
	va.Tags = tags
}

// GetAdditionalInformation для поля AdditionalInformation
func (va *VerifiedBiZoneAlert) GetAdditionalInformation() *AdditionalInformation {
	return &va.AdditionalInformation
}

// SetAdditionalInformation для поля AdditionalInformation
func (va *VerifiedBiZoneAlert) SetAdditionalInformation(ai AdditionalInformation) {
	va.AdditionalInformation = ai
}

// ToStringBeautiful форматированный вывод
func (va *VerifiedBiZoneAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)
	wsInc := supportingfunctions.GetWhitespace(num + 1)

	str.WriteString(fmt.Sprintf("%s'id': '%d'\n", ws, va.IDNum))
	str.WriteString(fmt.Sprintf("%s'uuid': '%s'\n", ws, va.UUID))
	str.WriteString(fmt.Sprintf("%s'Title': '%s'\n", ws, va.Title))
	str.WriteString(fmt.Sprintf("%s'severity': '%s'\n", ws, va.Severity))
	str.WriteString(fmt.Sprintf("%s'external_id': '%s'\n", ws, va.ExternalID))
	str.WriteString(fmt.Sprintf("%s'confidence': '%s'\n", ws, va.Confidence))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, va.Description))
	str.WriteString(fmt.Sprintf("%s'created_time': '%s'\n", ws, va.CreatedTime))
	str.WriteString(fmt.Sprintf("%s'updated_time': '%s'\n", ws, va.UpdatedTime))
	str.WriteString(fmt.Sprintf("%s'event_start_time': '%s'\n", ws, va.EventStartTime))
	str.WriteString(fmt.Sprintf("%s'event_end_time': '%s'\n", ws, va.EventEndTime))
	str.WriteString(fmt.Sprintf("%s'first_detection_time': '%s'\n", ws, va.FirstDetectionTime))
	str.WriteString(fmt.Sprintf("%s'last_detection_time': '%s'\n", ws, va.LastDetectionTime))
	str.WriteString(fmt.Sprintf("%s'detection_rule': '%s'\n", ws, va.DetectionRule))
	str.WriteString(fmt.Sprintf("%s'customer_system': '%s'\n", ws, va.CustomerSystem))
	str.WriteString(fmt.Sprintf("%s'recommendations': '%s'\n", ws, va.Recommendations))
	str.WriteString(fmt.Sprintf("%s'platform_type': '%s'\n", ws, va.PlatformType))
	str.WriteString(fmt.Sprintf("%s'platform_hostname': '%s'\n", ws, va.PlatformHostname))
	str.WriteString(fmt.Sprintf("%s'affected_log_sources': '%s'\n", ws, va.AffectedLogSources))
	str.WriteString(fmt.Sprintf("%s'data':\n%s", ws, va.Data.ToStringBeautiful(num+1)))
	str.WriteString(fmt.Sprintf("%s'tags':\n", ws))
	for k, v := range va.Tags {
		str.WriteString(fmt.Sprintf("%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2)))
	}
	str.WriteString(fmt.Sprintf("%s'snapshots':\n", ws))
	for k, v := range va.Snapshots {
		str.WriteString(fmt.Sprintf("%s%d.\n%s", wsInc, k, v.ToStringBeautiful(num+2)))
	}
	str.WriteString(fmt.Sprintf("%s'@additional_information':\n%s", ws, va.AdditionalInformation.ToStringBeautiful(num+1)))

	return str.String()
}
