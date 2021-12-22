package restaurantstorage

import (
	"gorm.io/gorm"
)

// Encapsulation
type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
