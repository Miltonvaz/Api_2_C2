package Use_case

import (
	"API_notification/src/notification/application/Use_case/repositories"
	"API_notification/src/notification/domain/entities"
)

type CreateNotificationUseCase struct {
	notificationService *repositories.ServiceNotification
}

func NewCreateNotificationUseCase(notificationService *repositories.ServiceNotification) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{
		notificationService: notificationService,
	}
}

func (u *CreateNotificationUseCase) Execute(notification entities.Notification) error {
	err := u.notificationService.NotifyAppointmentCreated(notification)
	return err
}
