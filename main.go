package main

import (
	"API_notification/src/notification/infraestructure/dependencies"
	"API_notification/src/notification/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	_, handler, err := dependencies.InitializeDependencies()
	if err != nil {
		log.Fatalf("Error inicializando dependencias: %v", err)
		return
	}

	routes.SetupNotificationRoutes(r, handler)

	log.Println("Servidor corriendo en http://localhost:8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
