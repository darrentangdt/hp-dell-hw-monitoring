package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type HpNetwork struct {
	MonitorId       int64  `db:"MonitorId, notnull"`
	MacAddress      string `db:"MacAddress, notnull" json:"MacAddress"`
	SpeedMbps       int    `db:"SpeedMbps" json:"SpeedMbps"`
	HpNetworkStatus `db:"Status" json:"Status"`
	CreatedAt       string `db:"CreatedAt"`
}

type HpNetworkStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *HpNetwork) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

type HpNetworkJson struct {
	HpNetworks           []HpNetwork `json:"PhysicalPorts"`
	HpNetworkTotalStatus `json:"Status"`
}

type HpNetworkTotalStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

type HpNetworkLinkJson struct {
	HpNetworkLinks `json:"links"`
}

type HpNetworkLinks struct {
	HpNetworkMembers []HpNetworkMember `json:"Member"`
}

type HpNetworkMember struct {
	Href string `json:"href"`
}
