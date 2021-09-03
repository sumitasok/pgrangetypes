package main

import (
	"fmt"
	"log"
	pgtype "pgrangetypes/lib"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Tstzrgt table
type Tstzrgt struct {
	Room int               `gorm:""`
	Dttm *pgtype.Tstzrange `gorm:""`
}

type TstzrgtDummy struct {
	Room int    `gorm:""`
	Dttm string `gorm:""`
}

func (t Tstzrgt) TableName() string {
	return "tstzrgt"
}

func (t TstzrgtDummy) TableName() string {
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

	tsRow := TstzrgtDummy{}
	tx := db.First(&tsRow)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	log.Println("First row: ", tsRow)

	tstzRange, err := pgtype.NewTstzrange('[', time.Now(), time.Now().Add(1*time.Hour), ')')
	if err != nil {
		log.Println(err)
	}

	ts1 := Tstzrgt{Room: 1077, Dttm: tstzRange}
	tx = db.Create(&ts1)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	ts2 := Tstzrgt{}
	tx = db.Last(&ts2)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	log.Printf("prefix: %s", ts2.Dttm.ToString())

	log.Println("program run successfully")
}
