package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"virtual-plant-nursery/api-rest/models"
	"virtual-plant-nursery/api-rest/services"

	"github.com/gin-gonic/gin"
)

// Obtener el estado de todas las plantas
func GetAllPlants(c *gin.Context) {
	services.LoadListPlants()
	c.IndentedJSON(http.StatusOK, services.ListPlants)
}

// Obtener el estado de una planta específica
func GetPlant(c *gin.Context) {
	id := ParseUint(c.Param("id"))

	for _, plant := range services.ListPlants {
		if plant.ID == id {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")
			c.Header("Access-Control-Allow-Origin", "*")

			dataChannel := make(chan models.Plant)
			go services.GenerateConstantData(plant, dataChannel)

			encoder := json.NewEncoder(c.Writer)
			for constantData := range dataChannel {
				// Serializar la planta en formato JSON
				err := encoder.Encode(constantData)
				if err != nil {
					c.Error(err)
					return
				}
				c.Writer.WriteString("\n")
				c.Writer.Flush()
			}
			close(dataChannel)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Planta no encontrada."})
}

// Crear una nueva planta
func CreatePlant(c *gin.Context) {
	var plant *models.Plant
	if err := c.BindJSON(&plant); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Verifique la estructura de los datos."})
		return
	}
	if services.CreatePlant(*plant) {
		services.ListPlants = append(services.ListPlants, plant)
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Planta creada."})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar en base de datos."})
}

func ParseUint(str string) uint {
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint(value)
}

// Actualizar una planta existente
func UpdatePlant(c *gin.Context) {
	var updateplant *models.Plant
	if err := c.BindJSON(&updateplant); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Verifique la estructura de los datos."})
		return
	}

	if services.UpdatePlant(*updateplant) {
		services.LoadListPlants()
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Planta actualizada."})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al actualizar la planta."})
}

// Eliminar una planta
func DeletePlant(c *gin.Context) {
	id := ParseUint(c.Param("id"))

	if services.DeletePlant(id) {
		services.LoadListPlants()
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Planta eliminada."})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error al eliminar la planta."})
}

// Agregar agua a una planta
func AddWaterToPlant(c *gin.Context) {
	plantID := ParseUint(c.Param("id"))
	amountStr := c.Param("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "El valor debe ser numérico."})
		return
	}
	if services.AddWaterToPlant(plantID, amount) {
		services.LoadListPlants()
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Agua agregada correctamente a la planta."})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Lamentablemente, se produjo un inconveniente durante el proceso de actualización de la planta."})
}

// Agregar nutrientes a una planta
func AddNutrientsToPlant(c *gin.Context) {
	plantID := ParseUint(c.Param("id"))
	amountStr := c.Param("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "El valor debe ser numérico."})
		return
	}

	if services.AddNutrientsToPlant(plantID, amount) {
		services.LoadListPlants()
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Nutrientes agregados correctamente a la planta."})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Lamentablemente, se produjo un inconveniente durante el proceso de actualización de la planta."})
}

// Agregar agua a todas las plantas
func AddWaterToPlants(c *gin.Context) {
	amountStr := c.Param("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "El valor debe ser numerico."})
		return
	}
	services.AddWaterToPlants(amount)
	services.LoadListPlants()
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
	services.AddNutrientsToPlants(amount)
	services.LoadListPlants()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Nutrientes agregados correctamente a todas las plantas."})
}
