package entities

import "time"

type (
	Admin struct {
		ID        string    `gorm:"primaryKey;varchar(64);"`
		Item      []Item    `gorm:"foreignKey:AdminID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Email     string    `gorm:"type:varchar(128);unique;not null;"`
		Name      string    `gorm:"type:varchar(128);not null;"`
		Avatar    string    `gorm:"type:varchar(256);not null;"`
		CreatedAt time.Time `grom:"not null;autoCreateTime;"`
		UpdatedAt time.Time `grom:"not null;autoUpdateTime;"`
	}
)
