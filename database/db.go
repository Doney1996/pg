package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db DbUtil

type DbUtil struct {
	db *sqlx.DB
}

func (d *DbUtil) SetDb(db *sqlx.DB) {
	d.db = db
}

func (d DbUtil) GetDb() *sqlx.DB {
	return d.db
}

func init(){
	database, err := sqlx.Open("mysql", "root:admin1984@tcp(114.96.105.111:3306)/pg")
	if err != nil {
		log.Fatal("open mysql failed,", err)
	}
	log.Println(database, "database init success")
	dataBase := new(DbUtil)
	dataBase.SetDb(database)
	Db = *dataBase
}
