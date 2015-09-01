package genmai

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	now := time.Now()
	memo := "memo"
	usePoint := uint64(1000)
	user := &User{Name: "test", CreateAt: &now, Memo: &memo, UsePoint: &usePoint}
	actual, err := Insert(user)

	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if actual != 1 {
		t.Errorf("insert failed. affected row amount is ", actual)
	}
}

func TestFind_findUser(t *testing.T) {
	user, err := FindOne(1)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if user == nil {
		t.Errorf("cannot find a user. %v", user)
	}
}

func TestFind_noexist_user(t *testing.T) {
	user, err := FindOne(0)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if user != nil {
		t.Errorf("cannot find a user. %v", user)
	}
}

func TestFind(t *testing.T) {
	users, err := Find("test")
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if len(users) == 0 {
		t.Errorf("cannot find a user. %v", users)
	}
}

func TestFind_notexist(t *testing.T) {
	users, err := Find("not_exist")
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if len(users) != 0 {
		t.Errorf("cannot find a user. %v", users)
	}
}
