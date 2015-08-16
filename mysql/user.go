package mysql

import "time"

type User struct {
	Id       uint64
	Name     string
	CreateAt time.Time
}
