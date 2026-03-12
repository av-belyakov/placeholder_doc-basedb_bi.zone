package datamodels_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"
	datamodelstest "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/test/datamodels"
)

func TestBiZoneCaseMethods(t testing.T) {
	biZoneCase := &datamodels.VerifiedIRPBiZoneCase{}

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

	listTesting["GossopkaErrors"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			biZoneCase.(listTesting.["GossopkaErrors"].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneCase.(), listTesting["GossopkaErrors"].ValueString)
		},
	}

	/*
	listTesting["GossopkaErrors"] = datamodelstest.TestOptions{
		ValueString: gofakeit.ID(),
		SetFunc: func() {
			biZoneCase.(listTesting[""].ValueString)
		},
		GetFunc: func() {
			assert.Equal(t, biZoneCase.(), listTesting[""].ValueString)
		},
	}
	*/

}
