package datamodels

import (
	"fmt"
	"slices"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneMITRECOV() *BiZoneMITRECOV {
	return &BiZoneMITRECOV{
		Persistence: make([]string, 0),
	}
}

// GetPersistence
func (mcov *BiZoneMITRECOV) GetPersistence() []string {
	if mcov.Persistence == nil {
		mcov.Persistence = make([]string, 0)
	}

	return mcov.Persistence
}

// SetPersistences список значений
func (mcov *BiZoneMITRECOV) SetPersistencies(persistences []string) {
	mcov.Persistence = persistences
}

// AddPersistence добавляет значение в список
func (mcov *BiZoneMITRECOV) AddPersistence(persistence string) {
	mcov.GetPersistence()

	if slices.Contains(mcov.Persistence, persistence) {
		return
	}

	mcov.Persistence = append(mcov.Persistence, persistence)
}

// AddAnyPersistence добавляет значение в список
func (mcov *BiZoneMITRECOV) AddAnyPersistence(a any) {
	mcov.AddPersistence(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (mcov *BiZoneMITRECOV) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	return fmt.Sprintf("%s'categores': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, mcov.Persistence))
}
