package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyFan(hp, dell models.Monitor) {
	Dbm.Insert(&models.HpFan{hp.Id, 9, "Fan 1", models.HpFanOem{models.HpFanHp{"System", "HpServerFan.0.9.5"}}, models.HpFanStatus{"OK", "Enabled"}, "Percent", "2017-08-01 10:04:09"})
}

func InsertFan(target interface{}) {
	if err := Dbm.Insert(target); err != nil {
		log.Println(err)
	}
}

func SelectFansNewestState(monitor models.Monitor) interface{} {
	if monitor.Manuf == models.HP {
		var fans []models.HpFan
		if _, err := Dbm.Select(&fans, GetNewestRecordsQueryByMonitorId("HpFan", monitor.Id)); err != nil {
			log.Println(err)
		}
		log.Println(fans)
		return fans
	} else if monitor.Manuf == models.DELL {
		var fans []models.DellFan
		if _, err := Dbm.Select(&fans, GetNewestRecordsQueryByMonitorId("DellFan", monitor.Id)); err != nil {
			log.Println(err)
		}
		return fans
	}

	return nil
}
