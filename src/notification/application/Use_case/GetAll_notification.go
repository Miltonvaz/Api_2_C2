package Use_case

import (
	"API_notification/src/notification/domain"
	"API_notification/src/notification/domain/entities"
)

type GetAllNotificationsUseCase struct {
	notificationPort domain.NotificationPort
}

func NewGetAllNotificationsUseCase(notificationPort domain.NotificationPort) *GetAllNotificationsUseCase {
	return &GetAllNotificationsUseCase{notificationPort: notificationPort}
}

func (uc *GetAllNotificationsUseCase) Execute() ([]entities.Notification, error) {
	return uc.notificationPort.GetAll()
}
