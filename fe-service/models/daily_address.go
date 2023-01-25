package models

import (
	"gorm.io/gorm"
)

type DailyAddress struct {
	Id        uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AddressId uint64 `json:"address_id"`
	Balance   string `gorm:"type:varchar(256)"  json:"balance"`
	*gorm.Model
}
