package controllers

import (
	"API_notification/src/notification/application/Use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteNotificationController struct {
	deleteNotificationUseCase *Use_case.DeleteNotificationUseCase
}

func NewDeleteNotificationController(deleteNotificationUseCase *Use_case.DeleteNotificationUseCase) *DeleteNotificationController {
	return &DeleteNotificationController{deleteNotificationUseCase: deleteNotificationUseCase}
}

func (c *DeleteNotificationController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.deleteNotificationUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}
