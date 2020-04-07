package goroutine

import "testing"

//  go test -v -run WaitingComplete groutine.go groutine_test.go
func Test_WaitingComplete(t *testing.T){
	ChannelWaitingCompleted()
}

// go test -v -run ChannelSum groutine.go groutine_test.go
func Test_ChannelSum(t *testing.T){
	ChannelSum()
}

// go test -v -run ChannelSelect groutine.go groutine_test.go
func Test_ChannelSelect(t *testing.T){
	ChannelSelect()
}

// go test -v -run ChannelBuffer groutine.go groutine_test.go
func Test_ChannelBuffer(t *testing.T){
	ChannelBuffer()
}

// go test -v -run ChannelNoBuffer groutine.go groutine_test.go
func Test_ChannelNoBuffer(t *testing.T){
	ChannelNoBuffer()
}

// go test -v -run ChannelClose groutine.go groutine_test.go
func Test_ChannelClose(t *testing.T){
	ChannelClose()
}