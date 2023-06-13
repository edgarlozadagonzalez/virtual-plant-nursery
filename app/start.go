package app

import (
	"log"
	"net/http"
	"virtual-plant-nursery-go/app/controllers"
)

func Start() {
	// Rutas de la API
	http.HandleFunc("/plants", controllers.GetAllPlants)
	http.HandleFunc("/plant", controllers.GetPlant)
	http.HandleFunc("/plant/create", controllers.CreatePlant)
	http.HandleFunc("/plant/update", controllers.UpdatePlant)
	http.HandleFunc("/plant/delete", controllers.DeletePlant)
	http.HandleFunc("/plants/water", controllers.AddWaterToPlants)
	http.HandleFunc("/plants/nutrients", controllers.AddNutrientsToPlants)

	// Iniciar el servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

}
