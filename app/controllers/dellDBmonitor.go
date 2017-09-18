package controllers

import (
	"log"
	"tms/app/dbs"
	"tms/app/models"
	"tms/app/sockets"
)

func DellSetCurrentSystemState(monitor models.Monitor, overallState *models.OverallState) {
	var system models.DellSystem
	url := "https://" + monitor.Host + ":443/redfish/v1/Systems/System.Embedded.1"
	if err := DellHttpGetRequest(url, monitor, &system); err != nil {
		log.Println(err)
	} else {
		system.MonitorId = monitor.Id
		dbs.InsertSystem(&system)

	}
	overallState.DellCheckSystemState(system)
}

func DellSetCurrentChassisState(monitor models.Monitor, overallState *models.OverallState) {
	var fanTemperatureJson models.DellFanTemperature
	var fans []models.DellFan
	var temperatures []models.DellTemperature

	url := "https://" + monitor.Host + ":443/redfish/v1/Chassis/System.Embedded.1/Thermal"
	if err := DellHttpGetRequest(url, monitor, &fanTemperatureJson); err != nil {
	} else {
		for _, fan := range fanTemperatureJson.DellFans {
			fan.MonitorId = monitor.Id
			dbs.InsertFan(&fan)
			fans = append(fans, fan)
		}
		for _, temperature := range fanTemperatureJson.DellTemperatures {
			temperature.MonitorId = monitor.Id
			dbs.InsertTemperature(&temperature)
			temperatures = append(temperatures, temperature)
		}
	}
	overallState.DellCheckFanState(fans)
	overallState.DellCheckTemperatureState(temperatures)
}

func DellSetCurrentStorageState(monitor models.Monitor, overallState *models.OverallState) {
	var StorageJson models.DellStorages
	var Storages []models.DellStorage
	url := "https://" + monitor.Host + ":443/redfish/v1/Systems/System.Embedded.1/Storage/Controllers/RAID.Integrated.1-1"
	if err := DellHttpGetRequest(url, monitor, &StorageJson); err != nil {
		log.Println(err)
	} else {
		for _, Storage := range StorageJson.DellDevices {
			Storage.MonitorId = monitor.Id
			dbs.InsertStorage(&Storage)
			Storages = append(Storages, Storage)
		}
	}
	overallState.DellCheckDiskState(Storages)
}

func DellSetCurrentPowerState(monitor models.Monitor, overallState *models.OverallState) {
	var powerJson models.DellPower
	var powers []models.DellPowerSupply

	url := "https://" + monitor.Host + ":443/redfish/v1/Chassis/System.Embedded.1/Power/"
	if err := DellHttpGetRequest(url, monitor, &powerJson); err != nil {
		log.Println(err)
	} else {
		for _, power := range powerJson.DellPowerSupplies {
			power.MonitorId = monitor.Id
			dbs.InsertPower(&power)
			powers = append(powers, power)
		}
	}
	overallState.DellCheckPowerSupplyState(powers)
}

func DellSetCurrentNetworkState(monitor models.Monitor, overallState *models.OverallState) {
	var NetworksJson models.DellNetworkCollection
	var DellNetworks []models.DellNetwork
	var Count int
	url := "https://" + monitor.Host + ":443/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces"
	if err := DellHttpGetRequest(url, monitor, &NetworksJson); err != nil {
		log.Println(err)
	} else {
		Count = NetworksJson.OdataCount
		DellNetworks = make([]models.DellNetwork, Count)
		for i, odatamember := range NetworksJson.Members {
			url = "https://" + monitor.Host + ":443" + odatamember.OdataID
			DellHttpGetRequest(url, monitor, &DellNetworks[i]) //err check
			DellNetworks[i].MonitorId = monitor.Id
			dbs.InsertNetwork(&DellNetworks[i])

		}
	}
	overallState.DellCheckNetworkState(DellNetworks)
}

func DellInsertCurrentState(monitor models.Monitor) {
	overallState := models.OverallState{}
	overallState.CheckMonitorState(monitor)

	DellSetCurrentSystemState(monitor, &overallState)
	DellSetCurrentChassisState(monitor, &overallState)
	DellSetCurrentPowerState(monitor, &overallState)
	DellSetCurrentNetworkState(monitor, &overallState)
	DellSetCurrentStorageState(monitor, &overallState)

	if !overallState.CheckAllStateNotInserted() {
		if components := dbs.GetComponentsIfFaultOccured(overallState); components != "" {
			sockets.Alert(overallState.Host, components)
			alarmlog := dbs.MakeAlarmLog(overallState)
		}
		dbs.InsertOverallState(overallState)
		dbs.InsertAlarmLog(alarmlog)
	}
}

func DellGetMonitorManuf(monitor models.Monitor) string {
	var system models.DellSystem
	url := "https://" + monitor.Host + "/redfish/v1/Systems/System.Embedded.1"
	if err := DellHttpGetRequest(url, monitor, &system); err != nil {
		log.Println(err)
	}
	return system.Manufacturer
}
