package connection

import (
	"fmt"
	"log"
	"rest_api_sederhana/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(conf *config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Host, conf.User, conf.Pass, conf.Name, conf.Port, conf.Tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err.Error())
	}

	// Migrate the schema // jika ingin otmatis, tapi ini jarang dipakai
	// db.AutoMigrate(&domain.Customer{}) // here replace domain.Customer with your actual model
	// db.AutoMigrate(&domain.Product{}) // here replace domain.Product with your actual model
	// db.AutoMigrate(&domain.Order{}) // here replace domain.Order with your actual model
	// db.AutoMigrate(&domain.OrderItem{}) // here replace domain.OrderItem with your actual model
	// db.AutoMigrate(&domain.User{}) // here replace domain.User with your actual model
	// db.AutoMigrate(&domain.Role{}) // here replace domain.Role with your actual model
	// db.AutoMigrate(&domain.Permission{}) // here replace
	return db
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
