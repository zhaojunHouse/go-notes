package testMethod

import (
	"net/http"
	"encoding/json"
	"net/http/httptest"
)

func Routes(){
	http.HandleFunc("/sendjson",SendJSON)
}

func SendJSON(rw http.ResponseWriter,r *http.Request){
	u := struct {
		Name string
		Age int32
	}{
		Name:"张三",
		Age:28,
	}

	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)
}

func MockServer() *httptest.Server {
	//API调用处理函数
	sendJson := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(u)
	}
	//适配器转换
	return httptest.NewServer(http.HandlerFunc(sendJson))
}