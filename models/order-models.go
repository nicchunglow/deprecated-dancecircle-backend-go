package models

import "time"

type Order struct {
	ID               uint `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time
	ProductReference int     `json:"product_id"`
	Product          Product `gorm:"foreignKey:ProductReference"`
	UserReference    int     `json:"user_id"`
	User             User    `gorm:"foreignKey:UserReference"`
}
