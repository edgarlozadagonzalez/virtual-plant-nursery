package repository

import (
	"log"
	"virtual-plant-nursery/api-rest/database"
	"virtual-plant-nursery/api-rest/models"
)

func GetListPlants() []*models.Plant {
	var plants []*models.Plant
	result := database.Db.Find(&plants)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return plants
}

func PlantCreate(plant models.Plant) bool {
	result := database.Db.Create(&plant)
	if result.Error != nil {
		return false
	}
	return true
}

func GetPlantById(id uint) (models.Plant, error) {
	var plant models.Plant
	result := database.Db.First(&plant, id)
	return plant, result.Error
}

func PlantUpdate(p models.Plant) bool {
	plant, err := GetPlantById(p.ID)

	if err != nil {
		return false
	}

	plant.Name = p.Name
	plant.SurvivalDegree = p.SurvivalDegree
	plant.WaterRequired = p.WaterRequired
	plant.WaterSystem = p.WaterSystem
	plant.HydrationLevel = p.HydrationLevel
	plant.NutrientsRequired = p.NutrientsRequired
	plant.NutrientSystem = p.NutrientSystem
	plant.NutritionLevel = p.NutritionLevel
	plant.Alive = p.Alive

	result := database.Db.Save(&plant)

	if result.Error != nil {
		return false
	}
	return true
}

func DeletePlant(id uint) bool {
	plant, err := GetPlantById(id)
	if err != nil {
		return false
	}
	result := database.Db.Delete(&plant)
	if result.Error != nil {
		return false
	}
	return true
}

func AddWaterToPlant(id uint, waterSystem float64) bool {
	plant, err := GetPlantById(id)
	if err != nil {
		return false
	}

	plant.WaterSystem += (waterSystem * 100) / plant.WaterRequired

	result := database.Db.Save(&plant)
	if result.Error != nil {
		return false
	}

	return true
}

func AddNutrientsToPlant(id uint, nutrientSystem float64) bool {
	plant, err := GetPlantById(id)
	if err != nil {
		return false
	}

	plant.NutrientSystem += (nutrientSystem * 100) / plant.NutrientsRequired

	result := database.Db.Save(&plant)
	if result.Error != nil {
		return false
	}
	return true
}
