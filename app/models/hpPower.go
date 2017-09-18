package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type HpPower struct {
	MonitorId     int64  `db:"MonitorId, notnull"`
	Model         string `db:"Model" json:"Model"`
	Name          string `db:"Name" json:"Name"`
	HpPowerOem    `json:"Oem"`
	HpPowerStatus `json:"Status"`
	CreatedAt     string `db:"CreatedAt"`
}

type HpPowerOem struct {
	HpPowerHp `json:"Hp"`
}

type HpPowerHp struct {
	BayNumber int `db:"BayNumber, notnull" json:"BayNumber"`
}

type HpPowerStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *HpPower) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func GetHpPowersHealth(powers []HpPower) string {
	for _, p := range powers {
		if p.HpPowerStatus.Health != "OK" {
			return "Warning"
		}
	}
	return "OK"
}

type HpPowerJson struct {
	Powers []HpPower `json:"PowerSupplies"`
}
