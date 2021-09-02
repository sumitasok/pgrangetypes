package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Tstzrgt table
type Tstzrgt struct {
	Room int    `gorm:""`
	Dttm string `gorm:""`
}

func (t Tstzrgt) TableName() string {
	return "tstzrgt"
}

func main() {
	var config = struct {
		PgHost     string
		PgPort     string
		PgDbname   string
		PgUser     string
		PgPassword string
		PgSslmode  string
	}{
		PgHost:     "localhost",
		PgPort:     "5432",
		PgDbname:   "dbtest",
		PgUser:     "sumitasok",
		PgPassword: "nest-skyward-decision",
		PgSslmode:  "disable",
	}

	var pgConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.PgHost, config.PgPort, config.PgUser, config.PgDbname, config.PgPassword, config.PgSslmode)

	db, err := gorm.Open(postgres.Open(pgConnStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	//db, err := gorm.Open(postgres.Open(pgConnStr), &gorm.Config{})
	if err != nil {
		log.Panicln(err.Error())
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	tsRow := Tstzrgt{}
	tx := db.First(&tsRow)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	log.Println("First row: ", tsRow)

	ts1 := Tstzrgt{Room: 107, Dttm: "[2010-01-02 01:01:02+5:30, 2010-01-03 02:04:07+5:30]"}
	tx = db.Create(&ts1)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	log.Println("program run successfully")
}
