package genmai

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InsertUserAccount(userAccount *UserAccount) (res int64, err error) {
	db, err := GetDb()
	if err != nil || db == nil {
		log.Printf("cannot get datasource. %v", err)
		return 0, err
	}
	defer func() {
		if err != nil {
			log.Printf("call rollback")
			db.Rollback()
		} else {
			log.Printf("call commit")
			db.Commit()
		}
		db.Close()
	}()
	if err := db.Begin(); err != nil {
		log.Printf("cannot begin transaction %v", err)
		return 0, err
	}
	res, err = db.Insert(userAccount)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func FindUserAccount(id uint64) ([]UserAccount, error) {
	db, err := GetDb()
	if err != nil || db == nil {
		log.Printf("cannot get datasource. %v", err)
		return nil, err
	}

	var accounts []UserAccount
	if err := db.Select(&accounts, db.Where("id", "=", 1)); err != nil {
		log.Println("an error occurred")
		return nil, err
	}
	return accounts, nil
}
