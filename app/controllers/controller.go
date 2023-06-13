package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"virtual-plant-nursery-go/app/models"
)

// Obtener el estado de todas las plantas

func GetAllPlants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ListPlants)
}

// Obtener el estado de una planta específica
func GetPlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for _, plant := range models.ListPlants {
		if plant.ID == id {
			json.NewEncoder(w).Encode(plant)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// Crear una nueva planta
func CreatePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var plant *models.Plant
	err := json.NewDecoder(r.Body).Decode(&plant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.ListPlants = append(models.ListPlants, plant)
	json.NewEncoder(w).Encode(plant)
}

// Actualizar el estado de una planta existente
func UpdatePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for index, plant := range models.ListPlants {
		if plant.ID == id {
			models.ListPlants = append(models.ListPlants[:index], models.ListPlants[index+1:]...)
			var updatedPlant *models.Plant
			err := json.NewDecoder(r.Body).Decode(&updatedPlant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			models.ListPlants = append(models.ListPlants, updatedPlant)
			json.NewEncoder(w).Encode(updatedPlant)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// Eliminar una planta
func DeletePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for index, plant := range models.ListPlants {
		if plant.ID == id {
			models.ListPlants = append(models.ListPlants[:index], models.ListPlants[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.ListPlants)
}

// Agregar agua a todas las plantas
func AddWaterToPlants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	//Aqui iria la logica para el agregado de agua en cada planta(calculado los porcentajes)
	for i := range models.ListPlants {
		models.ListPlants[i].WaterSystem += amount
	}

	json.NewEncoder(w).Encode(models.ListPlants)
}

// Agregar nutrientes a todas las plantas
func AddNutrientsToPlants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	//Aqui iria la logica para el agregado de nutrientas en cada planta(calculado los porcentajes)
	for i := range models.ListPlants {
		models.ListPlants[i].NutrientSystem += amount
	}

	json.NewEncoder(w).Encode(models.ListPlants)
}
