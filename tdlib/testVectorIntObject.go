// AUTOGENERATED - DO NOT EDIT

package tdlib

// TestVectorIntObject A simple object containing a vector of objects that hold a number; for testing only
type TestVectorIntObject struct {
	tdCommon
	Value []TestInt `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorIntObject
func (testVectorIntObject *TestVectorIntObject) MessageType() string {
	return "testVectorIntObject"
}

// NewTestVectorIntObject creates a new TestVectorIntObject
//
// @param value Vector of objects
func NewTestVectorIntObject(value []TestInt) *TestVectorIntObject {
	testVectorIntObjectTemp := TestVectorIntObject{
		tdCommon: tdCommon{Type: "testVectorIntObject"},
		Value:    value,
	}

	return &testVectorIntObjectTemp
}
