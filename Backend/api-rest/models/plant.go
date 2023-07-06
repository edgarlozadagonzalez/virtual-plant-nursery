package models

// Estructura de datos para representar una planta
type Plant struct {
	ID                    uint    `gorm:"primaryKey" json:"id"`
	Name                  string  `gorm:"column:name" json:"name"`
	SurvivalDegree        float64 `gorm:"column:survival_degree" json:"survival_degree"`
	WaterRequired         float64 `gorm:"column:water_required" json:"water_required"`
	WaterSystem           float64 `gorm:"column:water_system" json:"water_system"`
	HydrationLevel        float64 `gorm:"column:hydration_level" json:"hydration_level"`
	NutrientsRequired     float64 `gorm:"column:nutrients_required" json:"nutrients_required"`
	NutrientSystem        float64 `gorm:"column:nutrient_system" json:"nutrient_system"`
	NutritionLevel        float64 `gorm:"column:nutrition_level" json:"nutrition_level"`
	Alive                 bool    `gorm:"column:alive" json:"alive"`
	Img                   string  `gorm:"column:img" json:"img"`
	Descrip               string  `gorm:"column:descrip" json:"descrip"`
	SystemNutrientReserve float64 `gorm:"-" json:"system_nutrient_reserve"`
	SystemWaterReserve    float64 `gorm:"-" json:"system_water_reserve"`
}

func (Plant) TableName() string {
	return "plants"
}
