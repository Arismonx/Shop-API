package entities

import "time"

type (
	Item struct {
		ID          uint64    `gorm:"primaryKey;autoIncrement;"`
		AdminID     *string   `gorm:"type:varchar(64);"`
		Name        string    `gorm:"type:varchar(64);unique;not null;"`
		Description string    `grom:"type:varchar(128);not null;"`
		Picture     string    `grom:"type:varchar(256);not null;"`
		Price       uint      `grom:"not null"`
		IsArchive   bool      `grom:"not null;default:false;"`
		CreatedAt   time.Time `grom:"not null;autoCreateTime;"`
		UpdatedAt   time.Time `grom:"not null;autoUpdateTime;"`
	}
)
