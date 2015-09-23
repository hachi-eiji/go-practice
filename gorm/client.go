package gorm

import (
	"github.com/jinzhu/gorm"
	"log"
)

func getSequence(db gorm.DB) (uint64, error) {
	tx := db.Begin()

	if err := tx.Model(&Seq{}).Update("id", gorm.Expr("LAST_INSERT_ID(id + 1)")).Error; err != nil {
		return 0, err
	}
	var id uint64
	r := tx.Raw("SELECT LAST_INSERT_ID()").Row()
	r.Scan(&id)
	return id, nil

}

func Insert(user *User) (res uint64, err error) {
	db, err := GetDB()

	if err != nil {
		log.Printf("cannot get datasource. %v", err)
		return 0, err
	}
	defer db.Close()

	id, err := getSequence(db)
	if err != nil {
		log.Printf("cannot get id %v", err)
		return 0, err
	}
	user.Id = id
	r := db.Create(user)
	if err := r.Error; err != nil {
		return 0, err
	}
	return uint64(r.RowsAffected), nil
}

func FindOne(id uint64) (*User, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	// FIRSTはorder by してる...
	// SELECT  * FROM `user`  WHERE (id = '1') ORDER BY `user`.`id` ASC LIMIT 1
	err = db.Where("id = ?", id).Find(&user).Error
	if err != nil && err != gorm.RecordNotFound {
		return nil, err
	}
	if err == gorm.RecordNotFound {
		return nil, nil
	}
	return &user, nil
}

func Find(name string) ([]User, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	var users []User
	if err := db.Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
