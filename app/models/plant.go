package models

// Estructura de datos para representar una planta
type Plant struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	SurvivalDegree        float64 `json:"survival_degree"`
	WaterRequired         float64 `json:"water_required"`
	WaterSystem           float64 `json:"water_system"`
	HydrationLevel        float64 `json:"hydration_level"`
	NutrientsRequired     float64 `json:"nutrients_required"`
	NutrientSystem        float64 `json:"nutrient_system"`
	NutritionLevel        float64 `json:"nutrition_level"`
	SystemNutrientReserve float64 `json:"-"`
	SystemWaterReserve    float64 `json:"-"`
}

// Variable global para almacenar las plantas ,ejemplo para test provicional.
var ListPlants = []*Plant{
	{
		ID:                "1",
		Name:              "Rosa",
		SurvivalDegree:    20,
		WaterRequired:     0.2,
		WaterSystem:       100,
		HydrationLevel:    50,
		NutrientsRequired: 12,
		NutrientSystem:    100,
		NutritionLevel:    60,
	},
	{
		ID:                "2",
		Name:              "Orquídea",
		SurvivalDegree:    0.2,
		WaterRequired:     0.4,
		WaterSystem:       0.2,
		HydrationLevel:    60,
		NutrientsRequired: 10,
		NutrientSystem:    10,
		NutritionLevel:    70,
	},
}
