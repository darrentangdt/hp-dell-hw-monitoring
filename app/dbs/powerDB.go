package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyPower(hp, dell models.Monitor) {
	Dbm.Insert(&models.HpPower{hp.Id, "720478-B21", "HpServerPowerSupply.0.9.5", models.HpPowerOem{models.HpPowerHp{0}}, models.HpPowerStatus{"Warning", "Enabled"}, "2017-08-01 10:04:09"})
}

func InsertPower(target interface{}) {
	if err := Dbm.Insert(target); err != nil {
		log.Println(err)
	}
}

func SelectPowersNewestState(monitor models.Monitor) interface{} {
	if monitor.Manuf == models.HP {
		var powers []models.HpPower
		if _, err := Dbm.Select(&powers, GetNewestRecordsQueryByMonitorId("HpPower", monitor.Id)); err != nil {
			log.Println(err)
		}
		return powers
	} else if monitor.Manuf == models.DELL {
		var powers []models.DellPowerSupply
		if _, err := Dbm.Select(&powers, GetNewestRecordsQueryByMonitorId("DellPower", monitor.Id)); err != nil {
			log.Println(err)
		}
		return powers
	}

	return nil
}
