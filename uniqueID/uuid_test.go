package uniqueID

import "testing"

// go test -v -run=CreateUUID uuid.go uuid_test.go
func Test_CreateUUID(t *testing.T){
	CreateUUID()
}

// go test -v -run=IDWorker uuid.go uuid_test.go
func Test_IDWorker(t *testing.T){
	IDWroker()
}
