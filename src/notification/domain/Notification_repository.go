package domain

import "API_notification/src/notification/domain/entities"

type NotificationPort interface {
	PublishNotification(notification entities.Notification) error
	Save(notification entities.Notification) (entities.Notification, error)
	GetByID(id string) (entities.Notification, error)
	Delete(id string) error
	GetAll() ([]entities.Notification, error)
	Update(id string, notification entities.Notification) (entities.Notification, error)
}
