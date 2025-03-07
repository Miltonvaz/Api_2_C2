package routes

import (
	"API_notification/src/notification/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(router *gin.Engine, controller *controllers.NotificationController) {
	router.POST("/api/notifications", controller.CreateNotification)

}
