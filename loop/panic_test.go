package loop

import "testing"

// go test -v -run=PanicRecover panic.go panic_test.go
func Test_PanicRecover(t *testing.T){
	PanicRecover()
}
