package entity

import "time"

type User struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:enum('admin','campaign_owner','donor');default:donor"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
