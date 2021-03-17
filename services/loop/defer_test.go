package loop

import "testing"

// go test -v -run=DeferHelper defer.go defer_test.go
func Test_DeferHelper(t *testing.T) {
	DeferHelper()
}

// go test -v -run=DeferUpdateFuncResult defer.go defer_test.go
func Test_DeferUpdateFuncResult(t *testing.T) {
	DeferUpdateFuncResult()
}

// go test -v -run=DeferPanic defer.go defer_test.go
func Test_DeferPanic(t *testing.T) {
	DeferPanic()
}

// go test -v -run=DeferBibao defer.go defer_test.go
func Test_DeferBibao(t *testing.T) {
	DeferBibao()
}
