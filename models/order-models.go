package models

import "time"

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductReference"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserReference"`
}
