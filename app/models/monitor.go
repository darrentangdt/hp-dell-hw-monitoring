package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

const (
	HP   string = "HP"
	DELL        = "DELL"
)

type Monitor struct {
	Id          int64  `db:"Id"`
	Host        string `db:"Host, notnull"`
	Manager     string `db:"Manager"`
	Service     string `db:"Service"`
	MonitorName string `db:"MonitorName, notnull"`
	MonitorPass string `db:"MonitorPass, notnull"`
	Manuf       string `db:"Manuf"`
	CreatedAt   string `db:"CreatedAt"`
}

func (c *Monitor) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
