package entity

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name string `json:"name"`
}

type MerchantGross struct {
	Date         string
	MerchantID   int
	MerchantName string
	Gross        int
}
