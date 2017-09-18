package dbs

import (
	"log"
	"strconv"
	"tms/app/models"
)

func InsertDummyMonitor(idx int) (models.Monitor, models.Monitor) {
	hpIdx := strconv.Itoa(idx)
	dellIdx := strconv.Itoa(idx + 200)

	hp := models.Monitor{0, "192.168.0." + hpIdx, "김태엽", "Tman", "test", "test", "HP", "2017-08-01 10:02:03"}
	Dbm.Insert(&hp)

	dell := models.Monitor{0, "192.168.0." + dellIdx, "김태엽", "Tman", "test", "test", "DELL", "2017-08-01 10:02:03"}
	Dbm.Insert(&dell)

	return hp, dell
}

func InsertMonitor(monitor models.Monitor) {
	if monitor.Service == "Select" {
		monitor.Service = ""
	}

	if err := Dbm.Insert(&monitor); err != nil {
		log.Println(err)
	}
}

func DeleteMonitor(monitorId int64) {
	monitor := models.Monitor{}
	monitor.Id = monitorId
	if _, err := Dbm.Delete(&monitor); err != nil {
		log.Println(err)
	}
}

func SelectAllMonitors() []models.Monitor {
	var monitors []models.Monitor
	if _, err := Dbm.Select(&monitors, "select * from Monitor"); err != nil {
		log.Println(err)
	}

	return monitors
}

func SelectAllMonitorsByCondition(monitor models.Monitor) []models.Monitor {
	var condition string
	if monitor.Host != "" {
		condition += " Host like '%" + monitor.Host + "%' and"
	}
	if monitor.Service != "" && monitor.Service != "Select" {
		condition += " Service like '%" + monitor.Service + "%' and"
	}
	if monitor.Manager != "" {
		condition += " Manager like '%" + monitor.Manager + "%' and"
	}
	if monitor.MonitorName != "" {
		condition += " MonitorName like '%" + monitor.MonitorName + "%' and"
	}
	if monitor.MonitorPass != "" {
		condition += " MonitorPass like '%" + monitor.MonitorPass + "%' and"
	}

	var monitors []models.Monitor
	if _, err := Dbm.Select(&monitors, GetNewestRecordsQueryByCondition("Monitor", condition)); err != nil {
		log.Println(err)
	}

	return monitors
}

func SelectOneMonitorById(id int64) models.Monitor {
	var monitor models.Monitor
	if err := Dbm.SelectOne(&monitor, "select * from Monitor where Id = ?", id); err != nil {
		log.Println(err)
	}
	return monitor
}

func SelectMonitorsByService(service string) []models.Monitor {
	var monitors []models.Monitor
	if _, err := Dbm.Select(&monitors, "select * from Monitor where Service = ?", service); err != nil {
		log.Println(err)
	}

	return monitors
}

func SelectOneMonitorByHost(host string) models.Monitor {
	var monitor models.Monitor
	if err := Dbm.SelectOne(&monitor, "select * from Monitor where Host = ?", host); err != nil {
		log.Println(err)
	}

	return monitor
}

func UpdateMonitor(monitor models.Monitor) {
	if _, err := Dbm.Update(&monitor); err != nil {
		log.Println(err)
	}
}
