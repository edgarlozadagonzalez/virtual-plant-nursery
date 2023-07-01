package app

import (
	"time"
	"virtual-plant-nursery/api-rest/models"
	"virtual-plant-nursery/api-rest/routes"
	"virtual-plant-nursery/api-rest/services"
)

func Start() {
	// Ejecutar las funciones de cálculo cada segundo en goroutines separadas
	go func() {
		for {
			services.CalculateWaterSystem(models.ListPlants)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			services.CalculateNutrientSystem(models.ListPlants)
			time.Sleep(time.Second)
		}
	}()

	routes.Routes()
}
