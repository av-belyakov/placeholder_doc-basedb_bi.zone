package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// GetID уникальный идентификатор
func (b *BiZoneTypeRef) GetID() string {
	return b.ID
}

// SetID уникальный идентификатор
func (b *BiZoneTypeRef) SetID(id string) {
	b.ID = id
}

// SetAnyID уникальный идентификатор
func (b *BiZoneTypeRef) SetAnyID(a any) {
	b.ID = fmt.Sprint(a)
}

// GetTitle заголовок
func (b *BiZoneTypeRef) GetTitle() string {
	return b.Title
}

// SetTitle заголовок
func (b *BiZoneTypeRef) SetTitle(title string) {
	b.Title = title
}

// SetAnyTitle заголовок
func (b *BiZoneTypeRef) SetAnyTitle(a any) {
	b.Title = fmt.Sprint(a)
}

// GetDescription описание
func (b *BiZoneTypeRef) GetDescription() string {
	return b.Description
}

// SetDescription описание
func (b *BiZoneTypeRef) SetDescription(description string) {
	b.Description = description
}

// SetAnyDescription описание
func (b *BiZoneTypeRef) SetAnyDescription(a any) {
	b.Description = fmt.Sprint(a)
}

// ToStringBeautiful форматированный вывод
func (tr *BiZoneTypeRef) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, tr.ID))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, tr.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, tr.Description))

	return str.String()
}
