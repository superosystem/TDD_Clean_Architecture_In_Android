package mysql

import (
	"fmt"
	"github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ConfigMySQL struct {
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_NAME     string
	MYSQL_HOST     string
	MYSQL_PORT     string
}

func (cfg *ConfigMySQL) InitMySQLDatabase() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MYSQL_USERNAME,
		cfg.MYSQL_PASSWORD,
		cfg.MYSQL_HOST,
		cfg.MYSQL_PORT,
		cfg.MYSQL_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func MySQLAutoMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(&users.User{})
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}
