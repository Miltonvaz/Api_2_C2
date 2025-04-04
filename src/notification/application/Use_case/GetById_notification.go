package Use_case

import (
	"API_notification/src/notification/domain"
	"API_notification/src/notification/domain/entities"
)

type GetNotificationByIDUseCase struct {
	notificationPort domain.NotificationPort
}

func NewGetNotificationByIDUseCase(notificationPort domain.NotificationPort) *GetNotificationByIDUseCase {
	return &GetNotificationByIDUseCase{notificationPort: notificationPort}
}

func (uc *GetNotificationByIDUseCase) Execute(id string) (entities.Notification, error) {
	return uc.notificationPort.GetByID(id)
}
