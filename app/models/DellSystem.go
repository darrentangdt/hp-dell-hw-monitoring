package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type DellSystem struct {
	MonitorId        int64 `db:"MonitorId, notnull"`
	MemorySummary    `db:"MemorySummary" json:"MemorySummary"`
	Model            string `db:"Model" json:"Model"`
	PowerState       string `db:"PowerState" json:"PowerState"`
	ProcessorSummary `db:"ProcessorSummary" json:"ProcessorSummary"`
	SKU              string `db:"SKU" json:"SKU"`
	SerialNumber     string `db:"SerialNumber, notnull" json:"SerialNumber"`
	Status           `db:"Status" json:"Status"`
	Manufacturer     string `db:"Manufacturer" json:"Manufacturer"`
	CreatedAt        string `db:"CreatedAt"`
}

type MemorySummary struct {
	MemoryStatus         `db:"Status" json:"Status"`
	TotalSystemMemoryGiB float64 `db:"TotalSystemMemoryGiB" json:"TotalSystemMemoryGiB"`
}

type MemoryStatus struct {
	MemoryState  string `db:"MemoryState" json:"State"`
	MemoryHealth string `db:"MemoryHealth" json:"Health"`
}

type ProcessorSummary struct {
	ProcessorCount  int    `db:"ProcessorCount" json:"Count"`
	ProcessorModel  string `db:"ProcessorModel" json:"Model"`
	ProcessorStatus `json:"Status"`
}

type ProcessorStatus struct {
	ProcessorState  string `db:"processorState" json:"State"`
	ProcessorHealth string `db:"processorHealth" json:"Health"`
}

func (c *DellSystem) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
