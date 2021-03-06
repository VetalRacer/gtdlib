// AUTOGENERATED - DO NOT EDIT

package tdlib

// TestVectorInt A simple object containing a vector of numbers; for testing only
type TestVectorInt struct {
	tdCommon
	Value []int32 `json:"value"` // Vector of numbers
}

// MessageType return the string telegram-type of TestVectorInt
func (testVectorInt *TestVectorInt) MessageType() string {
	return "testVectorInt"
}

// NewTestVectorInt creates a new TestVectorInt
//
// @param value Vector of numbers
func NewTestVectorInt(value []int32) *TestVectorInt {
	testVectorIntTemp := TestVectorInt{
		tdCommon: tdCommon{Type: "testVectorInt"},
		Value:    value,
	}

	return &testVectorIntTemp
}
