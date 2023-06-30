package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host, user, password, dbname string
}

var DBInstance *gorm.DB

func ConnectPostgresDB() error {
	conf := &dbConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	}
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		conf.host,
		strconv.Itoa(32656),
		conf.user,
		conf.password,
		conf.dbname,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	log.Println("DB is Connected!!!!!")

	DBInstance = db
	return nil
}
