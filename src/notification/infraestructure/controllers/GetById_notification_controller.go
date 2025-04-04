package controllers

import (
	"API_notification/src/notification/application/Use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetNotificationByIDController struct {
	getNotificationByIDUseCase *Use_case.GetNotificationByIDUseCase
}

func NewGetNotificationByIDController(getNotificationByIDUseCase *Use_case.GetNotificationByIDUseCase) *GetNotificationByIDController {
	return &GetNotificationByIDController{getNotificationByIDUseCase: getNotificationByIDUseCase}
}

func (c *GetNotificationByIDController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")

	notification, err := c.getNotificationByIDUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	ctx.JSON(http.StatusOK, notification)
}
