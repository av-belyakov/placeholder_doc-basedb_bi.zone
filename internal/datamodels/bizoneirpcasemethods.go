package datamodels

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

func NewBiZoneIRPCase() *VerifiedIRPBiZoneCase {
	return &VerifiedIRPBiZoneCase{
		GossopkaErrors:       make(map[string]any),
		ObservedIocs:         make([]BiZoneObservedIocs, 0),
		SecondaryCategoryRef: make([]BiZoneTypeRef, 0),
		DetectionRules:       make([]string, 0),
		SecondaryCategory:    make([]string, 0),
		PlatformHostname:     make([]string, 0),
		WatchersId:           make([]uint64, 0),
		//		ActivityDetected:     make([]any, 0),
		//		Attachments:          make([]any, 0),
		//		Badges:               make([]any, 0),
		//		Emls:                 make([]any, 0),
		//		Slas:                 make([]any, 0),
		//		Tags:                 make([]any, 0),
		//		KeyField:             make([]any, 0),
	}
}

// GetGossopkaErrors список ошибок
func (vc *VerifiedIRPBiZoneCase) GetGossopkaErrors() map[string]any {
	if vc.GossopkaErrors == nil {
		vc.GossopkaErrors = make(map[string]any)
	}

	return vc.GossopkaErrors
}

// SetGossopkaErrors список ошибок
func (vc *VerifiedIRPBiZoneCase) SetGossopkaErrors(errors map[string]any) {
	vc.GossopkaErrors = errors
}

// SetAnyGossopkaErrors список ошибок
func (vc *VerifiedIRPBiZoneCase) SetAnyGossopkaErrors(k string, v any) bool {
	vc.GetGossopkaErrors()

	if _, ok := vc.GossopkaErrors[k]; !ok {
		vc.GossopkaErrors[k] = []string(nil)
	}

	return true
}

// GetObservedIocs список наблюдаемых IOC
func (vc *VerifiedIRPBiZoneCase) GetObservedIocs() []BiZoneObservedIocs {
	if vc.ObservedIocs == nil {
		vc.ObservedIocs = make([]BiZoneObservedIocs, 0)
	}

	return vc.ObservedIocs
}

// SetObservedIocs список наблюдаемых IOC
func (vc *VerifiedIRPBiZoneCase) SetObservedIocs(iocs []BiZoneObservedIocs) {
	vc.ObservedIocs = iocs
}

// AddObservedIocs добавить значение в список наблюдаемых IOC
func (vc *VerifiedIRPBiZoneCase) AddObservedIocs(iocs BiZoneObservedIocs) {
	vc.GetObservedIocs()

	vc.ObservedIocs = append(vc.ObservedIocs, iocs)
}

// GetSecondaryCategoryRef список ссылок на категории
func (vc *VerifiedIRPBiZoneCase) GetSecondaryCategoryRef() []BiZoneTypeRef {
	if vc.SecondaryCategoryRef == nil {
		vc.SecondaryCategoryRef = make([]BiZoneTypeRef, 0)
	}

	return vc.SecondaryCategoryRef
}

// SetSecondaryCategoryRef список ссылок на категории
func (vc *VerifiedIRPBiZoneCase) SetSecondaryCategoryRef(refs []BiZoneTypeRef) {
	vc.SecondaryCategoryRef = refs
}

// AddSecondaryCategoryRef добавляет ссылку в список категорий
func (vc *VerifiedIRPBiZoneCase) AddSecondaryCategoryRef(ref BiZoneTypeRef) {
	vc.GetSecondaryCategoryRef()

	vc.SecondaryCategoryRef = append(vc.SecondaryCategoryRef, ref)
}

// GetActivityDetected список обнаруженной активности
func (vc *VerifiedIRPBiZoneCase) GetActivityDetected() []any {
	if vc.ActivityDetected == nil {
		vc.ActivityDetected = make([]any, 0)
	}

	return vc.ActivityDetected
}

// SetActivityDetected список обнаруженной активности
func (vc *VerifiedIRPBiZoneCase) SetActivityDetected(detected []any) {
	vc.ActivityDetected = detected
}

// AddActivityDetected обнаруженная активность
func (vc *VerifiedIRPBiZoneCase) AddActivityDetected(detected any) {
	vc.GetActivityDetected()

	vc.ActivityDetected = append(vc.ActivityDetected, detected)
}

// GetAttachments список связий
func (vc *VerifiedIRPBiZoneCase) GetAttachments() []any {
	if vc.Attachments == nil {
		vc.Assignee = make([]any, 0)
	}

	return vc.Attachments
}

// SetAttachments список связей
func (vc *VerifiedIRPBiZoneCase) SetAttachments(attachments []any) {
	vc.Attachments = attachments
}

// AddAttachments связи
func (vc *VerifiedIRPBiZoneCase) AddAttachments(attachment any) {
	vc.GetAttachments()

	vc.Attachments = append(vc.Attachments, attachment)
}

// GetBadges значки
func (vc *VerifiedIRPBiZoneCase) GetBadges() []interface{} {
	if vc.Badges == nil {
		vc.Badges = make([]interface{}, 0)
	}

	return vc.Badges
}

// SetBadges значки
func (vc *VerifiedIRPBiZoneCase) SetBadges(badges []interface{}) {
	vc.Badges = badges
}

// AddBadges значок
func (vc *VerifiedIRPBiZoneCase) AddBadges(badge any) {
	vc.GetBadges()

	vc.Badges = append(vc.Badges, badge)
}

// GetEmls
func (vc *VerifiedIRPBiZoneCase) GetEmls() []interface{} {
	if vc.Emls == nil {
		vc.Emls = make([]interface{}, 0)
	}

	return vc.Emls
}

// SetEmls
func (vc *VerifiedIRPBiZoneCase) SetEmls(emls []interface{}) {
	vc.Emls = emls
}

// AddEmls
func (vc *VerifiedIRPBiZoneCase) AddEmls(emls any) {
	vc.GetEmls()

	vc.Emls = append(vc.Emls, emls)
}

// GetSlas
func (vc *VerifiedIRPBiZoneCase) GetSlas() []interface{} {
	if vc.Slas == nil {
		vc.Slas = make([]interface{}, 0)
	}

	return vc.Slas
}

// SetSlas
func (vc *VerifiedIRPBiZoneCase) SetSlas(slas []interface{}) {
	vc.Slas = slas
}

// AddSlas
func (vc *VerifiedIRPBiZoneCase) AddSlas(slas any) {
	vc.GetSlas()

	vc.Slas = append(vc.Slas, slas)
}

// GetTags теги
func (vc *VerifiedIRPBiZoneCase) GetTags() []interface{} {
	if vc.Tags == nil {
		vc.Tags = make([]interface{}, 0)
	}

	return vc.Tags
}

// SetTags теги
func (vc *VerifiedIRPBiZoneCase) SetTags(tags []interface{}) {
	vc.Tags = tags
}

// AddTags теги
func (vc *VerifiedIRPBiZoneCase) AddTags(tags any) {
	vc.GetTags()

	vc.Tags = append(vc.Tags, tags)
}

// GetKeyField ключевые поля
func (vc *VerifiedIRPBiZoneCase) GetKeyField() []interface{} {
	if vc.KeyField == nil {
		vc.KeyField = make([]interface{}, 0)
	}

	return vc.KeyField
}

// SetKeyField ключевые поля
func (vc *VerifiedIRPBiZoneCase) SetKeyField(keyField []interface{}) {
	vc.KeyField = keyField
}

// AddKeyField ключевые поля
func (vc *VerifiedIRPBiZoneCase) AddKeyField(keyField any) {
	vc.GetKeyField()

	vc.KeyField = append(vc.KeyField, keyField)
}

// GetMITRECOV
func (vc *VerifiedIRPBiZoneCase) GetMITRECOV() BiZoneMITRECOV {
	if vc.MITRECOV.GetPersistence() == nil {
		vc.MITRECOV = *NewBiZoneMITRECOV()
	}

	return vc.MITRECOV
}

// SetMITRECOV
func (vc *VerifiedIRPBiZoneCase) SetMITRECOV(mitreconv BiZoneMITRECOV) {
	vc.MITRECOV = mitreconv
}

// GetDetectionRules правила обнаружения
func (vc *VerifiedIRPBiZoneCase) GetDetectionRules() []string {
	if vc.DetectionRules == nil {
		return make([]string, 0)
	}

	return vc.DetectionRules
}

// SetDetectionRules правила обнаружения
func (vc *VerifiedIRPBiZoneCase) SetDetectionRules(rules []string) {
	vc.DetectionRules = rules
}

// AddDetectionRules добавляет правило в список
func (vc *VerifiedIRPBiZoneCase) AddDetectionRules(detectionRule string) {
	vc.GetDetectionRules()

	if slices.Contains(vc.DetectionRules, detectionRule) {
		return
	}

	vc.DetectionRules = append(vc.DetectionRules, detectionRule)
}

// AddAnyDetectionRule добавляет правило в список
func (vc *VerifiedIRPBiZoneCase) AddAnyDetectionRule(a any) {
	vc.AddDetectionRules(fmt.Sprint(a))
}

// GetSecondaryCategory список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) GetSecondaryCategory() []string {
	if vc.SecondaryCategory == nil {
		vc.SecondaryCategory = make([]string, 0)
	}

	return vc.SecondaryCategory
}

// SetSecondaryCategories список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) SetSecondaryCategories(categories []string) {
	vc.SecondaryCategory = categories
}

// AddSecondaryCategory добавляет значение в список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) AddSecondaryCategory(category string) {
	vc.GetSecondaryCategory()

	if slices.Contains(vc.SecondaryCategory, category) {
		return
	}

	vc.SecondaryCategory = append(vc.SecondaryCategory, category)
}

// AddAnySecondaryCategory добавляет значение в список вторичных категорий
func (vc *VerifiedIRPBiZoneCase) AddAnySecondaryCategory(a any) {
	vc.AddSecondaryCategory(fmt.Sprint(a))
}

// GetPlatformHostname список с наименованиями платформы
func (vc *VerifiedIRPBiZoneCase) GetPlatformHostname() []string {
	if vc.PlatformHostname == nil {
		vc.PlatformHostname = make([]string, 0)
	}

	return vc.PlatformHostname
}

// SetPlatformHostname список с наименованиями платформы
func (vc *VerifiedIRPBiZoneCase) SetPlatformHostname(rules []string) {
	vc.PlatformHostname = rules
}

// AddPlatformHostname добавляетнаименование платформы в список
func (vc *VerifiedIRPBiZoneCase) AddPlatformHostname(phn string) {
	vc.GetPlatformHostname()

	if slices.Contains(vc.PlatformHostname, phn) {
		return
	}

	vc.PlatformHostname = append(vc.PlatformHostname, phn)
}

// AddAnyPlatformHostname добавляетнаименование платформы в список
func (vc *VerifiedIRPBiZoneCase) AddAnyPlatformHostname(a any) {
	vc.AddPlatformHostname(fmt.Sprint(a))
}

// GetWatchersId список идентификаторов наблюдения
func (vc *VerifiedIRPBiZoneCase) GetWatchersId() []uint64 {
	if vc.WatchersId == nil {
		vc.WatchersId = make([]uint64, 0)
	}

	return vc.WatchersId
}

// SetWatchersId список идентификаторов наблюдения
func (vc *VerifiedIRPBiZoneCase) SetWatchersId(id []uint64) {
	vc.WatchersId = id
}

// AddWatchersId добавить идентификатор в список идентификаторов наблюдения
func (vc *VerifiedIRPBiZoneCase) AddWatchersId(id uint64) {
	vc.GetWatchersId()

	if slices.Contains(vc.WatchersId, id) {
		return
	}

	vc.WatchersId = append(vc.WatchersId, id)
}

// AddAnyPlatformHostname добавляетнаименование платформы в список
func (vc *VerifiedIRPBiZoneCase) AddAnyWatchersId(a any) {
	switch v := a.(type) {
	case int:
		vc.AddWatchersId(uint64(v))

	case int32:
		vc.AddWatchersId(uint64(v))

	case int64:
		vc.AddWatchersId(uint64(v))

	case float32:
		vc.AddWatchersId(uint64(v))

	case float64:
		vc.AddWatchersId(uint64(v))
	}
}

// GetAssignee правоприемник
func (vc *VerifiedIRPBiZoneCase) GetAssignee() interface{} {
	return vc.Assignee
}

// SetAssignee правоприемник
func (vc *VerifiedIRPBiZoneCase) SetAssignee(a any) {
	vc.Assignee = a
}

// GetAssigneeDisplayName отображаемое имя правопреемника
func (vc *VerifiedIRPBiZoneCase) GetAssigneeDisplayName() interface{} {
	return vc.AssigneeDisplayName
}

// SetAssigneeDisplayName отображаемое имя правопреемника
func (vc *VerifiedIRPBiZoneCase) SetAssigneeDisplayName(a any) {
	vc.AssigneeDisplayName = a
}

// GetService сервис
func (vc *VerifiedIRPBiZoneCase) GetService() interface{} {
	return vc.Service
}

// SetService сервис
func (vc *VerifiedIRPBiZoneCase) SetService(a any) {
	vc.Service = a
}

// GetResolutionDate дата разрешения
func (vc *VerifiedIRPBiZoneCase) GetResolutionDate() interface{} {
	return vc.ResolutionDate
}

// SetResolutionDate дата разрешения
func (vc *VerifiedIRPBiZoneCase) SetResolutionDate(a any) {
	vc.ResolutionDate = a
}

// GetResolutionName имя разрешения
func (vc *VerifiedIRPBiZoneCase) GetResolutionName() interface{} {
	return vc.ResolutionName
}

// SetResolutionName имя разрешения
func (vc *VerifiedIRPBiZoneCase) SetResolutionName(a any) {
	vc.ResolutionName = a
}

// GetResolutionNameRef ссылка на имя разрешения
func (vc *VerifiedIRPBiZoneCase) GetResolutionNameRef() interface{} {
	return vc.ResolutionNameRef
}

// SetResolutionNameRef ссылка на имя разрешения
func (vc *VerifiedIRPBiZoneCase) SetResolutionNameRef(a any) {
	vc.ResolutionNameRef = a
}

// GetGossopkaId идентификатор
func (vc *VerifiedIRPBiZoneCase) GetGossopkaId() interface{} {
	return vc.GossopkaId
}

// SetGossopkaId идентификатор
func (vc *VerifiedIRPBiZoneCase) SetGossopkaId(a any) {
	vc.GossopkaId = a
}

// GetGossopkaSendTime установленное время (возможно типа string или time)
func (vc *VerifiedIRPBiZoneCase) GetGossopkaSendTime() interface{} {
	return vc.GossopkaSendTime
}

// SetGossopkaSendTime установленное время (возможно типа string или time)
func (vc *VerifiedIRPBiZoneCase) SetGossopkaSendTime(a any) {
	vc.GossopkaSendTime = a
}

// GetGtiId идентификатор
func (vc *VerifiedIRPBiZoneCase) GetGtiId() interface{} {
	return vc.GtiId
}

// SetGtiId идентификатор
func (vc *VerifiedIRPBiZoneCase) SetGtiId(a any) {
	vc.GtiId = a
}

// GetGtiSendTime некое время
func (vc *VerifiedIRPBiZoneCase) GetGtiSendTime() interface{} {
	return vc.GtiSendTime
}

// SetGtiSendTime некое время
func (vc *VerifiedIRPBiZoneCase) SetGtiSendTime(a any) {
	vc.GtiSendTime = a
}

// GetCustomerStarRating звёздный рейтинг потребителей
func (vc *VerifiedIRPBiZoneCase) GetCustomerStarRating() interface{} {
	return vc.CustomerStarRating
}

// SetCustomerStarRating звёздный рейтинг потребителей
func (vc *VerifiedIRPBiZoneCase) SetCustomerStarRating(a any) {
	vc.CustomerStarRating = a
}

// GetFirstSentNotificationTime время первого сообщения
func (vc *VerifiedIRPBiZoneCase) GetFirstSentNotificationTime() interface{} {
	return vc.FirstSentNotificationTime
}

// SetFirstSentNotificationTime время первого сообщения
func (vc *VerifiedIRPBiZoneCase) SetFirstSentNotificationTime(a any) {
	vc.FirstSentNotificationTime = a
}

// GetResponseTeam группа ответа
func (vc *VerifiedIRPBiZoneCase) GetResponseTeam() interface{} {
	return vc.ResponseTeam
}

// SetResponseTeam группа ответа
func (vc *VerifiedIRPBiZoneCase) SetResponseTeam(a any) {
	vc.ResponseTeam = a
}

// GetGossopkaKeyLink ссылка на ключ
func (vc *VerifiedIRPBiZoneCase) GetGossopkaKeyLink() interface{} {
	return vc.GossopkaKeyLink
}

// SetGossopkaKeyLink ссылка на ключ
func (vc *VerifiedIRPBiZoneCase) SetGossopkaKeyLink(a any) {
	vc.GossopkaKeyLink = a
}

// GetIssueTypeRef ссылка на тип проблеммы
func (vc *VerifiedIRPBiZoneCase) GetIssueTypeRef() BiZoneTypeRef {
	return vc.IssueTypeRef
}

// SetIssueTypeRef ссылка на тип проблеммы
func (vc *VerifiedIRPBiZoneCase) SetIssueTypeRef(ref BiZoneTypeRef) {
	vc.IssueTypeRef = ref
}

// GetPrimaryCategoryRef ссылка на первичную категорию
func (vc *VerifiedIRPBiZoneCase) GetPrimaryCategoryRef() BiZoneTypeRef {
	return vc.PrimaryCategoryRef
}

// SetPrimaryCategoryRef ссылка на первичную категорию
func (vc *VerifiedIRPBiZoneCase) SetPrimaryCategoryRef(ref BiZoneTypeRef) {
	vc.PrimaryCategoryRef = ref
}

// GetAttackType тип атаки
func (vc *VerifiedIRPBiZoneCase) GetAttackType() string {
	return vc.AttackType
}

// SetAttackType тип атаки
func (vc *VerifiedIRPBiZoneCase) SetAttackType(at string) {
	vc.AttackType = at
}

// SetAnyAttackType тип атаки
func (vc *VerifiedIRPBiZoneCase) SetAnyAttackType(a any) {
	vc.SetAttackType(fmt.Sprint(a))
}

// GetCreated время создания (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetCreated() string {
	return vc.Created
}

// SetCreated время создания (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetCreated(str string) error {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	vc.Created = t.String()

	return nil
}

// SetAnyCreated время создания
func (vc *VerifiedIRPBiZoneCase) SetAnyCreated(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetCreated(v)
	}

	return errors.New("type conversion error")
}

// GetCreatorDisplayName имя создающего дисплей
func (vc *VerifiedIRPBiZoneCase) GetCreatorDisplayName() string {
	return vc.CreatorDisplayName
}

// SetCreatorDisplayName имя создающего дисплей
func (vc *VerifiedIRPBiZoneCase) SetCreatorDisplayName(cdn string) {
	vc.CreatorDisplayName = cdn
}

// SetAnyCreatorDisplayName имя создающего дисплей
func (vc *VerifiedIRPBiZoneCase) SetAnyCreatorDisplayName(a any) {
	vc.SetCreatorDisplayName(fmt.Sprint(a))
}

// GetDetectionDate дата обнаружения (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetDetectionDate() string {
	return vc.DetectionDate
}

// SetDetectionDate дата обнаружения (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetDetectionDate(str string) error {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	vc.DetectionDate = t.String()

	return nil
}

// SetAnyDetectionDate дата обнаружения
func (vc *VerifiedIRPBiZoneCase) SetAnyDetectionDate(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetDetectionDate(v)
	}

	return errors.New("type conversion error")
}

// GetKey ключ
func (vc *VerifiedIRPBiZoneCase) GetKey() string {
	return vc.Key
}

// SetKey ключ
func (vc *VerifiedIRPBiZoneCase) SetKey(key string) {
	vc.SetKey(key)
}

// SetAnyKey ключ
func (vc *VerifiedIRPBiZoneCase) SetAnyKey(a any) {
	vc.Key = fmt.Sprint(a)
}

// GetIssueType тип проблемы
func (vc *VerifiedIRPBiZoneCase) GetIssueType() string {
	return vc.IssueType
}

// SetIssueType тип проблемы
func (vc *VerifiedIRPBiZoneCase) SetIssueType(it string) {
	vc.SetIssueType(it)
}

// SetAnyIssueType тип проблемы
func (vc *VerifiedIRPBiZoneCase) SetAnyIssueType(a any) {
	vc.IssueType = fmt.Sprint(a)
}

// GetPriority приоритет
func (vc *VerifiedIRPBiZoneCase) GetPriority() string {
	return vc.Priority
}

// SetPriority приоритет
func (vc *VerifiedIRPBiZoneCase) SetPriority(p string) {
	vc.Priority = p
}

// SetAnyPriority приоритет
func (vc *VerifiedIRPBiZoneCase) SetAnyPriority(a any) {
	vc.SetPriority(fmt.Sprint(a))
}

// GetSummary резюме
func (vc *VerifiedIRPBiZoneCase) GetSummary() string {
	return vc.Summary
}

// SetSummary резюме
func (vc *VerifiedIRPBiZoneCase) SetSummary(sum string) {
	vc.Summary = sum
}

// SetAnySummary резюме
func (vc *VerifiedIRPBiZoneCase) SetAnySummary(a any) {
	vc.SetSummary(fmt.Sprint(a))
}

// GetDescription описание
func (vc *VerifiedIRPBiZoneCase) GetDescription() string {
	return vc.Description
}

// SetDescription описание
func (vc *VerifiedIRPBiZoneCase) SetDescription(d string) {
	vc.Description = d
}

// SetAnyDescription описание
func (vc *VerifiedIRPBiZoneCase) SetAnyDescription(a any) {
	vc.SetDescription(fmt.Sprint(a))
}

// GetRenderedDescription отрисованое описание
func (vc *VerifiedIRPBiZoneCase) GetRenderedDescription() string {
	return vc.RenderedDescription
}

// SetRenderedDescription отрисованое описание
func (vc *VerifiedIRPBiZoneCase) SetRenderedDescription(rd string) {
	vc.RenderedDescription = rd
}

// SetAnyRenderedDescription отрисованое описание
func (vc *VerifiedIRPBiZoneCase) SetAnyRenderedDescription(a any) {
	vc.SetRenderedDescription(fmt.Sprint(a))
}

// GetStatus статус
func (vc *VerifiedIRPBiZoneCase) GetStatus() string {
	return vc.Status
}

// SetStatus статус
func (vc *VerifiedIRPBiZoneCase) SetStatus(s string) {
	vc.Status = s
}

// SetAnyStatus статус
func (vc *VerifiedIRPBiZoneCase) SetAnyStatus(a any) {
	vc.SetStatus(fmt.Sprint(a))
}

// GetStatusDescription описание статуса
func (vc *VerifiedIRPBiZoneCase) GetStatusDescription() string {
	return vc.StatusDescription
}

// SetStatusDescription описание статуса
func (vc *VerifiedIRPBiZoneCase) SetStatusDescription(sd string) {
	vc.StatusDescription = sd
}

// SetAnyStatusDescription описание статуса
func (vc *VerifiedIRPBiZoneCase) SetAnyStatusDescription(a any) {
	vc.SetStatusDescription(fmt.Sprint(a))
}

// GetReporterDisplayName нименование докладчика
func (vc *VerifiedIRPBiZoneCase) GetReporterDisplayName() string {
	return vc.ReporterDisplayName
}

// SetReporterDisplayName нименование докладчика
func (vc *VerifiedIRPBiZoneCase) SetReporterDisplayName(rdn string) {
	vc.ReporterDisplayName = rdn
}

// SetAnyReporterDisplayName нименование докладчика
func (vc *VerifiedIRPBiZoneCase) SetAnyReporterDisplayName(a any) {
	vc.SetReporterDisplayName(fmt.Sprint(a))
}

// GetReporterEmailAddress email адрес докладчика
func (vc *VerifiedIRPBiZoneCase) GetReporterEmailAddress() string {
	return vc.ReporterEmailAddress
}

// SetReporterEmailAddress email адрес докладчика
func (vc *VerifiedIRPBiZoneCase) SetReporterEmailAddress(re string) {
	vc.ReporterEmailAddress = re
}

// SetAnyReporterEmailAddress email адрес докладчика
func (vc *VerifiedIRPBiZoneCase) SetAnyReporterEmailAddress(a any) {
	vc.SetReporterEmailAddress(fmt.Sprint(a))
}

// GetPrimaryCategory первичная категория
func (vc *VerifiedIRPBiZoneCase) GetPrimaryCategory() string {
	return vc.PrimaryCategory
}

// SetPrimaryCategory первичная категория
func (vc *VerifiedIRPBiZoneCase) SetPrimaryCategory(pc string) {
	vc.PrimaryCategory = pc
}

// SetAnyPrimaryCategory первичная категория
func (vc *VerifiedIRPBiZoneCase) SetAnyPrimaryCategory(a any) {
	vc.SetPrimaryCategory(fmt.Sprint(a))
}

// GetUpdated время обновления (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetUpdated() string {
	return vc.Updated
}

// SetUpdated время обновления (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetUpdated(str string) error {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	vc.Updated = t.String()

	return nil
}

// SetAnyUpdated время обновления
func (vc *VerifiedIRPBiZoneCase) SetAnyUpdated(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetUpdated(v)
	}

	return errors.New("type conversion error")
}

// GetTimestamp временная метка (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) GetTimestamp() string {
	return vc.Timestamp
}

// SetTimestamp временная метка (формат времени RFC3339)
func (vc *VerifiedIRPBiZoneCase) SetTimestamp(str string) error {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	vc.Timestamp = t.String()

	return nil
}

// SetAnyTimestamp временная метка
func (vc *VerifiedIRPBiZoneCase) SetAnyTimestamp(a any) error {
	if v, ok := a.(string); ok {
		return vc.SetTimestamp(v)
	}

	return errors.New("type conversion error")
}

// GetResolutionDetailed подробная информация о разрешении
func (vc *VerifiedIRPBiZoneCase) GetResolutionDetailed() string {
	return vc.ResolutionDetailed
}

// SetResolutionDetailed подробная информация о разрешении
func (vc *VerifiedIRPBiZoneCase) SetResolutionDetailed(rd string) {
	vc.ResolutionDetailed = rd
}

// SetAnyResolutionDetailed подробная информация о разрешении
func (vc *VerifiedIRPBiZoneCase) SetAnyResolutionDetailed(rd string) {
	vc.SetResolutionDetailed(rd)
}

// GetPlatformURL
func (vc *VerifiedIRPBiZoneCase) GetPlatformURL() string {
	return vc.PlatformURL
}

// SetPlatformURL
func (vc *VerifiedIRPBiZoneCase) SetPlatformURL(purl string) {
	vc.PlatformURL = purl
}

// SetAnyPlatformURL
func (vc *VerifiedIRPBiZoneCase) SetAnyPlatformURL(purl string) {
	vc.SetPlatformURL(purl)
}

func (vc *VerifiedIRPBiZoneCase) GetCustomerStarRatingComment() string {
	return vc.customerStarRatingComment
}

func (vc *VerifiedIRPBiZoneCase) GetRecommendations() string {
	return vc.recommendations
}

func (vc *VerifiedIRPBiZoneCase) GetParsedActivityDetected() string {
	return vc.parsedActivityDetected
}

func (vc *VerifiedIRPBiZoneCase) GetAffectedService() string {
	return vc.affectedService
}

func (vc *VerifiedIRPBiZoneCase) GetFakeAsPath() string {
	return vc.fakeAsPath
}

func (vc *VerifiedIRPBiZoneCase) GetFakePrefix() string {
	return vc.fakePrefix
}

func (vc *VerifiedIRPBiZoneCase) GetFpType() string {
	return vc.fpType
}

func (vc *VerifiedIRPBiZoneCase) GetLookingGlass() string {
	return vc.lookingGlass
}

func (vc *VerifiedIRPBiZoneCase) GetSpamRecipients() string {
	return vc.spamRecipients
}

func (vc *VerifiedIRPBiZoneCase) GetTLP() string {
	return vc.tlp
}

func (vc *VerifiedIRPBiZoneCase) GetUsualPrefix() string {
	return vc.usualPrefix
}

func (vc *VerifiedIRPBiZoneCase) GetUsualAsPath() string {
	return vc.usualAsPath
}

func (vc *VerifiedIRPBiZoneCase) GetSource() string {
	return vc.source
}

func (vc *VerifiedIRPBiZoneCase) GetExternalId() string {
	return vc.externalId
}

func (vc *VerifiedIRPBiZoneCase) GetUnderliningSource() string {
	return vc.underliningSource
}

func (vc *VerifiedIRPBiZoneCase) GetUpdatedAll() time.Time {
	return vc.updatedAll
}

func (vc *VerifiedIRPBiZoneCase) GetUpdatedAllString() string {
	if vc.updatedAll.IsZero() {
		return ""
	}
	return vc.updatedAll.Format(time.RFC3339)
}

func (vc *VerifiedIRPBiZoneCase) GetGossopkaKey() string {
	return vc.gossopkaKey
}

func (vc *VerifiedIRPBiZoneCase) GetWatchers() BiZoneWatcher {
	return vc.watchers
}

func (vc *VerifiedIRPBiZoneCase) GetID() int {
	return vc.id
}

func (vc *VerifiedIRPBiZoneCase) GetSystem() int {
	return vc.system
}

func (vc *VerifiedIRPBiZoneCase) GetCancelGossopkaSendButton() bool {
	return vc.cancelGossopkaSendButton
}

func (vc *VerifiedIRPBiZoneCase) GetIsVisibleForCustomer() bool {
	return vc.isVisibleForCustomer
}

func (vc *VerifiedIRPBiZoneCase) GetShowGossopkaButton() bool {
	return vc.showGossopkaButton
}

func (vc *VerifiedIRPBiZoneCase) GetShowGtiButton() bool {
	return vc.showGtiButton
}
