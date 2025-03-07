package domain

import "API_notification/src/notification/domain/entities"

type NotificationPort interface {
	PublishNotification(notification entities.Notification) error
	Save(notification entities.Notification) (entities.Notification, error)
}
