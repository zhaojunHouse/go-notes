package json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int32 `json:"age"`
	Addr string `json:"addr"`
}

func marshal(){
	user1 := &User{
		ID:   1,
		Name: "zj",
	}
	b, _ :=json.Marshal(user1)
	var u User
	json.Unmarshal(b, u)
	fmt.Println(u)
}
