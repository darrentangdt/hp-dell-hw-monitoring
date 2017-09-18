package controllers

import (
	"log"
	"time"
	"tms/app/dbs"
	"tms/app/models"
)

func InitDBMonitor() {
	go func() {
		log.Println("Intialize DB Monitor")

		for {
			monitors := dbs.SelectAllMonitors()
			for _, monitor := range monitors {
				go GetMonitorState(monitor)
			}

			time.Sleep(30 * time.Second)
		}
	}()
}

func GetMonitorState(monitor models.Monitor) {
	switch monitor.Manuf {
	case models.HP:
		InsertCurrentHpState(monitor)
	case models.DELL:
		DellInsertCurrentState(monitor)
	default:
		SetMonitorManuf(monitor)
	}
}

func SetMonitorManuf(monitor models.Monitor) {
	if GetHpMonitorManuf(monitor) != "" {
		monitor.Manuf = models.HP
	} else if DellGetMonitorManuf(monitor) != "" {
		monitor.Manuf = models.DELL
	} else {
		return
	}

	dbs.UpdateMonitor(monitor)
	go GetMonitorState(monitor)
}
