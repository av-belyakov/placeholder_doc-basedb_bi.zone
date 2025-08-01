package datamodels

import (
	"fmt"
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

// GetMACAddresses для поля MACAddresses
func (s *BiZoneSnapshot) GetMACAddresses() []string {
	return s.IPAddresses
}

// SetMACAddresses для поля MACAddresses
func (s *BiZoneSnapshot) SetMACAddresses(macAddress []string) {
	s.MACAddresses = macAddress
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
func (s *BiZoneSnapshot) SetAnyDomain(i any) {
	s.Domain = fmt.Sprint(i)
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
func (s *BiZoneSnapshot) SetAnyFqdn(i any) {
	s.Fqdn = fmt.Sprint(i)
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
func (s *BiZoneSnapshot) SetAnyOS(i any) {
	s.OS = fmt.Sprint(i)
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
func (s *BiZoneSnapshot) SetAnyUserCMDBName(i any) {
	s.UserCMDBName = fmt.Sprint(i)
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
func (s *BiZoneSnapshot) SetAnyCMDBID(i any) {
	s.CMDBID = fmt.Sprint(i)
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
func (s *BiZoneSnapshot) SetAnyHostname(i any) {
	s.Hostname = fmt.Sprint(i)
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
func (s *BiZoneSnapshot) SetAnyOSType(i any) {
	s.OSType = fmt.Sprint(i)
}
