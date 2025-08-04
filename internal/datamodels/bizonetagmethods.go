package datamodels

import (
	"fmt"
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
func (t *BiZoneTag) SetAnyName(i any) {
	t.Name = fmt.Sprint(i)
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
func (t *BiZoneTag) SetAnyColor(i any) {
	t.Color = fmt.Sprint(i)
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
func (t *BiZoneTag) SetAnyCreatedByUsername(i any) {
	t.CreatedBy.Username = fmt.Sprint(i)
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
func (t *BiZoneTag) SetAnyCreatedByID(i any) {
	if v, ok := i.(float32); ok {
		t.CreatedBy.ID = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		t.CreatedBy.ID = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		t.CreatedBy.ID = v
	}
}
