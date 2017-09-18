package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummySystem(hp, dell models.Monitor) {
	Dbm.Insert(&models.HpSystem{hp.Id, models.HpBios{models.HpCurrent{"P89 v1.40 (05/06/2015)"}}, "HP", "ProLiant DL380 Gen9", "Computer System", "Off", models.HpProcessors{1, "Intel(R) Xeon(R) CPU E5-2640 v3 @ 2.60GHz", models.HpProcessorStatus{"OK"}}, "719064-B21", "SGH528YA4M", models.HpSystemStatus{"Warning", "Disabled"}, "2017-08-01 10:04:08"})
}

func InsertSystem(target interface{}) {
	if err := Dbm.Insert(target); err != nil {
		log.Println(err)
	}
}

func SelectSystemNewestState(monitor models.Monitor) interface{} {
	if monitor.Manuf == models.HP {
		var system models.HpSystem
		if err := Dbm.SelectOne(&system,
			GetNewestRecordsQueryByMonitorId("HpSystem", monitor.Id)); err != nil {
			log.Println(err)
		}
		return system
	} else if monitor.Manuf == models.DELL {
		var system models.DellSystem
		if err := Dbm.SelectOne(&system,
			GetNewestRecordsQueryByMonitorId("DellSystem", monitor.Id)); err != nil {
			log.Println(err)
		}
		return system
	}

	return nil
}
