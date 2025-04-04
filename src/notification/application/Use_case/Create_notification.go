package Use_case

import (
	"API_notification/src/notification/application/Use_case/repositories"
	"API_notification/src/notification/domain"
	"API_notification/src/notification/domain/entities"
	"log"
)

type CreateNotificationUseCase struct {
	notificationService *repositories.ServiceNotification
	notificationRepo    domain.NotificationPort
}

func NewCreateNotificationUseCase(notificationService *repositories.ServiceNotification, notificationRepo domain.NotificationPort) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{
		notificationService: notificationService,
		notificationRepo:    notificationRepo,
	}
}

func (u *CreateNotificationUseCase) Execute(notification entities.Notification) error {
	savedNotification, err := u.notificationRepo.Save(notification)
	if err != nil {
		return err
	}

	go func() {
		if err := u.notificationService.NotifyAppointmentCreated(savedNotification); err != nil {
			log.Printf("Error publicando notificación en RabbitMQ: %v", err)
		}
	}()

	return nil
}
