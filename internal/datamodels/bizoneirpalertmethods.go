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

// NewVerifiedBiZoneIRPAlert новый объект
func NewVerifiedBiZoneIRPAlert() *VerifiedBiZoneIRPAlert {
	oldTime := time.Date(1970, time.January, 1, 0, 0, 0, 1, time.UTC).Format(time.RFC3339)

	return &VerifiedBiZoneIRPAlert{
		SpecialUUID:        uuid.NewString(),
		CreatedTime:        oldTime,
		UpdatedTime:        oldTime,
		EventEndTime:       oldTime,
		EventStartTime:     oldTime,
		LastDetectionTime:  oldTime,
		FirstDetectionTime: oldTime,
		Tags:               []BiZoneIRPTag(nil),
		Snapshots:          []BiZoneIRPSnapshot(nil),
	}
}

func (va *VerifiedBiZoneIRPAlert) Get() *VerifiedBiZoneIRPAlert {
	return va
}

// GetSpecialUUID униакальное значение поля SpecialUUID
func (va *VerifiedBiZoneIRPAlert) GetSpecialUUID() string {
	return va.SpecialUUID
}

// GetIDNum для поля IDNum
func (va *VerifiedBiZoneIRPAlert) GetIDNum() uint64 {
	return va.IDNum
}

// SetIDNum для поля IDNum
func (va *VerifiedBiZoneIRPAlert) SetIDNum(idNum uint64) {
	va.IDNum = idNum
}

// SetAnyIDNum для поля IDNum
func (va *VerifiedBiZoneIRPAlert) SetAnyIDNum(a any) {
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
func (va *VerifiedBiZoneIRPAlert) GetUUID() string {
	return va.UUID
}

// SetUUID для поля UUID
func (va *VerifiedBiZoneIRPAlert) SetUUID(UUID string) {
	va.UUID = UUID
}

// SetAnyUUID для поля UUID
func (va *VerifiedBiZoneIRPAlert) SetAnyUUID(a any) {
	va.UUID = fmt.Sprint(a)
}

// GetExternalID для поля ExternalID
func (va *VerifiedBiZoneIRPAlert) GetExternalID() string {
	return va.ExternalID
}

// SetExternalID для поля ExternalID
func (va *VerifiedBiZoneIRPAlert) SetExternalID(externalID string) {
	va.ExternalID = externalID
}

// SetAnyExternalID для поля ExternalID
func (va *VerifiedBiZoneIRPAlert) SetAnyExternalID(a any) {
	va.ExternalID = fmt.Sprint(a)
}

// GetCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneIRPAlert) GetCustomerSystem() string {
	return va.CustomerSystem
}

// SetCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneIRPAlert) SetCustomerSystem(customerSystem string) {
	va.CustomerSystem = customerSystem
}

// SetAnyCustomerSystem для поля CustomerSystem
func (va *VerifiedBiZoneIRPAlert) SetAnyCustomerSystem(a any) {
	va.CustomerSystem = fmt.Sprint(a)
}

// GetPlatformType для поля PlatformType
func (va *VerifiedBiZoneIRPAlert) GetPlatformType() string {
	return va.PlatformType
}

// SetPlatformType для поля PlatformType
func (va *VerifiedBiZoneIRPAlert) SetPlatformType(platformType string) {
	va.PlatformType = platformType
}

// SetAnyPlatformType для поля PlatformType
func (va *VerifiedBiZoneIRPAlert) SetAnyPlatformType(a any) {
	va.PlatformType = fmt.Sprint(a)
}

// GetAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneIRPAlert) GetAffectedLogSources() string {
	return va.AffectedLogSources
}

// SetAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneIRPAlert) SetAffectedLogSources(affectedLogSources string) {
	va.AffectedLogSources = affectedLogSources
}

// SetAnyAffectedLogSources для поля AffectedLogSources
func (va *VerifiedBiZoneIRPAlert) SetAnyAffectedLogSources(a any) {
	va.AffectedLogSources = fmt.Sprint(a)
}

// GetConfidence для поля Confidence
func (va *VerifiedBiZoneIRPAlert) GetConfidence() string {
	return va.Confidence
}

// SetConfidence для поля Confidence
func (va *VerifiedBiZoneIRPAlert) SetConfidence(confidence string) {
	va.Confidence = confidence
}

// SetAnyConfidence для поля Confidence
func (va *VerifiedBiZoneIRPAlert) SetAnyConfidence(a any) {
	va.Confidence = fmt.Sprint(a)
}

// GetDescription для поля Description
func (va *VerifiedBiZoneIRPAlert) GetDescription() string {
	return va.Description
}

// SetDescription для поля Description
func (va *VerifiedBiZoneIRPAlert) SetDescription(description string) {
	description = strings.ReplaceAll(description, "\t", "")
	description = strings.ReplaceAll(description, "\n", "")

	va.Description = description
}

// SetAnyDescription для поля Description
func (va *VerifiedBiZoneIRPAlert) SetAnyDescription(a any) {
	va.SetDescription(fmt.Sprint(a))
}

// GetDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneIRPAlert) GetDetectionRule() string {
	return va.DetectionRule
}

// SetDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneIRPAlert) SetDetectionRule(detectionRule string) {
	va.DetectionRule = detectionRule
}

// SetAnyDetectionRule для поля DetectionRule
func (va *VerifiedBiZoneIRPAlert) SetAnyDetectionRule(a any) {
	va.DetectionRule = fmt.Sprint(a)
}

// GetCreatedTime для поля CreatedTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetCreatedTime() string {
	return va.CreatedTime
}

// SetCreatedTime для поля CreatedTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetCreatedTime(createdTime string) {
	va.CreatedTime = createdTime
}

// SetAnyCreatedTime для поля CreatedTime
func (va *VerifiedBiZoneIRPAlert) SetAnyCreatedTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.CreatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.SetCreatedTime(fmt.Sprint(a))
}

// GetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetUpdatedTime() string {
	return va.UpdatedTime
}

// SetUpdatedTime для поля UpdatedTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetUpdatedTime(updatedTime string) {
	va.UpdatedTime = updatedTime
}

// SetAnyUpdatedTime для поля UpdatedTime (строка)
func (va *VerifiedBiZoneIRPAlert) SetAnyUpdatedTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.UpdatedTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.UpdatedTime = fmt.Sprint(a)
}

// GetEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetEventStartTime() string {
	return va.EventStartTime
}

// SetEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetEventStartTime(eventStartTime string) {
	va.EventStartTime = eventStartTime
}

// SetAnyEventStartTime для поля EventStartTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetAnyEventStartTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.EventStartTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.EventStartTime = fmt.Sprint(a)
}

// GetVerifiedBiZoneIRPAlertEndTime для поля EventEndTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetEventEndTime() string {
	return va.EventEndTime
}

// SetEventEndTime для поля EventEndTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetEventEndTime(eventEndTime string) {
	va.EventEndTime = eventEndTime
}

// SetAnyEventEndTime для поля EventEndTime (строка)
func (va *VerifiedBiZoneIRPAlert) SetAnyEventEndTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.EventEndTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.EventEndTime = fmt.Sprint(a)
}

// GetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetFirstDetectionTime() string {
	return va.FirstDetectionTime
}

// SetFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetFirstDetectionTime(firstDetectionTime string) {
	va.FirstDetectionTime = firstDetectionTime
}

// SetAnyFirstDetectionTime для поля FirstDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetAnyFirstDetectionTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.FirstDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.FirstDetectionTime = fmt.Sprint(a)
}

// GetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) GetLastDetectionTime() string {
	return va.LastDetectionTime
}

// SetLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetLastDetectionTime(lastDetectionTime string) {
	va.LastDetectionTime = lastDetectionTime
}

// SetAnyLastDetectionTime для поля LastDetectionTime (формат RFC3339)
func (va *VerifiedBiZoneIRPAlert) SetAnyLastDetectionTime(a any) {
	//это только когда время в типе unixtime
	//tmp := supportingfunctions.ConversionAnyToInt(a)
	//va.LastDetectionTime = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
	va.LastDetectionTime = fmt.Sprint(a)
}

// GetPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneIRPAlert) GetPlatformHostname() string {
	return va.PlatformHostname
}

// SetPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneIRPAlert) SetPlatformHostname(platformHostname string) {
	va.PlatformHostname = platformHostname
}

// SetAnyPlatformHostname для поля PlatformHostname
func (va *VerifiedBiZoneIRPAlert) SetAnyPlatformHostname(a any) {
	va.PlatformHostname = fmt.Sprint(a)
}

// GetTitle для поля Title
func (va *VerifiedBiZoneIRPAlert) GetTitle() string {
	return va.Title
}

// SetTitle для поля Title
func (va *VerifiedBiZoneIRPAlert) SetTitle(title string) {
	va.Title = title
}

// SetAnyTitle для поля Title
func (va *VerifiedBiZoneIRPAlert) SetAnyTitle(a any) {
	va.Title = fmt.Sprint(a)
}

// GetSeverity для поля Severity
func (va *VerifiedBiZoneIRPAlert) GetSeverity() string {
	return va.Severity
}

// SetSeverity для поля Severity
func (va *VerifiedBiZoneIRPAlert) SetSeverity(severity string) {
	va.Severity = severity
}

// SetAnySeverity для поля Severity
func (va *VerifiedBiZoneIRPAlert) SetAnySeverity(a any) {
	va.Severity = fmt.Sprint(a)
}

// GetRecommendations для поля Recommendations
func (va *VerifiedBiZoneIRPAlert) GetRecommendations() string {
	return va.Recommendations
}

// SetRecommendations для поля Recommendations
func (va *VerifiedBiZoneIRPAlert) SetRecommendations(recommendations string) {
	va.Recommendations = recommendations
}

// SetAnyRecommendations для поля Recommendations
func (va *VerifiedBiZoneIRPAlert) SetAnyRecommendations(a any) {
	va.Recommendations = fmt.Sprint(a)
}

// GetData для поля Data
func (va *VerifiedBiZoneIRPAlert) GetData() *BiZoneIRPData {
	return &va.Data
}

// SetData для поля Data
func (va *VerifiedBiZoneIRPAlert) SetData(data BiZoneIRPData) {
	va.Data = data
}

// GetSnapshots для поля Snapshots
func (va *VerifiedBiZoneIRPAlert) GetSnapshots() []BiZoneIRPSnapshot {
	return va.Snapshots
}

// SetSnapshots для поля Snapshots
func (va *VerifiedBiZoneIRPAlert) SetSnapshots(snapshots []BiZoneIRPSnapshot) {
	va.Snapshots = snapshots
}

// GetTags для поля Tags
func (va *VerifiedBiZoneIRPAlert) GetTags() []BiZoneIRPTag {
	return va.Tags
}

// SetTags для поля Tags
func (va *VerifiedBiZoneIRPAlert) SetTags(tags []BiZoneIRPTag) {
	va.Tags = tags
}

// GetAdditionalInformation для поля AdditionalInformation
func (va *VerifiedBiZoneIRPAlert) GetAdditionalInformation() *AdditionalInformation {
	return &va.AdditionalInformation
}

// SetAdditionalInformation для поля AdditionalInformation
func (va *VerifiedBiZoneIRPAlert) SetAdditionalInformation(ai AdditionalInformation) {
	va.AdditionalInformation = ai
}

// ToStringBeautiful форматированный вывод
func (va *VerifiedBiZoneIRPAlert) ToStringBeautiful(num int) string {
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
