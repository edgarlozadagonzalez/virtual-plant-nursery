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

	//router.POST("/plants/water/:amount", database.AddWaterToPlants)
	//router.POST("/plants/nutrients/:amount", database.AddNutrientsToPlants)

	// router.GET("/alerts", database.AlertGet)
	// router.GET("/alert/:id", database.AlertGetByID)
	// router.POST("/alert/create", database.AlertCreate)
	// router.PUT("/alert/update/:id", database.AlertUpdate)
	// router.DELETE("/alert/delete/:id", database.AlertDelete)

	// router.GET("/records", database.RecordGet)
	// router.GET("/record/:id", database.RecordGetById)
	// router.POST("/record/create", database.RecordCreate)
	// router.PUT("/record/update/:id", database.RecordUpdate)
	// router.DELETE("/record/delete/:id", database.RecordDelete)

}
