package routes

import (
	"net/http"
	"virtual-plant-nursery/app/controllers"
)

func Routes() {
	// Rutas de la API
	http.HandleFunc("/plants", controllers.GetAllPlants)
	http.HandleFunc("/plant", controllers.GetPlant)
	http.HandleFunc("/plant/create", controllers.CreatePlant)
	http.HandleFunc("/plant/update", controllers.UpdatePlant)
	http.HandleFunc("/plant/delete", controllers.DeletePlant)
	http.HandleFunc("/plants/water", controllers.AddWaterToPlants)
	http.HandleFunc("/plants/nutrients", controllers.AddNutrientsToPlants)
}
