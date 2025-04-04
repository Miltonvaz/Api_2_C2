package Use_case

import (
	"API_notification/src/notification/domain"
	"API_notification/src/notification/domain/entities"
)

type UpdateNotificationUseCase struct {
	notificationPort domain.NotificationPort
}

func NewUpdateNotificationUseCase(notificationPort domain.NotificationPort) *UpdateNotificationUseCase {
	return &UpdateNotificationUseCase{notificationPort: notificationPort}
}

func (uc *UpdateNotificationUseCase) Execute(id string, notification entities.Notification) (entities.Notification, error) {
	return uc.notificationPort.Update(id, notification)
}
