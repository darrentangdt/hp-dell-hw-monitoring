package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type Alarm struct {
	Id        int64  `db:"Id"`
	UserName  string `db:"UserName"`
	UserPhone string `db:"UserPhone, notnull"`
	Host      string `db:"Host"`
	Service   string `db:"Service"`
	Power     bool   `db:"Power"`
	Thermal   bool   `db:"Thermal"`
	Fan       bool   `db:"Fan"`
	Network   bool   `db:"Network"`
	Disk      bool   `db:"Disk"`
	CreatedAt string `db:"CreatedAt"`
}

func (c *Alarm) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
