package gorp

import (
	"time"
)

type User struct {
	Id       uint64 `db:"id"`
	Name     string `db:"name"`
	CreateAt *time.Time `db:"create_at"`
	Memo     *string `db:"memo"`
	UsePoint *uint64 `db:"use_point"`
}
