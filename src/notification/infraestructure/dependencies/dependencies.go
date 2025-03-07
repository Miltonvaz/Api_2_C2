package dependencies

import (
	usecases "API_notification/src/notification/application/Use_case"
	"API_notification/src/notification/application/Use_case/repositories"
	"API_notification/src/notification/infraestructure/adapters/a_rabbit"
	"API_notification/src/notification/infraestructure/controllers"
	"log"
)

func InitializeDependencies() (*usecases.CreateNotificationUseCase, *controllers.NotificationController, error) {

	rabbitAdapter, err := a_rabbit.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error al crear adaptador RabbitMQ: %v", err)
	}

	serviceNotification := repositories.NewServiceNotification(rabbitAdapter)

	createNotificationUseCase := usecases.NewCreateNotificationUseCase(serviceNotification)

	notificationController := controllers.NewNotificationController(createNotificationUseCase, serviceNotification)

	return createNotificationUseCase, notificationController, nil
}
