package json

import "testing"

// go test -v -run=MarshalHelper *.go
func Test_MarshalHelper(t *testing.T) {
	MarshalHelper()
}

// go test -v -run=StructJSON *.go
func Test_StructJSON(t *testing.T) {
	StructJSON()
}
