package genmai

import (
	"testing"
)

func TestInsertUserAccount(t *testing.T) {
	u := &UserAccount{
		Id:          2,
		AccountType: "twitter",
		AccountName: "test_tw",
	}
	actual, err := InsertUserAccount(u)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if actual != 1 {
		t.Errorf("insert failed. affected row amount is ", actual)
	}
}

func TestFindUserAccount(t *testing.T) {
	accounts, err := FindUserAccount(1)
	if err != nil {
		t.Errorf("an error occurred. %v", err)
	}
	if len(accounts) != 2 {
		t.Errorf("account lenght is not 2. %v", accounts)
	}
}
