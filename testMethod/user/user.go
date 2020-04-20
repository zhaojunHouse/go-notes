package user

import (
	"go-notes/testMethod/person"
	"fmt"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) (int32, error) {
	male, err := u.Person.Get(id)
	if err != nil {
		fmt.Println(err)
	}
	return male.Male, err
}
