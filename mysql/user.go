package mysql

import "time"

// User Object
type User struct {
	Id       uint64
	Name     string
	CreateAt *time.Time
	Memo     *string
	UsePoint *uint64
}
