package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	AlertID    string    `json:"alert_id"`
	Reason     string    `json:"reason"`
	DateTime   time.Time `json:"date_time"`
	//PlantID    string    `json:"plant_id"`
	Plant      Plant     `gorm:"foreignkey:PlantID"`
}

func (Alert) TableName() string {
	return "alerts"
}