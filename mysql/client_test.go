package mysql

import "testing"
import "time"

func TestInsert(t *testing.T) {
	user := &User{Name: "bar", CreateAt: time.Now()}
	ret, err := Insert(user)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if ret != 1 {
		t.Errorf("result is not 1. actual %d", ret)
	}
}

func TestFind(t *testing.T) {
	// 2015-08-16 13:57:11
	createAt := time.Date(2015, time.August, 16, 13, 57, 11, 0, time.UTC)
	expect := &User{Id: 2, Name: "hoge", CreateAt: createAt}
	actual, err := FindOne(2)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}

	if !(expect.Id == actual.Id && expect.Name == actual.Name && expect.CreateAt == actual.CreateAt) {
		t.Errorf("expect not equal actual %v=%v", expect, actual)
	}
}
