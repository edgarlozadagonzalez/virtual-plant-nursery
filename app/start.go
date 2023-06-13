package app

import (
	"log"
	"net/http"
	"virtual-plant-nursery-go/app/routes"
)

func Start() {
	routes.Routes()
	// Iniciar el servidor en el puerto 8080
	log.Println("Iniciando Servidor en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
