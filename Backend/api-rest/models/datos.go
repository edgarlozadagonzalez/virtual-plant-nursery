package models

// Variable para almacenar las plantas ,ejemplo para test provicional.
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
