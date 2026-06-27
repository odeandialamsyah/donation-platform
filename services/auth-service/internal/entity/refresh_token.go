package entity

import "time"

type RefreshToken struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	UserID    string `gorm:"type:char(36);index;not null"`
	Token     string `gorm:"type:text;not null"`
	ExpiredAt time.Time
	CreatedAt time.Time
}
