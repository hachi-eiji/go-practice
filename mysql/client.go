package mysql

import (
	"bytes"
	"database/sql"
	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"text/template"
)

type config struct {
	Master dbConfig `toml:"master"`
	Slave  dbConfig
}

type dbConfig struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"db"`
	Param    string `toml:"param"`
}

var appConfig config

func loadConfig() {
	if &appConfig == nil {
		return
	}
	_, err := toml.DecodeFile(os.Getenv("appConf"), &appConfig)
	if err != nil {
		panic(err)
	}
}

func getDb() (*sql.DB, error) {
	loadConfig()
	tmpl, err := template.New("test").Parse("{{.User}}:{{.Password}}@tcp({{.Host}}:{{.Port}})/{{.Name}}?{{.Param}}")
	if err != nil {
		return nil, err
	}
	var url bytes.Buffer
	if err := tmpl.Execute(&url, appConfig.Master); err != nil {
		return nil, err
	}
	return sql.Open("mysql", url.String())
}

// getSequence get sequence number
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

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.CreateAt)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("scan error %v\n", err)
		return nil, err

		// data not found
	} else if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return user, nil
}

// Find is executing select query
func Find(users *[]User, name string) error {
	db, err := getDb()
	defer db.Close()
	if err != nil {
		log.Printf("can not get database. %v\n", err.Error())
		return err
	}
	stmt, err := db.Prepare("SELECT id, name,createAt,memo, use_point FROM user WHERE name like ?")
	if err != nil {
		log.Printf("statement error %v\n", err.Error())
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	defer rows.Close()
	if err != nil {
		log.Printf("statement error %v\n", err.Error())
		return err
	}

	//	columns, err := rows.Columns()
	_, err = rows.Columns()
	if err != nil {
		log.Printf("cannot get columns%v\n", err.Error())
		return err
	}
	//	results := make([]User, 0)
	for rows.Next() {
		//		user := new(User)
		var (
			id       uint64
			name     string
			createAt mysql.NullTime
			memo     sql.NullString
			usePoint *uint64
		)
		if err := rows.Scan(&id, &name, &createAt, &memo, &usePoint); err != nil {
			log.Printf("cannot scan %v\n", err.Error())
			return err
		}

		user := User{Id: id, Name: name}

		if createAt.Valid {
			user.CreateAt = &createAt.Time
		}
		if memo.Valid {
			user.Memo = &memo.String
		}

		if usePoint != nil {
			user.UsePoint = usePoint
		}
		*users = append(*users, user)
	}
	return nil
}
