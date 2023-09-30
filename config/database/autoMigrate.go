package database

import (
	"TRS/app/models"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Team{},
	)

	return err
}
