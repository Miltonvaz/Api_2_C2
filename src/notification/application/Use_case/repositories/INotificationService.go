package repositories

import "API_notification/src/notification/domain/entities"

type NotificationPort interface {
	PublishNotification(notification entities.Notification) error
}
