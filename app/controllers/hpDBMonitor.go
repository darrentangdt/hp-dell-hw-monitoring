package controllers

import (
	"log"
	"tms/app/dbs"
	"tms/app/models"
	"tms/app/sockets"
)

func SetCurrentHpSystemState(monitor models.Monitor, overallState *models.OverallState) {
	var system models.HpSystem
	url := "https://" + monitor.Host + "/rest/v1/Systems/1"
	if err := HttpGetRequest(url, monitor, &system); err != nil {
		log.Println(err)
		return
	}
	system.MonitorId = monitor.Id
	dbs.InsertSystem(&system)
	overallState.HpCheckSystemState(system)
}

func SetCurrentHpChassisState(monitor models.Monitor, overallState *models.OverallState) {
	var fanTemperatureJson models.HpFanTemperatureJson
	var fans []models.HpFan
	var temperatures []models.HpTemperature

	url := "https://" + monitor.Host + "/rest/v1/Chassis/1/ThermalMetrics"
	if err := HttpGetRequest(url, monitor, &fanTemperatureJson); err != nil {
		log.Println(err)
		return
	}
	for _, fan := range fanTemperatureJson.Fans {
		fan.MonitorId = monitor.Id
		dbs.InsertFan(&fan)
		fans = append(fans, fan)
	}
	for _, temperature := range fanTemperatureJson.Temperatures {
		temperature.MonitorId = monitor.Id
		dbs.InsertTemperature(&temperature)
		temperatures = append(temperatures, temperature)
	}
	overallState.HpCheckFanState(fans)
	overallState.HpCheckTemperatureState(temperatures)
}

func SetCurrentHpPowerState(monitor models.Monitor, overallState *models.OverallState) {
	var powerJson models.HpPowerJson
	var powers []models.HpPower

	url := "https://" + monitor.Host + "/rest/v1/Chassis/1/PowerMetrics"
	if err := HttpGetRequest(url, monitor, &powerJson); err != nil {
		log.Println(err)
		return
	}
	for _, power := range powerJson.Powers {
		power.MonitorId = monitor.Id
		dbs.InsertPower(&power)
		powers = append(powers, power)
	}
	overallState.HpCheckPowerState(powers)
}

func SetCurrentHpNetworkState(monitor models.Monitor, overallState *models.OverallState) {
	var networkLinkJson models.HpNetworkLinkJson
	var networkJsons []models.HpNetworkJson
	var networks []models.HpNetwork

	url := "https://" + monitor.Host + "/rest/v1/Systems/1/NetworkAdapters"
	if err := HttpGetRequest(url, monitor, &networkLinkJson); err != nil {
		log.Println(err)
		return
	}
	networkJsons = make([]models.HpNetworkJson, len(networkLinkJson.HpNetworkLinks.HpNetworkMembers))
	for i, member := range networkLinkJson.HpNetworkLinks.HpNetworkMembers {
		url := "https://" + monitor.Host + member.Href
		if err := HttpGetRequest(url, monitor, &networkJsons[i]); err != nil {
			log.Println(err)
			return
		}
	}

	for _, networkJson := range networkJsons {
		for _, network := range networkJson.HpNetworks {
			network.MonitorId = monitor.Id
			dbs.InsertNetwork(&network)
			networks = append(networks, network)
		}
	}
	overallState.HpCheckNetworkState(networks)
}

func SetCurrentHpStorageState(monitor models.Monitor, overallState *models.OverallState) {
	var storageLinkJson models.HpStorageLinkJson
	var storages []models.HpStorage

	url := "https://" + monitor.Host + "/rest/v1/Systems/1/SmartStorage/ArrayControllers/0/DiskDrives"
	if err := HttpGetRequest(url, monitor, &storageLinkJson); err != nil {
		log.Println(err)
		return
	}
	storages = make([]models.HpStorage, len(storageLinkJson.HpStorageLinks.HpStorageMembers))
	for i, member := range storageLinkJson.HpStorageLinks.HpStorageMembers {
		url := "https://" + monitor.Host + member.Href
		if err := HttpGetRequest(url, monitor, &storages[i]); err != nil {
			log.Println(err)
			return
		}
		storages[i].MonitorId = monitor.Id
		dbs.InsertStorage(&storages[i])
	}
	overallState.HpCheckDiskState(storages)
}

func InsertCurrentHpState(monitor models.Monitor) {
	overallState := models.OverallState{}
	overallState.CheckMonitorState(monitor)

	SetCurrentHpSystemState(monitor, &overallState)
	SetCurrentHpChassisState(monitor, &overallState)
	SetCurrentHpPowerState(monitor, &overallState)
	SetCurrentHpNetworkState(monitor, &overallState)
	SetCurrentHpStorageState(monitor, &overallState)

	if !overallState.CheckAllStateNotInserted() {
		if components := dbs.GetComponentsIfFaultOccured(overallState); components != "" {
			sockets.Alert(overallState.Host, components)
		}
		dbs.InsertOverallState(overallState)
	}
}

func GetHpMonitorManuf(monitor models.Monitor) string {
	var system models.HpSystem
	url := "https://" + monitor.Host + "/rest/v1/Systems/1"
	if err := HttpGetRequest(url, monitor, &system); err != nil {
		log.Println(err)
	}
	return system.Manufacturer
}
