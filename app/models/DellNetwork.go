package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

type DellNetworkCollection struct {
	OdataCount int      `json:"@odata.count"`
	Members    []Member `json:"Members"`
}

type Member struct {
	OdataID string `json:"@odata.id"`
}

type DellNetwork struct {
	MonitorId              int64  `db:"MonitorId, notnull"`
	FQDN                   string `db:"FQDN" json:"FQDN"`
	FullDuplex             bool   `db:"FullDuplex" json:"FullDuplex"`
	HostName               string `db:"HostName" json:"HostName"`
	IPV6DefaultGateway     string `db:"IPV6DefaultGateway" json:"IPV6DefaultGateway"`
	IPv4Addresses          string `db:"IPv4Addresses" json:"IPv4Addresses"`
	IPv6AddressPolicyTable string `db:"IPv6AddressPolicyTable" json:"IPv6AddressPolicyTable"`
	IPv6Addresses          string `db:"IPv6Addresses" json:"IPv6Addresses"`
	IPv6StaticAddresses    string `db:"IPv6StaticAddresses" json:"IPv6StaticAddresses"`
	ID                     string `db:"Id" json:"Id"`
	InterfaceEnabled       string `db:"InterfaceEnabled" json:"InterfaceEnabled"`
	MTUSize                string `db:"MTUSize" json:"MTUSize"`
	MacAddress             string `db:"MacAddress, notnull" json:"MacAddress"`
	MaxIPv6StaticAddresses string `db:"MaxIPv6StaticAddresses" json:"MaxIPv6StaticAddresses"`
	NameServers            string `db:"NameServers" json:"NameServers"`
	PermanentMACAddress    string `db:"PermanentMACAddress" json:"PermanentMACAddress"`
	SpeedMbps              int    `db:"SpeedMbps" json:"SpeedMbps"`
	Status                 `db:"Status" json:"Status"`
	UefiDevicePath         string `db:"UefiDevicePath" json:"UefiDevicePath"`
	VLAN                   string `db:"VLAN" json:"VLAN"`
	VLANs                  string `db:"VLANs" json:"VLANs"`
	CreatedAt              string `db:"CreatedAt"`
}

func (c *DellNetwork) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
