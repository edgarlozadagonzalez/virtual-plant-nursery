package services

import (
	"fmt"
	"virtual-plant-nursery-go/app/models"
)

// Función que calcula la cantidad de agua en el sistema (AS) de las plantas
func CalculateWaterSystem(plants []*models.Plant) {
	for _, plant := range plants {
		if plant.SystemWaterReserve == 0 {
			plant.SystemWaterReserve = plant.SurvivalDegree
		}
		if plant.WaterSystem <= 0 {
			if plant.SystemWaterReserve > 0 {
				plant.SystemWaterReserve = plant.SystemWaterReserve - (2 / plant.HydrationLevel)
				fmt.Printf("ALERTA: La planta %s está en riesgo extremo de deshidratación.\n", plant.Name)
				fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.SystemWaterReserve)
			} else {
				fmt.Printf("ALERTA: La planta %s ha muerto por deshidratación.\n", plant.Name)
				fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.WaterSystem)
			}
		} else {
			plant.WaterSystem = plant.WaterSystem - (1 / plant.HydrationLevel)
			fmt.Printf("Planta: %s, Cantidad de agua en el sistema: %.2f%%\n", plant.Name, plant.WaterSystem)
		}
	}
}

// Función que calcula la cantidad de nutrientes en el sistema (NS) de las plantas
func CalculateNutrientSystem(plants []*models.Plant) {
	for _, plant := range plants {
		if plant.SystemNutrientReserve == 0 {
			plant.SystemNutrientReserve = plant.SurvivalDegree
		}
		if plant.NutrientSystem <= 0 {
			if plant.SystemNutrientReserve > 0 {
				plant.SystemNutrientReserve = plant.SystemNutrientReserve - (2 / plant.NutritionLevel)
				fmt.Printf("ALERTA: La planta %s está en riesgo extremo por pérdida de nutrientes.\n", plant.Name)
				fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.SystemNutrientReserve)
			} else {
				fmt.Printf("ALERTA: La planta %s ha muerto por pérdida de nutrientes.\n", plant.Name)
				fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.NutrientSystem)
			}
		} else {
			plant.NutrientSystem = plant.NutrientSystem - (1 / plant.NutritionLevel)
			fmt.Printf("Planta: %s, Cantidad de nutrientes en el sistema: %.2f%%\n", plant.Name, plant.NutrientSystem)
		}
	}
}
