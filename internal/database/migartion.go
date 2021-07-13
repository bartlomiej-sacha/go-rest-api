package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sasz94/go-rest-api/internal/comment"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	// automigrate takes model from comment struct and defines all columns in database
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
