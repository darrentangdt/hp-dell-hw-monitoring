package dbs

import (
	"database/sql"
	"log"
	"os"
	"tms/app/models"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Dbm *gorp.DbMap
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitDB() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tms")
	checkErr(err, "sql.Open failed")
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	Dbm.AddTableWithName(models.OverallState{}, "OverallState")
	Dbm.TruncateTables()
	Dbm.AddTableWithName(models.Monitor{}, "Monitor").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Alarm{}, "Alarm").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.HpNetwork{}, "HpNetwork")
	Dbm.AddTableWithName(models.HpStorage{}, "HpStorage")
	Dbm.AddTableWithName(models.HpFan{}, "HpFan")
	Dbm.AddTableWithName(models.HpTemperature{}, "HpTemperature")
	Dbm.AddTableWithName(models.HpPower{}, "HpPower")
	Dbm.AddTableWithName(models.HpSystem{}, "HpSystem")
	Dbm.AddTableWithName(models.DellNetwork{}, "DellNetwork")
	Dbm.AddTableWithName(models.DellStorage{}, "DellStorage")
	Dbm.AddTableWithName(models.DellFan{}, "DellFan")
	Dbm.AddTableWithName(models.DellTemperature{}, "DellTemperature")
	Dbm.AddTableWithName(models.DellPowerSupply{}, "DellPower")
	Dbm.AddTableWithName(models.DellSystem{}, "DellSystem")
	Dbm.AddTableWithName(models.AlarmLog{}, "AlarmLog")
	err = Dbm.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "tms:", log.Lmicroseconds))
	log.Println("Success gorp initialize")
}
