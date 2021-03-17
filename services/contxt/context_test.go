package contxt

import "testing"

// go test -v -run CancelContext context.go context_test.go
func Test_CancelContext(t *testing.T) {
	CancelContext()
}

// go test -v -run ContextWithDeadline context.go context_test.go
func Test_ContextWithDeadline(t *testing.T) {
	ContextWithDeadline()
}

// go test -v -run ContextWithTimeout context.go context_test.go
func Test_ContextWithTimeout(t *testing.T) {
	ContextWithTimeout()
}

// go test -v -run ContextWithValue context.go context_test.go
func Test_ContextWithValue(t *testing.T) {
	ContextWithValue()
}
