package datamodels

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/datamodels"

type TestOptions struct {
	ValueMapAny                  map[string]any
	ValueSliceBiZoneObservedIocs []datamodels.BiZoneIRPObservedIocs
	//ValueSliceBiZoneTypeRef      []datamodels.BiZoneIRPTypeRef
	//ValueSliceWatchers           []datamodels.BiZoneIRPWatcher
	ValueSliceAny    []any
	ValueSliceString []string
	ValueSliceUint64 []uint64
	ValueSliceInt    []int
	ValueSliceBool   []bool
	ValueString      string
	ValueUint64      uint64
	ValueInt         int
	ValueBool        bool
	SetFunc          func()
	GetFunc          func()
}
