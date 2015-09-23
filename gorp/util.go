package gorp

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func GetDB() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "test:", log.Lmicroseconds))
	// この設定をしないとinsert,Getがうごかない
	dbMap.AddTableWithName(User{}, "user").SetKeys(false, "id")
	return dbMap, nil
}
