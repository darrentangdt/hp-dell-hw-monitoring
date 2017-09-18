package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyAlarm(hp, dell models.Monitor) {
	Dbm.Insert(&models.Alarm{0, "test", "01012341234", hp.Host, "", true, true, true, true, true, ""})
	Dbm.Insert(&models.Alarm{0, "test", "01012341234", dell.Host, "", true, true, true, true, true, ""})
}

func InsertAlarm(alarm models.Alarm) {
	if alarm.Service == "Select" {
		alarm.Service = ""
	}

	if err := Dbm.Insert(&alarm); err != nil {
		log.Println(err)
	}
}

func DeleteAlarm(alarmId int64) {
	alarm := models.Alarm{}
	alarm.Id = alarmId
	if _, err := Dbm.Delete(&alarm); err != nil {
		log.Println(err)
	}
}

func UpdateAlarm(alarm models.Alarm) {
	if _, err := Dbm.Update(&alarm); err != nil {
		log.Println(err)
	}
}

func SelectOneAlarmById(id int64) models.Alarm {
	var alarm models.Alarm
	if err := Dbm.SelectOne(&alarm, "select * from Alarm where Id = ?", id); err != nil {
		log.Println(err)
	}
	return alarm
}

func SelectAllAlarmsByCondition(alarm models.Alarm) []models.Alarm {
	var condition string
	if alarm.Host != "" {
		condition += " Host like '%" + alarm.Host + "%' and"
	}
	if alarm.Service != "" && alarm.Service != "Select" {
		condition += " Service like '%" + alarm.Service + "%' and"
	}
	if alarm.UserName != "" {
		condition += " UserName like '%" + alarm.UserName + "%' and"
	}
	if alarm.UserPhone != "" {
		condition += " UserPhone like '%" + alarm.UserPhone + "%' and"
	}
	if alarm.Power == true {
		condition += " Power = 1 and"
	}
	if alarm.Thermal == true {
		condition += " Thermal = 1 and"
	}
	if alarm.Fan == true {
		condition += " Fan = 1 and"
	}
	if alarm.Network == true {
		condition += " Network = 1 and"
	}
	if alarm.Disk == true {
		condition += " Disk = 1 and"
	}

	var alarms []models.Alarm
	if _, err := Dbm.Select(&alarms, GetNewestRecordsQueryByCondition("Alarm", condition)); err != nil {
		log.Println(err)
	}
	return alarms
}
