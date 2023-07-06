package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Record struct {
	gorm.Model
	RecordID string    `json:"record_id"`
	Reason   string    `json:"reason"`
	DateTime time.Time `json:"date_time"`
	PlantID  string    `json:"plant_id"`
	Plant    Plant     `gorm:"foreignkey:PlantID"`
}

func (Record) TableName() string {
	return "records"
}