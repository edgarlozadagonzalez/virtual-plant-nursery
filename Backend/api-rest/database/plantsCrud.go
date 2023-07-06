package database

import (
	//"net/http"
	//"strconv"
	"github.com/gin-gonic/gin"
	"virtual-plant-nursery/api-rest/models"
	
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PlantRequestBody struct {

	Name                  string  `json:"name"`
	SurvivalDegree        float64 `json:"survival_degree"`
	WaterRequired         float64 `json:"water_required"`
	WaterSystem           float64 `json:"water_system"`
	HydrationLevel        float64 `json:"hydration_level"`
	NutrientsRequired     float64 `json:"nutrients_required"`
	NutrientSystem        float64 `json:"nutrient_system"`
	NutritionLevel        float64 `json:"nutrition_level"`
	Alive                 bool    `json:"alive"`
}

func PlantCreate(c *gin.Context) {
	body := PlantRequestBody{}
	c.BindJSON(&body)

	plant := &models.Plant{

		Name:                  body.Name,
		SurvivalDegree:        body.SurvivalDegree,
		WaterRequired:         body.WaterRequired,
		WaterSystem:           body.WaterSystem,
		HydrationLevel:        body.HydrationLevel,
		NutrientsRequired:     body.NutrientsRequired,
		NutrientSystem:        body.NutrientSystem,
		NutritionLevel:        body.NutritionLevel,
		Alive:                 body.Alive,
	}

	result := db.Create(&plant)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &plant)
}

func PlantGet(c *gin.Context) {
	var plants []models.Plant
	db.Find(&plants)
	c.JSON(200, &plants)
	return
}

func PlantGetById(c *gin.Context) {
	id := c.Param("id")
	var plant models.Plant
	db.First(&plant, id)
	c.JSON(200, &plant)
	return
}

func PlantUpdate(c *gin.Context) {
	id := c.Param("id")
	var plant models.Plant
	db.First(&plant, id)

	body := PlantRequestBody{}
	c.BindJSON(&body)

	data := &models.Plant{

		Name:                  body.Name,
		SurvivalDegree:        body.SurvivalDegree,
		WaterRequired:         body.WaterRequired,
		WaterSystem:           body.WaterSystem,
		HydrationLevel:        body.HydrationLevel,
		NutrientsRequired:     body.NutrientsRequired,
		NutrientSystem:        body.NutrientSystem,
		NutritionLevel:        body.NutritionLevel,
		Alive:                 body.Alive,
	}

	result := db.Model(&plant).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &plant)
}

func PlantDelete(c *gin.Context) {
	id := c.Param("id")
	var plant models.Plant
	db.Delete(&plant, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
/*
// Agregar agua a todas las plantas
func AddWaterToPlants(c *gin.Context) {
	amountStr := c.Param("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "El valor debe ser numerico."})
		return
	}
	for i := range models.ListPlants {
		models.ListPlants[i].WaterSystem += amount
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Agua agregada correctamente a todas las plantas."})
}

// Agregar nutrientes a todas las plantas
func AddNutrientsToPlants(c *gin.Context) {
	amountStr := c.Param("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "El valor debe ser numerico."})
		return
	}
	for i := range models.ListPlants {
		models.ListPlants[i].NutrientSystem += amount
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Nutrientes agregados correctamente a todas las plantas."})
}


*/