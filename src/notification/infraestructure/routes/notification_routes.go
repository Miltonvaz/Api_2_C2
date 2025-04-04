package routes

import (
	"API_notification/src/notification/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(router *gin.Engine, notificationController *controllers.NotificationController, deleteNotificationController *controllers.DeleteNotificationController, getAllNotificationsController *controllers.GetAllNotificationsController, getNotificationByIDController *controllers.GetNotificationByIDController, updateNotificationController *controllers.UpdateNotificationController) {
	router.POST("/api/notifications", notificationController.Execute)
	router.DELETE("/api/notifications/:id", deleteNotificationController.Execute)
	router.GET("/api/notifications", getAllNotificationsController.Execute)
	router.GET("/api/notifications/:id", getNotificationByIDController.Execute)
	router.PUT("/api/notifications/:id", updateNotificationController.Execute)
}
