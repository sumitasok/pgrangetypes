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

	tstzRangeWrong, err := pgtype.NewTstzrange('[', time.Now(), time.Now().Add(-1*time.Hour), ')')
	if err != nil {
		log.Println(err)
	}

	tsWrongdatetime := Tstzrgt{Room: 1078, Dttm: tstzRangeWrong}
	tx = db.Create(&tsWrongdatetime)
	if tx.Error != nil {
		log.Println("Error:", tx.Error)
	}

	log.Printf("ts2: %s", ts2.Dttm.ToString())

	var tss1 []Tstzrgt
	tx = db.Where("dttm @> ?::timestamptz", "2010-01-03 02:04:07+5:30").Find(&tss1)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	for i, tstzrgt := range tss1 {
		log.Printf("tss1: %d %d %s", i, tstzrgt.Room, tstzrgt.Dttm.ToString())
	}

	//select * from tstzrgt where dttm @> tstzrange('2000-01-02 12:00:00+5:30', '2000-01-03 04:00:00+5:30');
	var tss2 []Tstzrgt
	tx = db.Where("dttm @> tstzrange(?, ?)", "2000-01-02 12:00:00+5:30", "2000-01-03 04:00:00+5:30").Find(&tss2)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	for i, tstzrgt := range tss2 {
		log.Printf("tss2: %d %d %s", i, tstzrgt.Room, tstzrgt.Dttm.ToString())
	}

	//select * from tstzrgt where dttm @> tstzrange('2000-01-02 12:00:00+5:30', '2000-01-03 04:00:00+5:30');
	var tss3 []Tstzrgt
	tx = db.Where("dttm && tstzrange(?, ?)", "2000-01-02 12:00:00+5:30", "2000-01-03 04:00:00+5:30").Find(&tss3)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	for i, tstzrgt := range tss3 {
		log.Printf("tss3: %d %d %s", i, tstzrgt.Room, tstzrgt.Dttm.ToString())
	}

	// select * from tstzrgt where NOT dttm && tstzrange('2000-01-02 12:00:00+5:30', '2000-01-03 04:00:00+5:30');
	var tss4 []Tstzrgt
	tx = db.Where("NOT dttm && tstzrange(?, ?)", "2000-01-02 12:00:00+5:30", "2000-01-03 04:00:00+5:30").Find(&tss4)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	for i, tstzrgt := range tss4 {
		log.Printf("tss4: %d %d %s", i, tstzrgt.Room, tstzrgt.Dttm.ToString())
	}

	//select unnest(array[100, 101, 102, 103, 104, 107])
	//    EXCEPT
	//    select room from tstzrgt where dttm && tstzrange('2000-01-02 12:00:00+5:30', '2000-01-03 04:00:00+5:30');

	type dummy struct {
		Room int
	}

	var tss5 []dummy
	tx = db.Raw("select unnest(array[100, 101, 102, 103, 104, 107]) as room EXCEPT select room from tstzrgt where dttm && tstzrange(?, ?)", "2000-01-02 12:00:00+5:30", "2000-01-03 04:00:00+5:30").Scan(&tss5)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	for i, tstzrgt := range tss5 {
		log.Printf("tss5: %d %d", i, tstzrgt.Room)
	}

	log.Println("program run successfully")
}
