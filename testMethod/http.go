package testMethod

import (
	"net/http"
	"encoding/json"
	"net/http/httptest"
)

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int32 `json:"age"`
}

type IUser interface {
	GetUserInfo(userID int64)(*User, error)
	UpdateUser(user *User) error
	AddUser(user *User) error
}


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