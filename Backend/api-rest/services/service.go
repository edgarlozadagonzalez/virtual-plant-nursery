package services

import (
	"fmt"
	"time"
	"virtual-plant-nursery/api-rest/models"
	"virtual-plant-nursery/api-rest/repository"
)

var ListPlants []*models.Plant

func LoadListPlants() {
	ListPlants = repository.GetListPlants()
}

func GetPlant(id uint) (models.Plant, error) {
	plant, err := repository.GetPlantById(id)
	return plant, err
}

func CreatePlant(plant models.Plant) bool {
	return repository.PlantCreate(plant)
}

func UpdatePlant(plant models.Plant) bool {
	return repository.PlantUpdate(plant)
}

func DeletePlant(id uint) bool {
	return repository.DeletePlant(id)
}

func AddWaterToPlant(id uint, amount float64) bool {
	return repository.AddWaterToPlant(id, amount)
}

func AddNutrientsToPlant(id uint, amount float64) bool {
	return repository.AddNutrientsToPlant(id, amount)
}

func AddWaterToPlants(amount float64) {
	LoadListPlants()
	for i := range ListPlants {
		AddWaterToPlant(ListPlants[i].ID, amount)
	}
}

func AddNutrientsToPlants(amount float64) {
	LoadListPlants()
	for i := range ListPlants {
		AddNutrientsToPlant(ListPlants[i].ID, amount)
	}
}

// Función que calcula la cantidad de agua en el sistema (AS) de las plantas
func CalculateWaterSystem(plants []*models.Plant) {
	for index, plant := range plants {
		if plant.Alive {
			if plant.SystemWaterReserve == 0 {
				plant.SystemWaterReserve = plant.SurvivalDegree
			}
			if plant.WaterSystem <= 0 {
				plant.WaterSystem = 0
				if plant.SystemWaterReserve > 0 {
					plant.SystemWaterReserve = plant.SystemWaterReserve - (2 / plant.HydrationLevel)
					fmt.Printf("ALERTA: La planta %s está en riesgo extremo de deshidratación.\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.SystemWaterReserve)
				} else {
					plant.SystemWaterReserve = 0
					plant.Alive = false
					ListPlants = append(ListPlants[:index], ListPlants[index+1:]...)
					UpdatePlant(*plant)
					fmt.Printf("ALERTA: La planta %s ha muerto por deshidratación.\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.WaterSystem)
				}
			} else if plant.WaterSystem > 100 {
				if plant.WaterSystem-100 > plant.SurvivalDegree {
					ListPlants = append(ListPlants[:index], ListPlants[index+1:]...)
					plant.Alive = false
					UpdatePlant(*plant)
					fmt.Printf("ALERTA: La planta %s ha muerto por sobrehidratación..\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.WaterSystem)
				} else {
					plant.WaterSystem = plant.WaterSystem - (2 / plant.HydrationLevel)
					UpdatePlant(*plant)
					fmt.Printf("ALERTA: La planta %s está en riesgo de sobre hidratación.\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.WaterSystem)
				}
			} else {
				plant.WaterSystem = plant.WaterSystem - (1 / plant.HydrationLevel)
				UpdatePlant(*plant)
				fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.WaterSystem)
			}
		}
	}
}

// Función que calcula la cantidad de nutrientes en el sistema (NS) de las plantas
func CalculateNutrientSystem(plants []*models.Plant) {

	for index, plant := range plants {
		if plant.Alive {
			if plant.SystemNutrientReserve == 0 {
				plant.SystemNutrientReserve = plant.SurvivalDegree
			}
			if plant.NutrientSystem <= 0 {
				plant.NutrientSystem = 0
				if plant.SystemNutrientReserve > 0 {
					plant.SystemNutrientReserve = plant.SystemNutrientReserve - (2 / plant.NutritionLevel)
					fmt.Printf("ALERTA: La planta %s está en riesgo extremo por pérdida de nutrientes.\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.SystemNutrientReserve)
				} else {
					plant.SystemNutrientReserve = 0
					ListPlants = append(ListPlants[:index], ListPlants[index+1:]...)
					plant.Alive = false
					UpdatePlant(*plant)
					fmt.Printf("ALERTA: La planta %s ha muerto por pérdida de nutrientes.\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.NutrientSystem)
				}
			} else if plant.NutrientSystem > 100 {
				if plant.NutrientSystem-100 > plant.SurvivalDegree {
					ListPlants = append(ListPlants[:index], ListPlants[index+1:]...)
					plant.Alive = false
					UpdatePlant(*plant)
					fmt.Printf("ALERTA: La planta %s ha muerto por ecxeso de nutrientes..\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.NutrientSystem)
				} else {
					plant.NutrientSystem = plant.NutrientSystem - (2 / plant.HydrationLevel)
					UpdatePlant(*plant)
					fmt.Printf("ALERTA: La planta %s está en riesgo por ecxeso de nutrientes.\n", plant.Name)
					fmt.Printf("Planta: %s, Cantidad de nutrientesen el sistema: %.2f%%\n", plant.Name, plant.NutrientSystem)
				}
			} else {
				plant.NutrientSystem = plant.NutrientSystem - (1 / plant.NutritionLevel)
				UpdatePlant(*plant)
				fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.NutrientSystem)
			}
		}
	}
}

// Funcion para enviar datos de la planta por el canal
func GenerateConstantData(plant *models.Plant, dataChannel chan<- models.Plant) {

	for {
		dataChannel <- *plant
		time.Sleep(time.Second)
	}
}
