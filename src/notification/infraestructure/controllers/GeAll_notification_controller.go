package controllers

import (
	"API_notification/src/notification/application/Use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllNotificationsController struct {
	getAllNotificationsUseCase *Use_case.GetAllNotificationsUseCase
}

func NewGetAllNotificationsController(getAllNotificationsUseCase *Use_case.GetAllNotificationsUseCase) *GetAllNotificationsController {
	return &GetAllNotificationsController{getAllNotificationsUseCase: getAllNotificationsUseCase}
}

func (c *GetAllNotificationsController) Execute(ctx *gin.Context) {
	notifications, err := c.getAllNotificationsUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get notifications"})
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}
