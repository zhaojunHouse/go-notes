package testmethod

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// User user
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// IUser user interface
type IUser interface {
	GetUserInfo(userID int64) (*User, error)
	UpdateUser(user *User) error
	AddUser(user *User) error
}

// Routes router
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON send json
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name string
		Age  int32
	}{
		Name: "张三",
		Age:  28,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)
}

// MockServer mock server
func MockServer() *httptest.Server {
	//API调用处理函数
	sendJSON := func(rw http.ResponseWriter, r *http.Request) {
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
	return httptest.NewServer(http.HandlerFunc(sendJSON))
}
