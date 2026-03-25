package datamodels_test

import (
	"fmt"
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

	// ---- Data ----
	size = 11
	allIPHomes := make([]string, 0, size)
	for range size {
		allIPHomes = append(allIPHomes, gofakeit.IPv4Address())
	}

	snortIds := make([]uint64, 0, size+2)
	for range size {
		snortIds = append(snortIds, gofakeit.Uint64())
	}

	allSensors := make([]uint64, 0, size+1)
	for range size {
		allSensors = append(allSensors, gofakeit.Uint64())
	}

	dataExample := datamodels.NewBiZoneIRPData()
	dataExample.SetAnyIPHome(allIPHomes)
	dataExample.SetAnySnortSid(snortIds)
	dataExample.SetAnyAllSensor(allSensors)
	dataExample.SetAnyID(gofakeit.ID())
	dataExample.SetAnyTitle(gofakeit.BookTitle())
	dataExample.SetAnyIPHome(gofakeit.IPv4Address())
	dataExample.SetAnyURLFTP(gofakeit.URL())
	dataExample.SetAnyIPExter(gofakeit.IPv4Address())
	dataExample.SetAnyURLHTTP(gofakeit.URL())
	dataExample.SetAnyEventType(gofakeit.BankType())
	dataExample.SetAnyURLArkime(gofakeit.URL())
	dataExample.SetAnySensor(gofakeit.Uint64())
	dataExample.SetAnyAllIPExt(gofakeit.Uint64())
	dataExample.SetAnyResponseTeam(gofakeit.Uint64())

	listTesting["Data"] = datamodelstest.TestOptions{
		ValueAny: *dataExample,
		SetFunc: func() {
			biZoneIRPAlert.SetData(*dataExample)
		},
		GetFunc: func() {
			data, ok := listTesting["Data"].ValueAny.(datamodels.BiZoneIRPData)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				data,
				biZoneIRPAlert.GetData(),
				func(a, b datamodels.BiZoneIRPData) bool {
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

	var num int
	for k, v := range listTesting {
		num++
		t.Run(fmt.Sprintf("Test %d. Field %s", num, k), func(t *testing.T) {
			v.SetFunc()
			v.GetFunc()
		})
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
