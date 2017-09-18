package models

import (
	"time"

	"github.com/go-gorp/gorp"
)

const (
	OK       string = "OK"
	Ok              = "Ok"
	WARNING         = "Warning"
	CRITICAL        = "Critical"
	ABSENT          = "Absent"
)

type OverallState struct {
	MonitorId   int64  `db:"MonitorId, notnull"`
	Host        string `db:"Host, notnull"`
	Service     string `db:"Service, notnull"`
	Manager     string `db:"Manager, notnull"`
	Power       string `db:"Power, notnull"`
	State       string `db:"State, notnull"`
	Fan         string `db:"Fan, notnull"`
	Temperature string `db:"Temperature, notnull"`
	PowerSupply string `db:"PowerSupply, notnull"`
	Network     string `db:"Network, notnull"`
	Disk        string `db:"Disk, notnull"`
	CreatedAt   string `db:"CreatedAt"`
}

func (c *OverallState) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (c *OverallState) CheckMonitorState(monitor Monitor) {
	c.MonitorId = monitor.Id
	c.Host = monitor.Host
	c.Service = monitor.Service
	c.Manager = monitor.Manager
}

func (c *OverallState) HpCheckSystemState(system HpSystem) {
	c.Power = system.Power
	c.State = system.HpSystemStatus.Health
}

func (c *OverallState) DellCheckSystemState(system DellSystem) {
	c.Power = system.PowerState
	c.State = system.Status.Health
}

func (c *OverallState) HpCheckTemperatureState(temperatures []HpTemperature) {
	for _, temperature := range temperatures {
		if temperature.HpTemperatureStatus.Health != OK && temperature.HpTemperatureStatus.State != ABSENT {
			c.Temperature = temperature.HpTemperatureStatus.Health
			return
		}
	}
	c.Temperature = OK
}

func (c *OverallState) HpCheckFanState(fans []HpFan) {
	for _, fan := range fans {
		if fan.HpFanStatus.Health != OK {
			c.Fan = fan.HpFanStatus.Health
			return
		}
	}
	c.Fan = OK
}

func (c *OverallState) DellCheckFanState(fans []DellFan) {
	for _, fan := range fans {
		if fan.Status.Health != OK {
			c.Fan = fan.Status.Health
			return
		}
	}
	c.Fan = OK
}

func (c *OverallState) DellCheckTemperatureState(temperatures []DellTemperature) {
	for _, temperature := range temperatures {
		if temperature.Status.Health != OK {
			c.Temperature = temperature.Status.Health
			return
		}
	}
	c.Temperature = OK
}

func (c *OverallState) DellCheckNetworkState(networks []DellNetwork) {
	for _, network := range networks {
		if network.Status.Health != Ok {
			c.Network = network.Status.Health
			return
		}
	}
	c.Network = OK
}

func (c *OverallState) DellCheckPowerSupplyState(powersupplies []DellPowerSupply) {
	for _, powersupply := range powersupplies {
		if powersupply.Status.Health != OK {
			c.PowerSupply = powersupply.Status.Health
			return
		}
	}
	c.PowerSupply = OK
}

func (c *OverallState) DellCheckDiskState(storages []DellStorage) {
	for _, storage := range storages {
		if storage.Status.Health != OK {
			c.Disk = storage.Status.Health
			return
		}
	}
	c.Disk = OK
}
func (c *OverallState) HpCheckPowerState(powers []HpPower) {
	for _, power := range powers {
		if power.HpPowerStatus.Health != OK {
			c.PowerSupply = power.HpPowerStatus.Health
			return
		}
	}
	c.PowerSupply = OK
}

func (c *OverallState) HpCheckNetworkState(networks []HpNetwork) {
	for _, network := range networks {
		if network.HpNetworkStatus.Health != OK {
			c.Network = network.HpNetworkStatus.Health
			return
		}
	}
	c.Network = OK
}

func (c *OverallState) HpCheckDiskState(disks []HpStorage) {
	for _, disk := range disks {
		if disk.HpStorageStatus.Health != OK {
			c.Disk = disk.HpStorageStatus.Health
			return
		}
	}
	c.Disk = OK
}

func (c *OverallState) CheckAllStateNotInserted() bool {
	return c.MonitorId == 0 || c.Host == "" || c.Service == "" ||
		c.Manager == "" || c.Power == "" || c.State == "" ||
		c.Fan == "" || c.Temperature == "" || c.PowerSupply == "" ||
		c.Network == "" || c.Disk == ""
}
