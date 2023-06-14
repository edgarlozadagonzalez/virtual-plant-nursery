package app

import (
	"log"
	"net/http"
	"time"
	"virtual-plant-nursery-go/app/models"
	"virtual-plant-nursery-go/app/routes"
	"virtual-plant-nursery-go/app/services"
)

func Start() {

	routes.Routes()

	// Iniciar el servidor en el puerto 8080
	log.Println("Iniciando Servidor en el puerto 8080")

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

	log.Fatal(http.ListenAndServe(":8080", nil))

}
