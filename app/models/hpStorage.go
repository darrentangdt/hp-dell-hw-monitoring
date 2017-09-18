package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type HpStorage struct {
	MonitorId                 int64  `db:"MonitorId, notnull"`
	CapacityMiB               int    `db:"CapacityMiB" json:"CapacityMiB"`
	CurrentTemperatureCelsius int    `db:"CurrentTemperatureCelsius" json:"CurrentTemperatureCelsius"`
	InterfaceType             string `db:"InterfaceType" json:"InterfaceType"`
	Location                  string `db:"Location" json:"Location"`
	LocationFormat            string `db:"LocationFormat" json:"LocationFormat"`
	MaximumTemperatureCelsius int    `db:"MaximumTemperatureCelsius" json:"MaximumTemperatureCelsius"`
	Model                     string `db:"Model" json:"Model"`
	SerialNumber              string `db:"SerialNumber, notnull" json:"SerialNumber"`
	HpStorageStatus           `db:"Status" json:"Status"`
	CreatedAt                 string `db:"CreatedAt"`
}

type HpStorageStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *HpStorage) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

type HpStorageLinkJson struct {
	HpStorageLinks `json:"links"`
}

type HpStorageLinks struct {
	HpStorageMembers []HpNetworkMember `json:"Member"`
}

type HpStorageMember struct {
	Href string `json:"href"`
}
