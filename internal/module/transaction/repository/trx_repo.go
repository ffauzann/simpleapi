package repository

import (
	"github.com/ffauzann/simpleapi/internal/module/transaction"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) transaction.Repository {
	return &Repository{
		DB: db,
	}
}
