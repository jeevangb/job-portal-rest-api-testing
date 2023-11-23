package database

import (
	"fmt"
	"jeevan/jobportal/config"
	"jeevan/jobportal/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// dsn := "host=postgres user=postgres password=root dbname=updated-job-portal port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.Sslmode, cfg.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {

		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {

		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Jobs{})
	if err != nil {

		return nil, err
	}
	return db, nil
}
