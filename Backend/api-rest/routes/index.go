package routes

import (
	"virtual-plant-nursery/api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {

	router := gin.Default()
	router.GET("/plants", controllers.GetAllPlants)
	router.GET("/plant/:id", controllers.GetPlant)
	router.POST("/plant/create", controllers.CreatePlant)
	router.PUT("/plant/update/:id", controllers.UpdatePlant)
	router.DELETE("/plant/delete/:id", controllers.DeletePlant)

	router.POST("/plant/:id/water/:amount", controllers.AddWaterToPlant)
	router.POST("/plant/:id/nutrients/:amount", controllers.AddNutrientsToPlant)

	router.POST("/plants/water/:amount", controllers.AddWaterToPlants)
	router.POST("/plants/nutrients/:amount", controllers.AddNutrientsToPlants)
	router.Run("localhost:8080")

}
