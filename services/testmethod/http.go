package testmethod

import (
	"encoding/json"
	"fmt"
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

// GetUserInfo 获取用户信息
func GetUserInfo(userID int64) (User, error) {
	return User{
		ID:   1,
		Name: "里斯",
		Age:  22,
	}, nil
}

// SingleFunc 测试单个函数
func SingleFunc(rw http.ResponseWriter, r *http.Request) {
	u, err := GetUserInfo(1)
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)
}

// MockServer mock server
func MockServer() *httptest.Server {
	//适配器转换
	return httptest.NewServer(http.HandlerFunc(SingleFunc))
}
