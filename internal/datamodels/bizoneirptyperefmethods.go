package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// GetID уникальный идентификатор
func (b *BiZoneIRPTypeRef) GetID() string {
	return b.ID
}

// SetID уникальный идентификатор
func (b *BiZoneIRPTypeRef) SetID(id string) {
	b.ID = id
}

// SetAnyID уникальный идентификатор
func (b *BiZoneIRPTypeRef) SetAnyID(a any) {
	b.ID = fmt.Sprint(a)
}

// GetTitle заголовок
func (b *BiZoneIRPTypeRef) GetTitle() string {
	return b.Title
}

// SetTitle заголовок
func (b *BiZoneIRPTypeRef) SetTitle(title string) {
	b.Title = title
}

// SetAnyTitle заголовок
func (b *BiZoneIRPTypeRef) SetAnyTitle(a any) {
	b.Title = fmt.Sprint(a)
}

// GetDescription описание
func (b *BiZoneIRPTypeRef) GetDescription() string {
	return b.Description
}

// SetDescription описание
func (b *BiZoneIRPTypeRef) SetDescription(description string) {
	b.Description = description
}

// SetAnyDescription описание
func (b *BiZoneIRPTypeRef) SetAnyDescription(a any) {
	b.Description = fmt.Sprint(a)
}

// ToStringBeautiful форматированный вывод
func (tr *BiZoneIRPTypeRef) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, tr.ID))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, tr.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, tr.Description))

	return str.String()
}
