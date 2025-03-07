package repositories

import (
	"API_notification/src/notification/domain/entities"
	"log"
)

type ServiceNotification struct {
	rabbitPort NotificationPort
}

func NewServiceNotification(rabbitPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{
		rabbitPort: rabbitPort,
	}
}

func (sn *ServiceNotification) NotifyAppointmentCreated(notification entities.Notification) error {
	log.Println("Notificación...")

	err := sn.rabbitPort.PublishNotification(notification)
	if err != nil {
		log.Printf("Error al publicar el evento en RabbitMQ: %v", err)
		return err
	}

	return nil
}
