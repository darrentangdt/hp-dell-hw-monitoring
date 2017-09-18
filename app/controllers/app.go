package controllers

import (
	"tms/app/dbs"
	"tms/app/models"
	"tms/app/sockets"

	"golang.org/x/net/websocket"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Monitor(overallState models.OverallState) revel.Result {
	overallStates := dbs.SelectAllOverallStatesByCondition(overallState)
	return c.Render(overallStates)
}

func (c App) MonitorSocket(ws *websocket.Conn) revel.Result {
	subscription := sockets.Subscribe()
	defer subscription.Cancel()

	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				return nil
			}
		}
	}
	return nil
}

func (c App) Server(monitor models.Monitor) revel.Result {
	monitors := dbs.SelectAllMonitorsByCondition(monitor)
	return c.Render(monitors)
}

func (c App) AddServer(monitor models.Monitor) revel.Result {
	dbs.InsertMonitor(monitor)
	monitor = dbs.SelectOneMonitorByHost(monitor.Host)
	go SetMonitorManuf(monitor)
	return c.Redirect(App.Server)
}

func (c App) DelServer(monitorId int64) revel.Result {
	dbs.DeleteMonitor(monitorId)
	dbs.DeleteOverallState(monitorId)
	return c.Redirect(App.Server)
}

func (c App) ControlServer(monitorHost string, pwrCtl string) revel.Result {
	monitor := dbs.SelectOneMonitorByHost(monitorHost)
	ControlHost(monitor, pwrCtl)
	return c.Redirect("/detail/powercontrol/%d", monitor.Id)
}

func (c App) UpdateServer(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	return c.Render(monitor)
}
func (c App) UpdateSubmit(monitor models.Monitor) revel.Result {
	dbs.UpdateMonitor(monitor)
	return c.Redirect(App.Server)
}

func (c App) Alarm(alarm models.Alarm) revel.Result {
	alarms := dbs.SelectAllAlarmsByCondition(alarm)
	return c.Render(alarms)
}

func (c App) AddAlarm(alarm models.Alarm) revel.Result {
	dbs.InsertAlarm(alarm)
	return c.Redirect(App.Alarm)
}

func (c App) DelAlarm(alarmId int64) revel.Result {
	dbs.DeleteAlarm(alarmId)
	return c.Redirect(App.Alarm)
}

func (c App) UpdateAlarmWeb(alarmId int64) revel.Result {
	alarm := dbs.SelectOneAlarmById(alarmId)
	return c.Render(alarm)
}

func (c App) UpdateAlarmSubmit(alarm models.Alarm) revel.Result {
	dbs.UpdateAlarm(alarm)
	return c.Redirect(App.Alarm)
}

func (c App) AlarmLogView(alarmlog models.AlarmLog) revel.Result {
	alarmlogs := dbs.SelectAllAlarmLogsByCondition(alarmlog)
	return c.Render(alarmlogs)
}

func (c App) Dummy() revel.Result {
	for idx := 400; idx < 450; idx++ {
		hp, dell := dbs.InsertDummyMonitor(idx)
		dbs.InsertDummyOverallState(hp, dell)
		dbs.InsertDummyAlarm(hp, dell)
		dbs.InsertDummySystem(hp, dell)
		dbs.InsertDummyPower(hp, dell)
		dbs.InsertDummyFan(hp, dell)
		dbs.InsertDummyStorage(hp, dell)
		dbs.InsertDummyNetwork(hp, dell)
		dbs.InsertDummyTemperature(hp, dell)
	}
	return c.Redirect(App.Monitor)
}
