package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type DellFanTemperature struct {
	DellFans         []DellFan         `json:"Fans"`
	DellTemperatures []DellTemperature `json:"Temperatures"`
}

type DellFan struct {
	MonitorId       int64  `db:"MonitorId, notnull"`
	FanName         string `db:"FanName, notnull" json:"FanName"`
	PhysicalContext string `db:"PhysicalContext" json:"PhysicalContext"`
	ReadingRPM      int    `db:"ReadingRPM" json:"ReadingRPM"`
	Status          `db:"Status" json:"Status"`
	CreatedAt       string `db:"CreatedAt"`
}

type DellTemperature struct {
	MonitorId       int64  `db:"MonitorId, notnull"`
	Name            string `db:"Name, notnull" json:"Name"`
	PhysicalContext string `db:"PhysicalContext" json:"PhysicalContext"`
	ReadingCelsius  int    `db:"ReadingCelsius" json:"ReadingCelsius"`
	Status          `db:"Status" json:"Status"`
	CreatedAt       string `db:"CreatedAt"`
}

func (c *DellFan) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
func (c *DellTemperature) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
