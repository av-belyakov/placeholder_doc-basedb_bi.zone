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
		keyString string
		valueAny  any
	)

	biZoneIRPCase := &datamodels.VerifiedIRPBiZoneCase{}

	/*
		   type TestOptions struct {
			ValueMapAny                  map[string]any
			ValueSliceBiZoneObservedIocs []datamodels.BiZoneObservedIocs
			ValueSliceBiZoneTypeRef      []datamodels.BiZoneTypeRef
			ValueSliceWatchers           []datamodels.BiZoneWatcher
			ValueSliceAny                []any
			ValueSliceString             []string
			ValueSliceUint64             []uint64
			ValueSliceInt                []int
			ValueSliceBool               []bool
			ValueString                  string
			ValueUint64                  uint64
			ValueInt                     int
			ValueBool                    bool
			SetFunc                      func()
			GetFunc                      func()
		   }
	*/

	listTesting := map[string]datamodelstest.TestOptions{}

	// ---- GossopkaErrors ----
	gossopkaErrorsExample := make(map[string]any, 10)
	for range 10 {
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
	observedIocsExample := make([]datamodels.BiZoneIRPObservedIocs, 10)

	for range 10 {
		observedIoc := datamodels.NewBiZoneIRPObservedIocs()
		observedIoc.SetCategories(gofakeit.NiceColors())
		observedIoc.SetAnyIoc(gofakeit.Color())
		observedIoc.SetAnyIocType(gofakeit.CarType())

		observedIocsExample = append(observedIocsExample, *observedIoc)
	}

	listTesting["ObservedIocs"] = datamodelstest.TestOptions{
		ValueSliceBiZoneObservedIocs: observedIocsExample,
		SetFunc: func() {
			biZoneIRPCase.SetObservedIocs(listTesting["ObservedIocs"].ValueSliceBiZoneObservedIocs)
		},
		GetFunc: func() {
			assert.True(t, slices.EqualFunc(
				listTesting["ObservedIocs"].ValueSliceBiZoneObservedIocs,
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

	/*
		listTesting[""] = datamodelstest.TestOptions{
			ValueString: ,
			SetFunc: func() {
				biZoneCase.(listTesting[""].ValueString)
			},
			GetFunc: func() {
				assert.Equal(t, biZoneCase.(), listTesting[""].ValueString)
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
