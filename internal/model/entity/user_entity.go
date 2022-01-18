package entity

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Name      string
		UserName  string
		Password  string
		CreatedBy int
		UpdatedBy int
	}

	UserWithMerchant struct {
		User     User     `gorm:"embedded"`
		Merchant Merchant `gorm:"embedded;embeddedPrefix:merchant_"`
	}
)
