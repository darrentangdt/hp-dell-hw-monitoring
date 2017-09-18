package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type AlarmLog struct {
	MonitorId   int64  `db:"MonitorId"`
	MonitorHost string `db:"MonitorHost"`
	Component   string `db:"Component"`
	Describe    string `db:"Describe"`
	CreatedAt   string `db:"CreatedAt"`
}

func (c *AlarmLog) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2017-08-03 15:10:05")
	return nil
}
