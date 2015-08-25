package mysql

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func init() {
	// set env
	os.Setenv("appConf", "app_test.toml")
}

func TestInsert(t *testing.T) {
	now := time.Now()
	user := &User{Name: "bar", CreateAt: &now}
	ret, err := Insert(user)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if ret != 1 {
		t.Errorf("result is not 1. actual %d", ret)
	}
}

func TestFindOne(t *testing.T) {
	// 2015-08-16 13:57:11
	createAt := time.Date(2015, time.August, 16, 13, 57, 11, 0, time.UTC)
	expect := &User{Id: 2, Name: "hoge", CreateAt: &createAt}
	actual, err := FindOne(2)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}

	if !(expect.Id == actual.Id && expect.Name == actual.Name && *expect.CreateAt == *actual.CreateAt) {
		t.Errorf("expect not equal actual %v=%v", expect, actual)
	}
}

func TestFindOne_not_found(t *testing.T) {
	actual, err := FindOne(-1)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}

	if actual != nil {
		t.Errorf("found some data")
	}
}

func TestFind(t *testing.T) {
	users, err := Find("hoge")
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	for i, v := range users {
		fmt.Printf("%v, %v %d\n", i, v)
	}
}

func TestFind_not_found(t *testing.T) {
	users, err := Find("not_found")
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if len(users) != 0 {
		t.Errorf("data found %v.", users)
	}
}
