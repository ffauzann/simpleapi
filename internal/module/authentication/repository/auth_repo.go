package repository

import (
	"context"

	"github.com/ffauzann/simpleapi/internal/model/entity"
	"github.com/ffauzann/simpleapi/internal/module/authentication"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) authentication.Repository {
	return &Repository{
		DB: db,
	}
}

// GetUserDetailByUsername nodoc. See: irepository.go
func (r *Repository) GetUserDetailByUsername(ctx context.Context, username string) (user entity.UserWithMerchant, err error) {
	// Minimize unused values and prevent sql injection
	query := `
		SELECT 
			u.id,
			u.name,
			u.user_name,
			u.password,
			m.id merchant_id,
			m.merchant_name
		FROM
			Users u
		LEFT JOIN
			Merchants m ON m.user_id = u.id 
		WHERE
			user_name = ?
	`
	err = r.DB.Raw(query, username).Scan(&user).Error
	return
}
