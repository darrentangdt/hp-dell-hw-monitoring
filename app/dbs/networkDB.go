package dbs

import (
	"log"
	"tms/app/models"
)

func InsertDummyNetwork(hp, dell models.Monitor) {
	Dbm.Insert(&models.HpNetwork{hp.Id, "3C:A8:2A:21:2F:4C", 0, models.HpNetworkStatus{"Warning", "Disabled"}, "2017-08-01 10:04:12"})
}

func InsertNetwork(target interface{}) {
	if err := Dbm.Insert(target); err != nil {
		log.Println(err)
	}
}

func SelectNetworksNewestState(monitor models.Monitor) interface{} {
	if monitor.Manuf == models.HP {
		var networks []models.HpNetwork
		if _, err := Dbm.Select(&networks, GetNewestRecordsQueryByMonitorId("HpNetwork", monitor.Id)); err != nil {
			log.Println(err)
		}
		return networks
	} else if monitor.Manuf == models.DELL {
		var networks []models.DellNetwork
		if _, err := Dbm.Select(&networks, GetNewestRecordsQueryByMonitorId("DellNetwork", monitor.Id)); err != nil {
			log.Println(err)
		}
		return networks
	}

	return nil
}
