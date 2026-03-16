package datamodels_test

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodelstest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneCaseMethods(t *testing.T) {
	var (
		size      int
		keyString string
		valueAny  any
	)

	biZoneIRPCase := &datamodels.VerifiedIRPBiZoneCase{}
	listTesting := map[string]datamodelstest.TestOptions{}

	// ---- GossopkaErrors ----
	size = 10
	gossopkaErrorsExample := make(map[string]any, size)
	for range size {
		gossopkaErrorsExample[gofakeit.Error().Error()] = gofakeit.ErrorDatabase()
	}

	listTesting["GossopkaErrors"] = datamodelstest.TestOptions{
		ValueMapAny: gossopkaErrorsExample,
		SetFunc: func() {
			biZoneIRPCase.SetGossopkaErrors(listTesting["GossopkaErrors"].ValueMapAny)

			keyString = gofakeit.Error().Error()
			valueAny = gofakeit.ErrorDatabase()
			assert.True(t, biZoneIRPCase.SetAnyGossopkaErrors(keyString, valueAny))

			listTesting["GossopkaErrors"].ValueMapAny[keyString] = valueAny
		},
		GetFunc: func() {
			assert.True(t, maps.EqualFunc(
				listTesting["GossopkaErrors"].ValueMapAny,
				biZoneIRPCase.GetGossopkaErrors(),
				func(a1 any, a2 any) bool {
					v1, ok := a1.(error)
					if !ok {
						return false
					}

					v2, ok := a2.(error)
					if !ok {
						return false
					}

					return strings.EqualFold(v1.Error(), v2.Error())
				}))
		},
	}

	// ---- ObservedIocs ----
	size = 10
	observedIocsExample := make([]datamodels.BiZoneIRPObservedIocs, 0, size)

	for range size {
		observedIoc := datamodels.NewBiZoneIRPObservedIocs()
		observedIoc.SetCategories(gofakeit.NiceColors())
		observedIoc.SetAnyIoc(gofakeit.Color())
		observedIoc.SetAnyIocType(gofakeit.CarType())

		observedIocsExample = append(observedIocsExample, *observedIoc)
	}

	listTesting["ObservedIocs"] = datamodelstest.TestOptions{
		ValueAny: observedIocsExample,
		SetFunc: func() {
			biZoneIRPCase.SetObservedIocs(observedIocsExample)
		},
		GetFunc: func() {
			observedIocs, ok := listTesting["ObservedIocs"].ValueAny.([]datamodels.BiZoneIRPObservedIocs)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				observedIocs,
				biZoneIRPCase.GetObservedIocs(),
				func(a, b datamodels.BiZoneIRPObservedIocs) bool {
					if a.GetIoc() != b.GetIoc() {
						return false
					}

					if a.GetIocType() != b.GetIocType() {
						return false
					}

					return slices.Equal(a.GetCategories(), b.GetCategories())
				}))
		},
	}

	// ---- SecondaryCategoryRef ----
	size = 13
	secondaryCategoryRefExample := make([]datamodels.BiZoneIRPTypeRef, 0, size)
	for range size {
		secondaryCategoryRef := datamodels.NewBiZoneIRPTypeRef()
		secondaryCategoryRef.ID = gofakeit.ID()
		secondaryCategoryRef.Title = gofakeit.JobTitle()
		secondaryCategoryRef.Description = gofakeit.JobDescriptor()

		secondaryCategoryRefExample = append(secondaryCategoryRefExample, *secondaryCategoryRef)
	}

	listTesting["SecondaryCategoryRef"] = datamodelstest.TestOptions{
		ValueAny: secondaryCategoryRefExample,
		SetFunc: func() {
			biZoneIRPCase.SetSecondaryCategoryRef(secondaryCategoryRefExample)
		},
		GetFunc: func() {
			secondaryCategoryRefs, ok := listTesting["SecondaryCategoryRef"].ValueAny.([]datamodels.BiZoneIRPTypeRef)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				secondaryCategoryRefs,
				biZoneIRPCase.GetSecondaryCategoryRef(),
				func(a, b datamodels.BiZoneIRPTypeRef) bool {
					if a.GetId() != b.GetId() {
						return false
					}

					if a.GetTitle() != b.GetTitle() {
						return false
					}

					return a.GetDescription() == b.GetDescription()
				}))
		},
	}

	// ---- Watcher ----
	size = 13
	watchersExample := make([]datamodels.BiZoneIRPWatcher, 0, size)
	for range size {
		watcherExample := datamodels.NewBiZoneIRPWatcher()
		watcherExample.ID = gofakeit.ID()
		watcherExample.Username = gofakeit.Name()
		watcherExample.FirstName = gofakeit.FirstName()
		watcherExample.LastName = gofakeit.LastName()
		watcherExample.Email = gofakeit.Email()
		watcherExample.Patronimic = gofakeit.HipsterParagraph()
		watcherExample.IsActive = gofakeit.Bool()

		watchersExample = append(watchersExample, *watcherExample)
	}

	listTesting["Watchers"] = datamodelstest.TestOptions{
		ValueAny: watchersExample,
		SetFunc: func() {
			biZoneIRPCase.SetWatchers(watchersExample)
		},
		GetFunc: func() {
			watchers, ok := listTesting["Watchers"].ValueAny.([]datamodels.BiZoneIRPWatcher)
			assert.True(t, ok)

			assert.True(t, slices.EqualFunc(
				watchers,
				biZoneIRPCase.GetWatchers(),
				func(a, b datamodels.BiZoneIRPWatcher) bool {
					if a.GetId() != b.GetId() {
						return false
					}

					if a.GetUsername() != b.GetUsername() {
						return false
					}

					if a.GetFirstName() != b.GetFirstName() {
						return false
					}

					if a.GetLastName() != b.GetLastName() {
						return false
					}

					if a.GetEmail() != b.GetEmail() {
						return false
					}

					if a.GetPatronimic() != b.GetPatronimic() {
						return false
					}

					return a.GetIsActive() == b.GetIsActive()
				}))
		},
	}

	// ---- ActivityDetected ----
	size = 12
	activityDetectedExp := make([]any, 0, size)
	for range size {
		activityDetectedExp = append(activityDetectedExp, DataStorage[string]{Data: gofakeit.APIUserAgent()})
	}

	listTesting["ActivityDetected"] = datamodelstest.TestOptions{
		ValueAny: activityDetectedExp,
		SetFunc: func() {
			biZoneIRPCase.SetActivityDetected(activityDetectedExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetActivityDetected(), listTesting["ActivityDetected"].ValueAny)
		},
	}

	// ---- Attachments ----
	size = 15
	attachmentsExp := make([]any, 0, size)
	for range size {
		attachmentsExp = append(attachmentsExp, DataStorage[string]{Data: gofakeit.CarModel()})
	}

	listTesting["Attachments"] = datamodelstest.TestOptions{
		ValueAny: attachmentsExp,
		SetFunc: func() {
			biZoneIRPCase.SetAttachments(attachmentsExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetAttachments(), listTesting["Attachments"].ValueAny)
		},
	}

	// ---- Badges ----
	size = 11
	badgesExp := make([]any, 0, size)
	for range size {
		badgesExp = append(badgesExp, DataStorage[string]{Data: gofakeit.City()})
	}

	listTesting["Badges"] = datamodelstest.TestOptions{
		ValueAny: badgesExp,
		SetFunc: func() {
			biZoneIRPCase.SetBadges(badgesExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetBadges(), listTesting["Badges"].ValueAny)
		},
	}

	// ---- Emls ----
	size = 12
	emlsExp := make([]any, 0, size)
	for range size {
		emlsExp = append(emlsExp, DataStorage[string]{Data: gofakeit.TimeZone()})
	}

	listTesting["Emls"] = datamodelstest.TestOptions{
		ValueAny: emlsExp,
		SetFunc: func() {
			biZoneIRPCase.SetEmls(emlsExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetEmls(), listTesting["Emls"].ValueAny)
		},
	}

	// ---- Slas ----
	size = 10
	slasExp := make([]any, 0, size)
	for range size {
		slasExp = append(slasExp, DataStorage[string]{Data: gofakeit.NounCommon()})
	}

	listTesting["Slas"] = datamodelstest.TestOptions{
		ValueAny: slasExp,
		SetFunc: func() {
			biZoneIRPCase.SetSlas(slasExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetSlas(), listTesting["Slas"].ValueAny)
		},
	}

	// ---- Tags ----
	size = 13
	tagsExp := make([]any, 0, size)
	for range size {
		tagsExp = append(tagsExp, DataStorage[string]{Data: gofakeit.EmojiTag()})
	}

	listTesting["Tags"] = datamodelstest.TestOptions{
		ValueAny: tagsExp,
		SetFunc: func() {
			biZoneIRPCase.SetTags(tagsExp)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetTags(), listTesting["Tags"].ValueAny)
		},
	}

	// ---- KeyField ----
	size = 12
	keyField := make([]any, 0, size)
	for range size {
		keyField = append(keyField, DataStorage[string]{Data: gofakeit.AnimalType()})
	}

	listTesting["KeyField"] = datamodelstest.TestOptions{
		ValueAny: keyField,
		SetFunc: func() {
			biZoneIRPCase.SetKeyField(keyField)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetKeyField(), listTesting["KeyField"].ValueAny)
		},
	}

	// ---- MITRECOV ----
	size = 11
	mitreCov := datamodels.NewBiZoneMITRECOV()
	for range size {
		mitreCov.AddAnyPersistence(gofakeit.AppAuthor())
	}

	listTesting["MITRECOV"] = datamodelstest.TestOptions{
		ValueAny: *mitreCov,
		SetFunc: func() {
			biZoneIRPCase.SetMITRECOV(*mitreCov)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneIRPCase.GetMITRECOV(), listTesting["MITRECOV"].ValueAny)
		},
	}

	/*
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

	var num int
	for k, v := range listTesting {
		num++
		t.Run(fmt.Sprintf("Test %d. Field %s", num, k), func(t *testing.T) {
			v.SetFunc()
			v.GetFunc()
		})
	}
}

type DataStorage[T comparable] struct {
	Data T
}

func (storage *DataStorage[T]) Compare(a []DataStorage[T]) bool {
	for _, v := range a {
		if v.Data == storage.Data {
			return true
		}
	}

	return false
}
