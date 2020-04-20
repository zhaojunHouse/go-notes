package loop

import "testing"

// go test -v -run=SelectRole1 select.go select_test.go
func Test_SelectRole1(t *testing.T) {
	SelectRole1()
}

// go test -v -run=SelectBreak select.go select_test.go
func Test_SelectBreak(t *testing.T) {
	SelectBreak()
}

// go test -v -run=GotoLoopHelper select.go select_test.go
func Test_GotoLoopHelper(t *testing.T) {
	GotoLoopHelper()
}

// go test -v -run=BreakLoopHelper select.go select_test.go
func Test_BreakLoopHelper(t *testing.T) {
	BreakLoopHelper()
}
