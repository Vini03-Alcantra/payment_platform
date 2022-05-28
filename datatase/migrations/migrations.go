package migrations

import "gorm.io/gorm"

func RunMigration(db *gorm.DB) {
	db.AutoMigrate()
}
