package genmai

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
	"log"
	_ "time"
)

type Seq struct {
	Id uint64 `db:pk default:0`
}

func getSequence(db *genmai.DB) (uint64, error) {
	tx, err := db.DB().Begin()
	if err != nil {
		return 0, err
	}
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

// Insert record into user table. return affected row amount
func Insert(user *User) (res int64, err error) {
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

	id, err := getSequence(db)
	if err != nil {
		log.Printf("cannot get id %v", err)
		return 0, err
	}

	user.Id = id
	log.Println(user.Id)
	res, err = db.Insert(user)

	if err != nil {
		return 0, err
	}

	return res, nil
}

// FindOne execute select * from user where id = ?
func FindOne(id uint64) (*User, error) {
	db, err := GetDb()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var user []User
	if err := db.Select(&user, db.Where("id", "=", id)); err != nil {
		log.Printf("an error occurred. %v", err)
		return nil, err
	}
	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}

func Find(name string) ([]User, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	var users []User
	if err := db.Select(&users, db.Where("name", "=", name)); err != nil {
		return nil, err
	}
	return users, nil
}
