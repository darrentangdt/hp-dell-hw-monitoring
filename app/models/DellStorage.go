package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type DellStorages struct {
	DellDevices []DellStorage `json:"Devices"`
}

type DellStorage struct {
	MonitorId    int64  `db:"MonitorId, notnull"`
	Manufacturer string `db:"Manufacturer" json:"Manufacturer"`
	Model        string `db:"Model" json:"Model"`
	Name         string `db:"Name, notnull" json:"Name"`
	Status       `db:"Status" json:"Status"`
	CreatedAt    string `db:"CreatedAt"`
}

type Status struct {
	Health       string `db:"Health" json:"Health"`
	HealthRollUp string `db:"HealthRollUp" json:"HealthRollUp"`
	State        string `db:"State" json:"State"`
}

func (c *DellStorage) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
