package user

import (
	mock "go-notes/testmethod/mock"
	"go-notes/testmethod/person"
	"testing"

	"github.com/golang/mock/gomock"
)

// go test -v -run=GetUserInfo *.go
func Test_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl)
	maleInfo := person.MaleInfo{Male: 222222}
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(maleInfo, nil),
	)

	user := NewUser(mockMale)
	male, err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
	t.Logf("user male is: %+v ", male)
}
