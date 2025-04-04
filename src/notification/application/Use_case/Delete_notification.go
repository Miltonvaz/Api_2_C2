package Use_case

import "API_notification/src/notification/domain"

type DeleteNotificationUseCase struct {
	notificationPort domain.NotificationPort
}

func NewDeleteNotificationUseCase(notificationPort domain.NotificationPort) *DeleteNotificationUseCase {
	return &DeleteNotificationUseCase{notificationPort: notificationPort}
}

func (uc *DeleteNotificationUseCase) Execute(id string) error {
	return uc.notificationPort.Delete(id)
}
