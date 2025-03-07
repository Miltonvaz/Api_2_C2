package controllers

import (
	usecases "API_notification/src/notification/application/Use_case"
	"API_notification/src/notification/application/Use_case/repositories"
	domain "API_notification/src/notification/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationController struct {
	createUseCase *usecases.CreateNotificationUseCase
	notifyService *repositories.ServiceNotification
}

func NewNotificationController(createUseCase *usecases.CreateNotificationUseCase, notifyService *repositories.ServiceNotification) *NotificationController {
	return &NotificationController{
		createUseCase: createUseCase,
		notifyService: notifyService,
	}
}

func (c *NotificationController) CreateNotification(ctx *gin.Context) {
	var notification domain.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.createUseCase.Execute(notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process notification"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Notification created"})
}
