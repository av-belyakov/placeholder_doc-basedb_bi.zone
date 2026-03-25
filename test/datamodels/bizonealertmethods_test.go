package datamodels_test

import (
	"slices"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodelstest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneIRPAlertMethods(t *testing.T) {
	var (
		size int
	)

	biZoneIRPAlert := &datamodels.VerifiedBiZoneIRPAlert{}
	listTesting := map[string]datamodelstest.TestOptions{}

	// ---- Snapshots ----
	size = 13
	ipAddresses := make([]string, 0, size)
	for range size {
		ipAddresses = append(ipAddresses, gofakeit.IPv4Address())
	}
	macAddresses := make([]string, 0, size)
	for range size {
		macAddresses = append(macAddresses, gofakeit.MacAddress())
	}

	snapshotsExample := make([]datamodels.BiZoneIRPSnapshot, 0, size)
	for range size {
		snapshot := datamodels.NewBiZoneIRPSnapshot()
		snapshot.SetAnyIPAddresse(ipAddresses)
		snapshot.SetAnyMACAddresse(macAddresses)
		snapshot.SetAnyOS(gofakeit.BankName())
		snapshot.SetAnyOSType(gofakeit.BankType())
		snapshot.SetAnyFqdn(gofakeit.DomainName())
		snapshot.SetAnyDomain(gofakeit.DomainName())
		snapshot.SetAnyCMDBID(gofakeit.CarModel())
		snapshot.SetAnyHostname(gofakeit.DomainName())
		snapshot.SetAnyUserCMDBName(gofakeit.DomainSuffix())

		snapshotsExample = append(snapshotsExample, *snapshot)
	}

	listTesting["Snapshots"] = datamodelstest.TestOptions{
		ValueAny: snapshotsExample,
		SetFunc: func() {
			biZoneIRPAlert.SetSnapshots(snapshotsExample)
		},
		GetFunc: func() {
			snapshots, ok := listTesting["Snapshots"].ValueAny.([]datamodels.BiZoneIRPSnapshot)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				snapshots,
				biZoneIRPAlert.GetSnapshots(),
				func(a, b datamodels.BiZoneIRPSnapshot) bool {
					if !slices.Equal(a.IPAddresses, b.IPAddresses) {
						return false
					}

					if !slices.Equal(a.MACAddresses, b.MACAddresses) {
						return false
					}

					if a.GetOS() != b.GetOS() {
						return false
					}

					if a.GetOSType() != b.GetOSType() {
						return false
					}

					if a.GetFqdn() != b.GetFqdn() {
						return false
					}

					if a.GetDomain() != b.GetDomain() {
						return false
					}

					if a.GetCMDBID() != b.GetCMDBID() {
						return false
					}
					if a.GetHostname() != b.GetHostname() {
						return false
					}
					return a.GetUserCMDBName() == b.GetUserCMDBName()
				}))
		},
	}

	// ---- Tags ----
	size = 14
	tagsExample := make([]datamodels.BiZoneIRPTag, 0, size)
	for range size {
		tag := datamodels.NewBiZoneIRPTag()
		tag.SetAnyName(gofakeit.EmojiTag())
		tag.SetAnyColor(gofakeit.Color())

		tagsExample = append(tagsExample, *tag)
	}

	listTesting["Tags"] = datamodelstest.TestOptions{
		ValueAny: snapshotsExample,
		SetFunc: func() {
			biZoneIRPAlert.SetSnapshots(snapshotsExample)
		},
		GetFunc: func() {
			snapshots, ok := listTesting["Tags"].ValueAny.([]datamodels.BiZoneIRPSnapshot)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				snapshots,
				biZoneIRPAlert.GetSnapshots(),
				func(a, b datamodels.BiZoneIRPSnapshot) bool {
					if !slices.Equal(a.IPAddresses, b.IPAddresses) {
						return false
					}

					if !slices.Equal(a.MACAddresses, b.MACAddresses) {
						return false
					}

					if a.GetOS() != b.GetOS() {
						return false
					}

					if a.GetOSType() != b.GetOSType() {
						return false
					}

					if a.GetFqdn() != b.GetFqdn() {
						return false
					}

					if a.GetDomain() != b.GetDomain() {
						return false
					}

					if a.GetCMDBID() != b.GetCMDBID() {
						return false
					}
					if a.GetHostname() != b.GetHostname() {
						return false
					}
					return a.GetUserCMDBName() == b.GetUserCMDBName()
				}))
		},
	}

	/*
		// ----  ----
		listTesting[""] = datamodelstest.TestOptions{
			ValueAny: gofakeit.TimeZoneRegion(),
			SetFunc: func() {
				biZoneIRPCase.SetAny(listTesting[""].ValueAny)
			},
			GetFunc: func() {
				assert.Equal(t, biZoneIRPCase.Get(), listTesting[""].ValueAny)
			},
		}


		listTesting[""] = datamodelstest.TestOptions{
			ValueString: ,
			SetFunc: func() {
				biZoneIRPCase.(listTesting[""].ValueString)
			},
			GetFunc: func() {
				assert.Equal(t, biZoneIRPCase.(), listTesting[""].ValueString)
			},
		}
	*/
}
