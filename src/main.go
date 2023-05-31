package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_config "github.com/superosystem/bantumanten-backend/src/app/config"
	_dbDriver "github.com/superosystem/bantumanten-backend/src/drivers/mysql"
)

const DEFAULT_PORT = "8080"

func main() {
	// DATABASE SETUP
	cfgDB := _dbDriver.ConfigMySQL{
		MYSQL_USERNAME: _config.GetEnvValue("DB_USERNAME"),
		MYSQL_PASSWORD: _config.GetEnvValue("DB_PASSWORD"),
		MYSQL_HOST:     _config.GetEnvValue("DB_HOST"),
		MYSQL_PORT:     _config.GetEnvValue("DB_PORT"),
		MYSQL_NAME:     _config.GetEnvValue("DB_NAME"),
	}
	db := cfgDB.InitMySQLDatabase()
	_dbDriver.MySQLAutoMigrate(db)

	// ECHO
	e := echo.New()
	var port string = _config.GetEnvValue("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}
	var SERVER_PORT string = fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(SERVER_PORT))
}
