package controllers

import (
	"tms/app/dbs"
	"tms/app/models"

	"github.com/revel/revel"
)

type Detail struct {
	*revel.Controller
}

func (c Detail) System(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	target := dbs.SelectSystemNewestState(monitor)

	if monitor.Manuf == models.HP {
		hpSystem, _ := target.(models.HpSystem)
		return c.Render(monitorId, overallState, hpSystem)
	} else if monitor.Manuf == models.DELL {
		dellSystem, _ := target.(models.DellSystem)
		return c.Render(monitorId, overallState, dellSystem)
	}

	return c.Render(monitorId)
}

func (c Detail) Fan(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	target := dbs.SelectFansNewestState(monitor)

	if monitor.Manuf == models.HP {
		hpFans, _ := target.([]models.HpFan)
		return c.Render(monitorId, overallState, hpFans)
	} else if monitor.Manuf == models.DELL {
		dellFans, _ := target.([]models.DellFan)
		return c.Render(monitorId, overallState, dellFans)
	}

	return c.Render(monitorId)
}

func (c Detail) Power(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	target := dbs.SelectPowersNewestState(monitor)

	if monitor.Manuf == models.HP {
		hpPowers, _ := target.([]models.HpPower)
		return c.Render(monitorId, overallState, hpPowers)
	} else if monitor.Manuf == models.DELL {
		dellPowers, _ := target.([]models.DellPowerSupply)
		return c.Render(monitorId, overallState, dellPowers)
	}

	return c.Render(monitorId)
}

func (c Detail) Temperature(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	target := dbs.SelectTemperaturesNewestState(monitor)

	if monitor.Manuf == models.HP {
		hpTemperatures, _ := target.([]models.HpTemperature)
		return c.Render(monitorId, overallState, hpTemperatures)
	} else if monitor.Manuf == models.DELL {
		dellTemperatures, _ := target.([]models.DellTemperature)
		return c.Render(monitorId, overallState, dellTemperatures)
	}

	return c.Render(monitorId)
}

func (c Detail) Network(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	target := dbs.SelectNetworksNewestState(monitor)

	if monitor.Manuf == models.HP {
		hpNetworks, _ := target.([]models.HpNetwork)
		return c.Render(monitorId, overallState, hpNetworks)
	} else if monitor.Manuf == models.DELL {
		dellNetworks, _ := target.([]models.DellNetwork)
		return c.Render(monitorId, overallState, dellNetworks)
	}

	return c.Render(monitorId)
}

func (c Detail) Storage(monitorId int64) revel.Result {
	monitor := dbs.SelectOneMonitorById(monitorId)
	target := dbs.SelectStoragesNewestState(monitor)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)

	if monitor.Manuf == models.HP {
		hpStorages, _ := target.([]models.HpStorage)
		return c.Render(monitorId, overallState, hpStorages)
	} else if monitor.Manuf == models.DELL {
		dellStorages, _ := target.([]models.DellStorage)
		return c.Render(monitorId, overallState, dellStorages)
	}

	return c.Render(monitorId)
}

func (c Detail) PowerControl(monitorId int64) revel.Result {
	monitor := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	overallState := dbs.SelectOneOverallStatesByMonitorId(monitorId)
	return c.Render(monitorId, overallState, monitor)
}
