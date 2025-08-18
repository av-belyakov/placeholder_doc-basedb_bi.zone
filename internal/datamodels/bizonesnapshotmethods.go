package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// ********** Структура Snapshot ***********
// *****************************************

// NewBiZneSnapshot новый объект
func NewBiZneSnapshot() *BiZoneSnapshot {
	return &BiZoneSnapshot{}
}

// GetIPAddresses для поля IPAddresses
func (s *BiZoneSnapshot) GetIPAddresses() []string {
	return s.IPAddresses
}

// SetIPAddresses для поля IPAddresses
func (s *BiZoneSnapshot) SetIPAddresses(ipAddress []string) {
	s.IPAddresses = ipAddress
}

// SetIPAddresse добавляет ip адрес в поле IPAddresses
func (s *BiZoneSnapshot) SetIPAddresse(ipAddress string) {
	if s.IPAddresses == nil {
		s.IPAddresses = []string(nil)
	}

	s.IPAddresses = append(s.IPAddresses, ipAddress)
}

// SetAnyIPAddresse добавляет ip адрес в поле IPAddresses
func (s *BiZoneSnapshot) SetAnyIPAddresse(a any) {
	if s.IPAddresses == nil {
		s.IPAddresses = []string(nil)
	}

	s.IPAddresses = append(s.IPAddresses, fmt.Sprint(a))
}

// GetMACAddresses для поля MACAddresses
func (s *BiZoneSnapshot) GetMACAddresses() []string {
	return s.IPAddresses
}

// SetMACAddresses для поля MACAddresses
func (s *BiZoneSnapshot) SetMACAddresses(macAddress []string) {
	s.MACAddresses = macAddress
}

// SetMACAddresse добавляет mac адрес в поле MACAddresses
func (s *BiZoneSnapshot) SetMACAddresse(macAddress string) {
	if s.MACAddresses == nil {
		s.MACAddresses = []string(nil)
	}

	s.MACAddresses = append(s.MACAddresses, macAddress)
}

// SetAnyMACAddresse добавляет mac адрес в поле MACAddresses
func (s *BiZoneSnapshot) SetAnyMACAddresse(a any) {
	if s.MACAddresses == nil {
		s.MACAddresses = []string(nil)
	}

	s.MACAddresses = append(s.MACAddresses, fmt.Sprint(a))
}

// GetDomain для поля Domain
func (s *BiZoneSnapshot) GetDomain() string {
	return s.Domain
}

// SetDomain для поля Domain
func (s *BiZoneSnapshot) SetDomain(domain string) {
	s.Domain = domain
}

// SetAnyDomain для поля Domain
func (s *BiZoneSnapshot) SetAnyDomain(a any) {
	s.Domain = fmt.Sprint(a)
}

// GetFqdn для поля Fqdn
func (s *BiZoneSnapshot) GetFqdn() string {
	return s.Fqdn
}

// SetFqdn для поля Fqdn
func (s *BiZoneSnapshot) SetFqdn(fqdn string) {
	s.Fqdn = fqdn
}

// SetAnyFqdn для поля Fqdn
func (s *BiZoneSnapshot) SetAnyFqdn(a any) {
	s.Fqdn = fmt.Sprint(a)
}

// GetOS для поля OS
func (s *BiZoneSnapshot) GetOS() string {
	return s.OS
}

// SetOS для поля OS
func (s *BiZoneSnapshot) SetOS(value string) {
	s.OS = value
}

// SetAnyOS для поля OS
func (s *BiZoneSnapshot) SetAnyOS(a any) {
	s.OS = fmt.Sprint(a)
}

// GetUserCMDBName для поля UserCMDBName
func (s *BiZoneSnapshot) GetUserCMDBName() string {
	return s.UserCMDBName
}

// SetUserCMDBName для поля UserCMDBName
func (s *BiZoneSnapshot) SetUserCMDBName(value string) {
	s.UserCMDBName = value
}

// SetAnyUserCMDBName для поля UserCMDBName
func (s *BiZoneSnapshot) SetAnyUserCMDBName(a any) {
	s.UserCMDBName = fmt.Sprint(a)
}

// GetCMDBID для поля CMDBID
func (s *BiZoneSnapshot) GetCMDBID() string {
	return s.CMDBID
}

// SetCMDBID для поля CMDBID
func (s *BiZoneSnapshot) SetCMDBID(value string) {
	s.CMDBID = value
}

// SetAnyCMDBID для поля CMDBID
func (s *BiZoneSnapshot) SetAnyCMDBID(a any) {
	s.CMDBID = fmt.Sprint(a)
}

// GetHostname для поля Hostname
func (s *BiZoneSnapshot) GetHostname() string {
	return s.Hostname
}

// SetHostname для поля Hostname
func (s *BiZoneSnapshot) SetHostname(hostname string) {
	s.Hostname = hostname
}

// SetAnyHostname для поля Hostname
func (s *BiZoneSnapshot) SetAnyHostname(a any) {
	s.Hostname = fmt.Sprint(a)
}

// GetOSType для поля OSType
func (s *BiZoneSnapshot) GetOSType() string {
	return s.OSType
}

// SetOSType для поля OSType
func (s *BiZoneSnapshot) SetOSType(osType string) {
	s.OSType = osType
}

// SetAnyOSType для поля OSType
func (s *BiZoneSnapshot) SetAnyOSType(a any) {
	s.OSType = fmt.Sprint(a)
}

// ToStringBeautiful форматированный вывод
func (s *BiZoneSnapshot) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'os': '%s'\n", ws, s.OS))
	str.WriteString(fmt.Sprintf("%s'fqdn': '%s'\n", ws, s.Fqdn))
	str.WriteString(fmt.Sprintf("%s'domain': '%s'\n", ws, s.Domain))
	str.WriteString(fmt.Sprintf("%s'cmdb_id': '%s'\n", ws, s.CMDBID))
	str.WriteString(fmt.Sprintf("%s'os_type': '%s'\n", ws, s.OSType))
	str.WriteString(fmt.Sprintf("%s'hostname': '%s'\n", ws, s.Hostname))
	str.WriteString(fmt.Sprintf("%s'user_cmdb_name': '%s'\n", ws, s.UserCMDBName))
	str.WriteString(fmt.Sprintf("%s'ip_addressess': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.IPAddresses)))
	str.WriteString(fmt.Sprintf("%s'mac_addressess': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, s.MACAddresses)))

	return str.String()
}
