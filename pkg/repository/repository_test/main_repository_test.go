package repository_test

import (
	"clean-arch-hicoll/config"
	"clean-arch-hicoll/shared/db"
	"database/sql"
	"testing"
)

var dbConn *sql.DB

func TestMain(m *testing.M) {
	conf := config.NewConfigurationPath("../../../config/config_test.json")
	dbConn = db.NewInstanceDb(conf)

	m.Run()
}
