package app

import (
	"time"
	"virtual-plant-nursery/api-rest/database"
	"virtual-plant-nursery/api-rest/routes"
	"virtual-plant-nursery/api-rest/services"
)

func Start() {
	database.InitDB()
	services.LoadListPlants()

	//Ejecutar las funciones de cálculo cada segundo en goroutines separadas
	go func() {
		for {
			services.CalculateWaterSystem(services.ListPlants)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			services.CalculateNutrientSystem(services.ListPlants)
			time.Sleep(time.Second)
		}
	}()

	routes.Routes()
}
