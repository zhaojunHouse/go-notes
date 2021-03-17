package http

import "testing"

// go test -v -run=HttpServerV1 *.go
func Test_HttpServerV1(t *testing.T) {
	HttpServerV1()
}

// go test -v -run=HttpServerV2 *.go
func Test_HttpServerV2(t *testing.T) {
	HttpServerV2()
}

// go test -v -run=HttpServerV3 *.go
func Test_HttpServerV3(t *testing.T) {
	HttpServerV3()
}
