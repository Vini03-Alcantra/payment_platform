package migrations

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.DebitCard{})
}
