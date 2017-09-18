package dbs

import "strconv"

var dbTableGroupMap = map[string]string{
	"HpFan":           "FanName",
	"HpPower":         "BayNumber",
	"HpTemperature":   "Name",
	"HpSystem":        "SerialNumber",
	"HpNetwork":       "MacAddress",
	"HpStorage":       "SerialNumber",
	"OverallState":    "Host",
	"DellFan":         "FanName",
	"DellPower":       "MemberID",
	"DellTemperature": "Name",
	"DellSystem":      "SerialNumber",
	"DellNetwork":     "MacAddress",
	"DellStorage":     "Name",
	"Monitor":         "Id",
	"Alarm":           "Id",
}

func GetNewestRecordsQueryByMonitorId(table string, monitorId int64) string {
	return "select * from " + table +
		" where MonitorId = " + strconv.FormatInt(monitorId, 10) +
		" and CreatedAt in " +
		"(select max(CreatedAt) from " + table +
		" where MonitorId = " + strconv.FormatInt(monitorId, 10) +
		" group by " + dbTableGroupMap[table] + ")"
}

func GetNewestRecordsQueryByCondition(table string, condition string) string {
	return "select * from " + table +
		" where" + condition +
		" CreatedAt in " +
		"(select max(CreatedAt) from " + table +
		" group by " + dbTableGroupMap[table] + ")"
}

func GetRecordsQueryByCondition(table string, condition string) string {
	if condition != "" {
		return "select * from " + table + " where" + condition
	} else {
		return "select * from " + table
	}
}
