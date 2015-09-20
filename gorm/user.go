package gorm

import "time"

type User struct {
	Id       uint64 `gorm:"primary_key"`
	Name     string
	CreateAt *time.Time
	Memo     *string
	UsePoint *uint64
}
