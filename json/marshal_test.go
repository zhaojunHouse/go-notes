package json

import "testing"

// go test -v -run=MarshalHelper *.go
func Test_MarshalHelper(t *testing.T){
	MarshalHelper()
}

// go test -v -run=StructJson *.go
func Test_StructJson(t *testing.T){
	StructJson()
}