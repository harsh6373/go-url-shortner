package model

import "time"

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	Slug      string `gorm:"uniqueIndex"`
	Original  string
	CreatedAt time.Time
	ExpiresAt *time.Time
}

type Click struct {
	ID        uint `gorm:"primaryKey"`
	Slug      string
	UserAgent string
	Timestamp time.Time
}
