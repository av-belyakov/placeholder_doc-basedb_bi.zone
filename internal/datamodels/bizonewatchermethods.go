package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// GetId идентификатор
func (b *BiZoneWatcher) GetId() string {
	return b.ID
}

// SetId идентификатор
func (b *BiZoneWatcher) SetId(id string) {
	// Здесь можно добавить валидацию
	b.ID = id
}

// SetAnyId идентификатор
func (b *BiZoneWatcher) SetAnyId(a any) {
	b.ID = fmt.Sprint(a)
}

// GetUsername имя пользователя
func (b *BiZoneWatcher) GetUsername() string {
	return b.Username
}

// SetUsername имя пользователя
func (b *BiZoneWatcher) SetUsername(username string) {
	b.Username = username
}

// SetAnyUsername имя пользователя
func (b *BiZoneWatcher) SetAnyUsername(a any) {
	b.Username = fmt.Sprint(a)
}

// GetEmail электронный адрес
func (b *BiZoneWatcher) GetEmail() string {
	return b.Email
}

// SetEmail электронный адрес
func (b *BiZoneWatcher) SetEmail(email string) {
	b.Email = email
}

// SetAnyEmail электронный адрес
func (b *BiZoneWatcher) SetAnyEmail(a any) {
	b.Email = fmt.Sprint(a)
}

// GetFirstName имя пользователя
func (b *BiZoneWatcher) GetFirstName() string {
	return b.FirstName
}

// SetFirstName имя пользователя
func (b *BiZoneWatcher) SetFirstName(firstName string) {
	b.FirstName = firstName
}

// SetAnyFirstName имя пользователя
func (b *BiZoneWatcher) SetAnyFirstName(a any) {
	b.FirstName = fmt.Sprint(a)
}

// GetLastName фамилия пользователя
func (b *BiZoneWatcher) GetLastName() string {
	return b.LastName
}

// SetLastName фамилия пользователя
func (b *BiZoneWatcher) SetLastName(lastName string) {
	b.LastName = lastName
}

// SetAnyLastName фамилия пользователя
func (b *BiZoneWatcher) SetAnyLastName(a any) {
	b.LastName = fmt.Sprint(a)
}

// GetPatronimic
func (b *BiZoneWatcher) GetPatronimic() string {
	return b.Patronimic
}

// SetPatronimic
func (b *BiZoneWatcher) SetPatronimic(patronimic string) {
	b.Patronimic = patronimic
}

// SetAnyPatronimic
func (b *BiZoneWatcher) SetAnyPatronimic(a any) {
	b.Patronimic = fmt.Sprint(a)
}

// GetIsActive
func (b *BiZoneWatcher) GetIsActive() bool {
	return b.IsActive
}

// SetIsActive
func (b *BiZoneWatcher) SetIsActive(isActive bool) {
	b.IsActive = isActive
}

// SetIsActive
func (b *BiZoneWatcher) SetAnyIsActive(a any) {
	if v, ok := a.(bool); ok {
		b.IsActive = v
	}
}

// ToStringBeautiful форматированный вывод
func (b *BiZoneWatcher) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, b.ID))
	str.WriteString(fmt.Sprintf("%s'username': '%s'\n", ws, b.Username))
	str.WriteString(fmt.Sprintf("%s'firstName': '%s'\n", ws, b.FirstName))
	str.WriteString(fmt.Sprintf("%s'lastName': '%s'\n", ws, b.LastName))
	str.WriteString(fmt.Sprintf("%s'email': '%s'\n", ws, b.Email))
	str.WriteString(fmt.Sprintf("%s'patronimic': '%s'\n", ws, b.Patronimic))
	str.WriteString(fmt.Sprintf("%s'isActive': '%t'\n", ws, b.IsActive))

	return str.String()
}
