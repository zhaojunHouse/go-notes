package testmethod

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	Routes()
}

// go test -v -run=SendJSON *.go
func Test_SendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建Request失败")
	}
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	log.Println("code:", rw.Code)

	log.Println("body:", rw.Body.String())
}

// go test -v -run=MockServer *.go
func Test_MockServer(t *testing.T) {
	//创建一个模拟的服务器
	server := MockServer()
	defer server.Close()
	//Get请求发往模拟服务器的地址
	resq, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("创建Get失败")
	}
	defer resq.Body.Close()

	log.Println("code:", resq.StatusCode)
	json, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", json)
}
