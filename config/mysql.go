package config

import (
	"campsite_reservation/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection
func InitDBMysql(cfg *AppConfig) *gorm.DB {

	// declare struct config & variable connectionString
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Booking{})
	db.AutoMigrate(&model.Campsite{})
	db.AutoMigrate(&model.Facility{})
	db.AutoMigrate(&model.Availability{})
	db.AutoMigrate(&model.Admin{})
	// db.AutoMigrate(&model.BookingDetail})
	// db.AutoMigrate(&model.Payment{})

}
