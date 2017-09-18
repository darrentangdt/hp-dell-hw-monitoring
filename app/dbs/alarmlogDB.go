package dbs

import (
	"fmt"
	"log"
	"tms/app/models"
)

type FaultStruct struct {
	Component string
	Fault     string
}

func InsertAlarmLog(alarmlog models.AlarmLog) {
	if err := Dbm.Insert(&alarmlog); err != nil {
		log.Println(err)
	}
}

func SelectAllAlarmLogsByCondition(alarmlog models.AlarmLog) []models.AlarmLog {
	var condition string
	if alarmlog.MonitorHost != "" {
		condition += " MonitorHost like '%" + alarmlog.MonitorHost + "%' "
		if alarmlog.Component != "" {
			condition += "and"
		}
	}
	if alarmlog.Component != "" {
		condition += " Component like '%" + alarmlog.Component + "%' "
	}
	var alarmlogs []models.AlarmLog
	if _, err := Dbm.Select(&alarmlogs, GetRecordsQueryByCondition("AlarmLog", condition)); err != nil {
		log.Println(err)
	}
	fmt.Println(alarmlogs)
	return alarmlogs
}

func MakeAlarmLogOverall(overallState models.OverallState) models.AlarmLog {
	var alarmlog models.AlarmLog
	alarmlog.MonitorId = overallState.MonitorId
	alarmlog.MonitorHost = overallState.Host
	alarmlog.Component = GetComponentsIfFaultOccured(overallState)
	alarmlog.Describe = alarmlog.Component + "Fault"
	return alarmlog
}

func MakeAlarmLog(overallState models.OverallState) models.AlarmLog {
	var alarmlog models.AlarmLog
	alarmlog.MonitorId = overallState.MonitorId
	alarmlog.MonitorHost = overallState.Host
	alarmlog.Component = overallState.Host
	alarmlog.Describe = overallState.Host

	return alarmlog
}

func GetFaultFan(monitor models.Monitor) []FaultStruct {
	var fault FaultStruct
	var faults []FaultStruct
	fans := SelectFansNewestState(monitor)
	if monitor.Manuf == models.HP {
		for _, fan := range fans.([]models.HpFan) {
			if fan.HpFanStatus.Health != models.OK {
				fault.Component = fan.FanName
				fault.Fault = fan.HpFanStatus.Health + " , " + fan.HpFanStatus.State
				faults = append(faults, fault)
			}
		}
	} else {
		for _, fan := range fans.([]models.DellFan) {
			if fan.Status.Health != models.OK {
				fault.Component = fan.FanName
				fault.Fault = fan.Status.Health + " , " + fan.Status.State
				faults = append(faults, fault)
			}
		}
	}
	return faults
}

func GetFaultTemperature(monitor models.Monitor) []FaultStruct {
	var fault FaultStruct
	var faults []FaultStruct
	temperatures := SelectTemperaturesNewestState(monitor)
	if monitor.Manuf == models.HP {
		for _, temperature := range temperatures.([]models.HpTemperature) {
			if temperature.HpTemperatureStatus.Health != models.OK {
				fault.Component = temperature.Name
				fault.Fault = temperature.HpTemperatureStatus.Health + " , " + temperature.HpTemperatureStatus.State
				faults = append(faults, fault)
			}
		}
	} else {
		for _, temperature := range temperatures.([]models.DellTemperature) {
			if temperature.Status.Health != models.OK {
				fault.Component = temperature.Name
				fault.Fault = temperature.Status.Health + " , " + temperature.Status.State
				faults = append(faults, fault)
			}
		}
	}
	return faults
}

func GetFaultPower(monitor models.Monitor) []FaultStruct {
	var fault FaultStruct
	var faults []FaultStruct
	powers := SelectPowersNewestState(monitor)
	if monitor.Manuf == models.HP {
		for _, power := range powers.([]models.HpPower) {
			if power.HpPowerStatus.Health != models.OK {
				fault.Component = power.Model
				fault.Fault = power.HpPowerStatus.Health + " , " + power.HpPowerStatus.State
				faults = append(faults, fault)
			}
		}
	} else {
		for _, power := range powers.([]models.DellPowerSupply) {
			if power.Status.Health != models.OK {
				fault.Component = power.Model
				fault.Fault = power.Status.Health + " , " + power.Status.State
				faults = append(faults, fault)
			}
		}
	}
	return faults
}

func GetFaultDisk(monitor models.Monitor) []FaultStruct {
	var fault FaultStruct
	var faults []FaultStruct
	storages := SelectStoragesNewestState(monitor)
	if monitor.Manuf == models.HP {
		for _, storage := range storages.([]models.HpStorage) {
			if storage.HpStorageStatus.Health != models.OK {
				fault.Component = storage.Location + "-" + storage.Model
				fault.Fault = storage.HpStorageStatus.Health + " , " + storage.HpStorageStatus.State
				faults = append(faults, fault)
			}
		}
	} else {
		for _, storage := range storages.([]models.DellStorage) {
			if storage.Status.Health != models.OK {
				fault.Component = storage.Manufacturer + " " + storage.Model
				fault.Fault = storage.Status.Health + " , " + storage.Status.State
				faults = append(faults, fault)
			}
		}
	}
	return faults
}
func GetFaultNetwork(monitor models.Monitor) []FaultStruct {
	var fault FaultStruct
	var faults []FaultStruct
	networks := SelectNetworksNewestState(monitor)
	if monitor.Manuf == models.HP {
		for _, network := range networks.([]models.HpNetwork) {
			if network.HpNetworkStatus.Health != models.OK {
				fault.Component = network.MacAddress
				fault.Fault = network.HpNetworkStatus.Health + " , " + network.HpNetworkStatus.State
				faults = append(faults, fault)
			}
		}
	} else {
		for _, network := range networks.([]models.DellNetwork) {
			if network.Status.Health != models.OK {
				fault.Component = network.MacAddress
				fault.Fault = network.Status.Health + " , " + network.Status.State
				faults = append(faults, fault)
			}
		}
	}
	return faults
}
