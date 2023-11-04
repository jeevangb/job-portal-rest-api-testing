package database

import (
	"jeevan/jobportal/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
