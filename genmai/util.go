package genmai

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
	"os"
)

func GetDb() (*genmai.DB, error) {
	url := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err := genmai.New(&genmai.MySQLDialect{}, url)
	if err != nil {
		return nil, err
	}
	db.SetLogOutput(os.Stdout)
	return db, nil
}
