package controllers

import (
	"net/http"
	"strconv"
	"virtual-plant-nursery/api-rest/models"

	"github.com/gin-gonic/gin"
)

// Obtener el estado de todas las plantas
func GetAllPlants(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.ListPlants)
}

// Obtener el estado de una planta específica
func GetPlant(c *gin.Context) {
	id := c.Param("id")

	for _, plant := range models.ListPlants {
		if plant.ID == id {
			c.IndentedJSON(http.StatusOK, plant)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Planta no encontrada."})
}

// Crear una nueva planta
func CreatePlant(c *gin.Context) {
	var plant *models.Plant
	if err := c.BindJSON(&plant); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Verifique la estructura de los datos."})
		return
	}
	models.ListPlants = append(models.ListPlants, plant)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Planta creada."})
}

// Actualizar una planta existente
func UpdatePlant(c *gin.Context) {
	id := c.Param("id")
	var updateplant *models.Plant
	if err := c.BindJSON(&updateplant); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Verifique la estructura de los datos."})
		return
	}
	for index, plant := range models.ListPlants {
		if plant.ID == id {
			models.ListPlants = append(models.ListPlants[:index], models.ListPlants[index+1:]...)
			models.ListPlants = append(models.ListPlants, updateplant)
			c.IndentedJSON(http.StatusCreated, gin.H{"message": "Planta actualizada."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No existe la planta para actualizar."})
}

// Eliminar una planta
func DeletePlant(c *gin.Context) {
	id := c.Param("id")

	for index, plant := range models.ListPlants {
		if plant.ID == id {
			models.ListPlants = append(models.ListPlants[:index], models.ListPlants[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Planta eliminada."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Planta no encontrada."})
}

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
