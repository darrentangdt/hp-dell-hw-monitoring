package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyTemperature(hp, dell models.Monitor) {
	Dbm.Insert(&models.HpTemperature{hp.Id, "Ambient", 23, "01-Inlet Ambient", 1, models.HpTemperatureOem{models.HpTemperatureHp{1, 1, "HpSeaOfSensors.0.9.5"}}, models.HpTemperatureStatus{"OK", "Enabled"}, "Celsius", "2017-08-01 10:04:09"})
}

func InsertTemperature(target interface{}) {
	if err := Dbm.Insert(target); err != nil {
		log.Println(err)
	}
}

func SelectTemperaturesNewestState(monitor models.Monitor) interface{} {
	if monitor.Manuf == models.HP {
		var temperatures []models.HpTemperature
		if _, err := Dbm.Select(&temperatures, GetNewestRecordsQueryByMonitorId("HpTemperature", monitor.Id)); err != nil {
			log.Println(err)
		}
		return temperatures
	} else if monitor.Manuf == models.DELL {
		var temperatures []models.DellTemperature
		if _, err := Dbm.Select(&temperatures, GetNewestRecordsQueryByMonitorId("DellTemperature", monitor.Id)); err != nil {
			log.Println(err)
		}
		return temperatures
	}

	return nil
}
