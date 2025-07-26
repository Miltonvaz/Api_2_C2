package main

import (
	"API_notification/src/notification/infraestructure/dependencies"
	"API_notification/src/notification/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
	}))

	notificationController, deleteNotificationController, getAllNotificationsController, getNotificationByIDController, updateNotificationController, err := dependencies.InitializeDependencies()
	if err != nil {
		log.Fatalf("Error inicializando dependencias: %v", err)
	}

	routes.SetupNotificationRoutes(r, notificationController, deleteNotificationController, getAllNotificationsController, getNotificationByIDController, updateNotificationController)

	go func() {
		log.Println("Servidor corriendo en http://localhost:8083")
		if err := r.Run(":8083"); err != nil {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	<-signalChannel

	log.Println("Deteniendo servidor...")
	time.Sleep(2 * time.Second)
	log.Println("Servidor detenido.")
}
