package datamodels

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneObservedIocs() *BiZoneObservedIocs {
	return &BiZoneObservedIocs{
		Category: make([]string, 0),
	}
}

// GetIoc сам Ioc
func (oiocs *BiZoneObservedIocs) GetIoc() string {
	return oiocs.Ioc
}

// SetIoc сам Ioc
func (oiocs *BiZoneObservedIocs) SetIoc(ioc string) {
	oiocs.Ioc = ioc
}

// SetAnyIoc сам Ioc
func (oiocs *BiZoneObservedIocs) SetAnyIoc(a any) {
	oiocs.Ioc = fmt.Sprint(a)
}

// GetIocType тип Ioc
func (oiocs *BiZoneObservedIocs) GetIocType() string {
	return oiocs.IocType
}

// SetIocType тип Ioc
func (oiocs *BiZoneObservedIocs) SetIocType(iocType string) {
	oiocs.IocType = iocType
}

// SetAnyIocType тип Ioc
func (oiocs *BiZoneObservedIocs) SetAnyIocType(a any) {
	oiocs.IocType = fmt.Sprint(a)
}

// GetCategories категории
func (oiocs *BiZoneObservedIocs) GetCategories() []string {
	if oiocs.Category == nil {
		oiocs.Category = make([]string, 0)
	}

	return oiocs.Category
}

// SetCategories категории
func (oiocs *BiZoneObservedIocs) SetCategories(categories []string) {
	oiocs.Category = categories
}

// AddCategory добавляет значение в список
func (oiocs *BiZoneObservedIocs) AddCategory(category string) {
	oiocs.GetCategories()

	if slices.Contains(oiocs.Category, category) {
		return
	}

	oiocs.Category = append(oiocs.Category, category)
}

// AddAnyCategory добавляет значение в список
func (oiocs *BiZoneObservedIocs) AddAnyCategory(a any) {
	oiocs.AddCategory(fmt.Sprint(a))
}

// ToStringBeautiful форматированный вывод
func (oiocs *BiZoneObservedIocs) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'ioc': '%s'\n", ws, oiocs.Ioc))
	str.WriteString(fmt.Sprintf("%s'iocType': '%s'\n", ws, oiocs.IocType))
	str.WriteString(fmt.Sprintf("%s'categores': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, oiocs.Category)))

	return str.String()
}
