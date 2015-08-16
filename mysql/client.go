package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getDb() (*sql.DB, error) {
	return sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&columnsWithAlias=true")

}

// getSequence get sequcence number
func getSequence(tx *sql.Tx) (id uint64, err error) {
	if _, err := tx.Query("UPDATE seq SET id = LAST_INSERT_ID(id + 1)"); err != nil {
		return 0, err
	}
	if err := tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id); err != nil {
		return 0, err
	}
	return
}

// Insert user data
func Insert(user *User) (int64, error) {
	db, err := getDb()
	defer func() {
		log.Println("close connection")
		db.Close()
	}()
	if err != nil {
		log.Println("can not get database. %v", err.Error())
		return 0, err
	}

	stmt, err := db.Prepare("INSERT INTO user(id, name, createAt) values(?,?,?)")
	if err != nil {
		log.Printf("statement error %v\n", err.Error())
		return 0, err
	}
	defer func() {
		log.Println("close statment")
		stmt.Close()
	}()

	tx, err := db.Begin()
	if err != nil {
		log.Printf("transaction begin failed. %v\n", err)
		return 0, err
	}

	user.Id, err = getSequence(tx)
	if err != nil {
		log.Println("can not get sequence")
		return 0, err
	}

	result, err := tx.Stmt(stmt).Exec(user.Id, user.Name, user.CreateAt)
	if err != nil {
		_err := tx.Rollback()
		if _err != nil {
			log.Panicf("can not rollback transaction %v", _err.Error())
			return 0, _err
		}
		log.Println("execute query error %v", err.Error())
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("can not commit transaction. %v\n", err.Error())
		return 0, err
	}
	var res, _ = result.RowsAffected()
	return res, nil
}

// FindOne execute select query and returns a User
func FindOne(id int64) (*User, error) {
	db, err := getDb()
	defer func() {
		log.Println("close connection")
		db.Close()
	}()
	if err != nil {
		log.Printf("can not get database. %v\n", err.Error())
		return nil, err
	}
	stmt, err := db.Prepare("SELECT id, name, createAt FROM user WHERE id = ?")
	if err != nil {
		log.Printf("statement error %v\n", err.Error())
		return nil, err
	}
	defer func() {
		log.Println("close statment")
		stmt.Close()
	}()
	user := new(User) // ポインタを返却

	if err := stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.CreateAt); err != nil {
		log.Printf("scan error %v\n", err)
		return nil, err
	}
	return user, nil
}
