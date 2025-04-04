package controllers

import (
	"API_notification/src/notification/application/Use_case"
	"API_notification/src/notification/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateNotificationController struct {
	updateNotificationUseCase *Use_case.UpdateNotificationUseCase
}

func NewUpdateNotificationController(updateNotificationUseCase *Use_case.UpdateNotificationUseCase) *UpdateNotificationController {
	return &UpdateNotificationController{updateNotificationUseCase: updateNotificationUseCase}
}

func (c *UpdateNotificationController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")

	var notification entities.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updatedNotification, err := c.updateNotificationUseCase.Execute(id, notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notification"})
		return
	}

	ctx.JSON(http.StatusOK, updatedNotification)
}
