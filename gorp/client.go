package gorp

import (
	"github.com/go-gorp/gorp"
	"log"
)

type Seq struct {
	Id uint64
}

func getSequence(dbMap *gorp.DbMap) (uint64, error) {
	tx, err := dbMap.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	st, err := tx.Prepare("UPDATE seq set id = LAST_INSERT_ID(id + 1)")
	if err != nil {
		return 0, err
	}
	defer st.Close()
	st.Exec()

	stmt, err := tx.Prepare("SELECT LAST_INSERT_ID()")
	defer stmt.Close()
	var id uint64
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Insert(user *User) (res int64, err error) {
	dbMap, err := GetDB()
	if err != nil {
		log.Printf("cannot get datasource. %v", err)
		return 0, err
	}
	id, err := getSequence(dbMap)
	if err != nil {
		return 0, nil
	}
	user.Id = id

	tx, err := dbMap.Begin()
	if err != nil {
		log.Printf("cannot begin transaction %v", err)
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		dbMap.Db.Close()
	}()

	if err != nil {
		log.Printf("cannot get id %v", err)
	}
	if err = dbMap.Insert(user); err != nil {
		log.Printf("an error occurred %v", err)
		return 0, err
	}
	return 1, nil
}

func FindOne(id uint64) (*User, error) {
	dbMap, err := GetDB()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}
	o, err := dbMap.Get(User{}, id)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, nil
	}

	user := o.(*User)
	return user, nil
}

func Find(name string) ([]User, error) {
	dbMap, err := GetDB()
	if err != nil {
		return nil, err
	}
	var users []User
	_, err = dbMap.Select(&users, "select * from user where name = ?", name)
	if err != nil {
		return nil, err
	}
	return users, nil
}
