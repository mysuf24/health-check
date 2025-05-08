package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // variable untuk diakses dari controller

func ConnectDB() {
	dns := "host=localhost user=postgres password=123 port=5432 dbname=db_merchant sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal untuk connect ke database: %v", err)
	}
	log.Println("[INFO] Berhasil koneksi ke database")
	
	DB = database
}
func PingDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
