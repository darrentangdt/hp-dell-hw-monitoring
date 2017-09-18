package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type HpFan struct {
	MonitorId      int64  `db:"MonitorId, notnull"`
	CurrentReading int    `db:"CurrentReading" json:"CurrentReading"`
	FanName        string `db:"FanName, notnull" json:"FanName"`
	HpFanOem       `db:"Oem" json:"Oem"`
	HpFanStatus    `db:"Status" json:"Status"`
	Units          string `db:"Units" json:"Units"`
	CreatedAt      string `db:"CreatedAt"`
}

type HpFanOem struct {
	HpFanHp `db:"Hp" json:"Hp"`
}

type HpFanHp struct {
	Location string `db:"Location" json:"Location"`
	Type     string `db:"Type" json:"Type"`
}

type HpFanStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *HpFan) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func GetFansHealth(fans []HpFan) string {
	for _, f := range fans {
		if f.HpFanStatus.Health != "OK" {
			return "Warning"
		}
	}
	return "OK"
}

type HpTemperature struct {
	MonitorId           int64  `db:"MonitorId, notnull"`
	Context             string `db:"Context" json:"Context"`
	CurrentReading      int    `db:"CurrentReading" json:"CurrentReading"`
	Name                string `db:"Name, notnull" json:"Name"`
	Number              int    `db:"Number" json:"Number"`
	HpTemperatureOem    `db:"Oem" json:"Oem"`
	HpTemperatureStatus `db:"Status" json:"Status"`
	Units               string `db:"Units" json:"Units"`
	CreatedAt           string `db:"CreatedAt"`
}

type HpTemperatureOem struct {
	HpTemperatureHp `db:"Hp" json:"Hp"`
}

type HpTemperatureHp struct {
	LocationXmm int    `db:"LocationXmm" json:"LocationXmm"`
	LocationYmm int    `db:"LocationYmm" json:"LocationYmm"`
	Type        string `db:"Type" json:"Type"`
}

type HpTemperatureStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *HpTemperature) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func GetTemperaturesHealth(temperatures []HpTemperature) string {
	for _, t := range temperatures {
		if t.HpTemperatureStatus.Health != "OK" && t.HpTemperatureStatus.State != "Absent" {
			return "Warning"
		}
	}
	return "OK"
}

type HpFanTemperatureJson struct {
	Fans         []HpFan         `json:"Fans"`
	Temperatures []HpTemperature `json:"Temperatures"`
}
