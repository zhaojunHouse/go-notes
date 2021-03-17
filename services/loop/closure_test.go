package loop

import "testing"

// go test -v -run=ClosureBug1 *.go
func Test_ClosureBug1(t *testing.T) {
	ClosureBug1()
}

// go test -v -run=ClosureBug2 *.go
func Test_ClosureBug2(t *testing.T) {
	ClosureBug2()
}

// go test -v -run=ClosureDefer *.go
func Test_ClosureDefer(t *testing.T) {
	ClosureDefer()
}
