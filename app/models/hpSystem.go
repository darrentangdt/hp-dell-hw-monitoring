package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type HpSystem struct {
	MonitorId      int64 `db:"MonitorId, notnull"`
	HpBios         `db:"Bios" json:"Bios"`
	Manufacturer   string `db:"Manufacturer" json:"Manufacturer"`
	Model          string `db:"Model" json:"Model"`
	Name           string `db:"Name" json:"Name"`
	Power          string `db:"Power" json:"Power"`
	HpProcessors   `db:"Processors" json:"Processors"`
	SKU            string `db:"SKU" json:"SKU"`
	SerialNumber   string `db:"SerialNumber, notnull" json:"SerialNumber"`
	HpSystemStatus `db:"Status" json:"Status"`
	CreatedAt      string `db:"CreatedAt"`
}

type HpBios struct {
	HpCurrent `db:"Current" json:"Current"`
}

type HpCurrent struct {
	VersionString string `db:"VersionString" json:"VersionString"`
}

type HpProcessors struct {
	Count             int    `db:"Count" json:"Count"`
	ProcessorFamily   string `db:"ProcessorFamily" json:"ProcessorFamily"`
	HpProcessorStatus `db:"Status" json:"Status"`
}

type HpProcessorStatus struct {
	HealthRollUp string `db:"HealthRollUp" json:"HealthRollUp"`
}

type HpSystemStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *HpSystem) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
