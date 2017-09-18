package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type DellPowerSupply struct {
	MonitorId  int64  `db:"MonitorId, notnull"`
	Model      string `db:"Model" json:"Model"`
	MemberID   string `db:"MemberID" json:"MemberID"`
	PartNumber string `db:"PartNumber, notnull" json:"PartNumber"`
	Status     `db:"Status" json:"Status"`
	CreatedAt  string `db:"CreatedAt"`
}

type DellPower struct {
	DellPowerSupplies []DellPowerSupply `json:"PowerSupplies"`
}

func (c *DellPowerSupply) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
