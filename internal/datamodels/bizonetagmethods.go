package datamodels

import (
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// ************* Структура BiZoneTag *************
// ***********************************************

// NewBiZoneTag новый объект
func NewBiZoneTag() *BiZoneTag {
	return &BiZoneTag{
		Created: time.Now().Format(time.RFC3339),
	}
}

// GetCreated для поля Created (формат RFC3339)
func (t *BiZoneTag) GetCreated() string {
	return t.Created
}

// SetCreated для поля Created (формат RFC3339)
func (t *BiZoneTag) SetCreated(created string) {
	t.Created = created
}

// SetAnyCreated для поля Created (формат RFC3339)
func (t *BiZoneTag) SetAnyCreated(i any) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	t.Created = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// GetName для поля Name
func (t *BiZoneTag) GetName() string {
	return t.Name
}

// SetName для поля Name
func (t *BiZoneTag) SetName(name string) {
	t.Name = name
}

// SetAnyName для поля Name
func (t *BiZoneTag) SetAnyName(a any) {
	t.Name = fmt.Sprint(a)
}

// GetColor для поля Color
func (t *BiZoneTag) GetColor() string {
	return t.Color
}

// SetColor для поля Color
func (t *BiZoneTag) SetColor(color string) {
	t.Color = color
}

// SetAnyColor для поля Color
func (t *BiZoneTag) SetAnyColor(a any) {
	t.Color = fmt.Sprint(a)
}

// GetCreatedByID для поля CreatedBy.ID
func (t *BiZoneTag) GetCreatedByID() uint64 {
	return t.CreatedBy.ID
}

// SetCreatedByID для поля CreatedBy.ID
func (t *BiZoneTag) SetCreatedByID(id uint64) {
	t.CreatedBy.ID = id
}

// SetAnyCreatedByID для поля CreatedBy.ID
func (t *BiZoneTag) SetAnyCreatedByID(a any) {
	if v, ok := a.(float32); ok {
		t.CreatedBy.ID = uint64(v)

		return
	}

	if v, ok := a.(float64); ok {
		t.CreatedBy.ID = uint64(v)

		return
	}

	if v, ok := a.(uint64); ok {
		t.CreatedBy.ID = v
	}
}

// GetCreatedByUsername для поля CreatedBy.Username
func (t *BiZoneTag) GetCreatedByUsername() string {
	return t.CreatedBy.Username
}

// SetCreatedByUsername для поля CreatedBy.Username
func (t *BiZoneTag) SetCreatedByUsername(username string) {
	t.CreatedBy.Username = username
}

// SetAnyCreatedByUsername для поля CreatedBy.Username
func (t *BiZoneTag) SetAnyCreatedByUsername(a any) {
	t.CreatedBy.Username = fmt.Sprint(a)
}

// ToStringBeautiful форматированный вывод
func (t *BiZoneTag) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, t.Name))
	str.WriteString(fmt.Sprintf("%s'color': '%s'\n", ws, t.Color))
	str.WriteString(fmt.Sprintf("%s'created': '%s'\n", ws, t.Created))
	str.WriteString(fmt.Sprintf("%s'created_by.id': '%d'\n", ws, t.CreatedBy.ID))
	str.WriteString(fmt.Sprintf("%s'created_by.username': '%s'\n", ws, t.CreatedBy.Username))

	return str.String()
}
