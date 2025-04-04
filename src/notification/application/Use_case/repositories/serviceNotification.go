package repositories

import (
	"API_notification/src/notification/domain/entities"
	"log"
)

type ServiceNotification struct {
	rabbitPort NotificationPort
}

func NewServiceNotification(rabbitPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{rabbitPort: rabbitPort}
}

func (sn *ServiceNotification) NotifyAppointmentCreated(notification entities.Notification) error {
	log.Printf("Notificación recibida: %+v", notification)

	go func() {
		err := sn.rabbitPort.PublishNotification(notification)
		if err != nil {
			log.Printf("Error al publicar en RabbitMQ: %v", err)
		} else {
			log.Println("Notificación enviada a RabbitMQ exitosamente")
		}
	}()

	return nil
}
