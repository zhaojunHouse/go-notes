package unicode

import "testing"

// go test -v -run=RuneCountHelper *.go
func Test_RuneCountHelper(t *testing.T) {
	t.Log(RuneCountHelper())
}

func Test_RuneCountInString(t *testing.T) {
	t.Log(RuneCountInString())
}
