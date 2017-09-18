package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyOverallState(hp, dell models.Monitor) {
	Dbm.Insert(&models.OverallState{hp.Id, hp.Host, hp.Service, hp.Manager, "On", "OK", "OK", "OK", "OK", "OK", "OK", ""})
	Dbm.Insert(&models.OverallState{dell.Id, dell.Host, dell.Service, dell.Manager, "On", "OK", "OK", "OK", "OK", "OK", "OK", ""})
}

func InsertOverallState(overallState models.OverallState) {
	if err := Dbm.Insert(&overallState); err != nil {
		log.Println(err)
	}
}

func DeleteOverallState(monitorId int64) {
	if _, err := Dbm.Exec("delete from OverallState where monitorId=?", monitorId); err != nil {
		log.Println(err)
	}
}

func SelectAllOverallStatesByCondition(overallState models.OverallState) []models.OverallState {
	var condition string
	if overallState.Host != "" {
		condition += " Host like '%" + overallState.Host + "%' and"
	}
	if overallState.Service != "" && overallState.Service != "Select" {
		condition += " Service like '%" + overallState.Service + "%' and"
	}
	if overallState.Manager != "" {
		condition += " Manager like '%" + overallState.Manager + "%' and"
	}
	if overallState.Power != "" && overallState.Power != "All" {
		condition += " Power = '" + overallState.Power + "' and"
	}
	if overallState.State != "" && overallState.State != "All" {
		condition += " State = '" + overallState.State + "' and"
	}

	var overallStates []models.OverallState
	if _, err := Dbm.Select(&overallStates, GetNewestRecordsQueryByCondition("OverallState", condition)); err != nil {
		log.Println(err)
	}
	return overallStates
}

func SelectOneOverallStatesByMonitorId(monitorId int64) models.OverallState {

	var overallState models.OverallState
	if err := Dbm.SelectOne(&overallState, GetNewestRecordsQueryByMonitorId("OverallState", monitorId)); err != nil {
		log.Println(err)
	}
	return overallState
}

func GetComponentsIfFaultOccured(overallState models.OverallState) string {
	var components string
	var prevOver models.OverallState
	prevOver = SelectOneOverallStatesByMonitorId(overallState.MonitorId)

	if prevOver.PowerSupply != overallState.PowerSupply && overallState.PowerSupply != models.OK {
		components += "PowerSupply "
	}
	if prevOver.Fan != overallState.Fan && overallState.Fan != models.OK {
		components += "Fan "
	}
	if prevOver.Temperature != overallState.Temperature && overallState.Temperature != models.OK {
		components += "Temperature "
	}
	if prevOver.Network != overallState.Network && overallState.Network != models.OK {
		components += "Network "
	}
	if prevOver.Disk != overallState.Disk && overallState.Disk != models.OK {
		components += "Disk "
	}
	return components
}
