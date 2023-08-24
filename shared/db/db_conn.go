package db

import (
	"clean-arch-hicoll/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewInstanceDb() *sql.DB {
	conf := config.NewConfiguration()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBName))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
