package models

import (
	"github.com/jinzhu/gorm"
)

// Estructura de datos para representar una planta
type Plant struct {
	gorm.Model
	//ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	SurvivalDegree        float64 `json:"survival_degree"`
	WaterRequired         float64 `json:"water_required"`
	WaterSystem           float64 `json:"water_system"`
	HydrationLevel        float64 `json:"hydration_level"`
	NutrientsRequired     float64 `json:"nutrients_required"`
	NutrientSystem        float64 `json:"nutrient_system"`
	NutritionLevel        float64 `json:"nutrition_level"`
	Alive	              bool    `json:"alive"`
}

func (Plant) TableName() string {
	return "plants"
}