package user

import (
	"fmt"
	"go-notes/testmethod/person"
)

// User user
type User struct {
	Person person.Male
}

// NewUser new user
func NewUser(p person.Male) *User {
	return &User{Person: p}
}

// GetUserInfo get user info
func (u *User) GetUserInfo(id int64) (int32, error) {
	male, err := u.Person.Get(id)
	if err != nil {
		fmt.Println(err)
	}
	return male.Male, err
}
