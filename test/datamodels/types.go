package datamodels

type TestOptions struct {
	ValueSliceString []string
	ValueSliceBool   []bool
	ValueSliceInt    []int
	ValueString      string
	ValueUint64      uint64
	ValueInt         int
	ValueBool        bool
	SetFunc          func()
	GetFunc          func()
}
