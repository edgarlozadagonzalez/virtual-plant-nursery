package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	AlertID    string    `json:"alert_id"`
	Motivo     string    `json:"motivo"`
	DateTime   time.Time `json:"date_time"`
	PlantID    string    `json:"plant_id"`
	Plant      Plant     `gorm:"foreignkey:PlantID"`
}
