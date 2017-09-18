package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyStorage(hp, dell models.Monitor) {
	Dbm.Insert(&models.HpStorage{hp.Id, 572325, 0, "SAS", "1I:3:4", "ControllerPort:Box:Bay", 40, "EG0600FCVBK", "S0M53XPC0000K5363XDX", models.HpStorageStatus{"OK", "Enabled"}, "2017-08-01 10:04:21"})
}

func InsertStorage(target interface{}) {
	if err := Dbm.Insert(target); err != nil {
		log.Println(err)
	}
}

func SelectStoragesNewestState(monitor models.Monitor) interface{} {
	if monitor.Manuf == models.HP {
		var storages []models.HpStorage
		if _, err := Dbm.Select(&storages, GetNewestRecordsQueryByMonitorId("HpStorage", monitor.Id)); err != nil {
			log.Println(err)
		}
		return storages
	} else if monitor.Manuf == models.DELL {
		var storages []models.DellStorage
		if _, err := Dbm.Select(&storages, GetNewestRecordsQueryByMonitorId("DellStorage", monitor.Id)); err != nil {
			log.Println(err)
		}
		return storages
	}

	return nil
}
