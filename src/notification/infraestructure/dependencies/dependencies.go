package dependencies

import (
	"API_notification/src/core"
	usecases "API_notification/src/notification/application/Use_case"
	"API_notification/src/notification/application/Use_case/repositories"
	"API_notification/src/notification/infraestructure/adapters"
	"API_notification/src/notification/infraestructure/adapters/a_rabbit"
	"API_notification/src/notification/infraestructure/controllers"
	"log"
)

func InitializeDependencies() (
	*controllers.NotificationController,
	*controllers.DeleteNotificationController,
	*controllers.GetAllNotificationsController,
	*controllers.GetNotificationByIDController,
	*controllers.UpdateNotificationController,
	error,
) {
	pool := core.GetDBPool()
	ps := adapters.NewMySQL(pool.DB)

	// Crear una sola instancia de RabbitMQAdapter
	rabbitAdapter, err := a_rabbit.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error al crear adaptador RabbitMQ: %v", err)
	}

	serviceNotification := repositories.NewServiceNotification(rabbitAdapter)

	// Crear casos de uso
	createNotificationUseCase := usecases.NewCreateNotificationUseCase(serviceNotification, ps)
	deleteNotificationUseCase := usecases.NewDeleteNotificationUseCase(ps)
	getAllNotificationsUseCase := usecases.NewGetAllNotificationsUseCase(ps)
	getNotificationByIDUseCase := usecases.NewGetNotificationByIDUseCase(ps)
	updateNotificationUseCase := usecases.NewUpdateNotificationUseCase(ps)

	// Crear controladores
	notificationController := controllers.NewNotificationController(createNotificationUseCase, serviceNotification)
	deleteNotificationController := controllers.NewDeleteNotificationController(deleteNotificationUseCase)
	getAllNotificationsController := controllers.NewGetAllNotificationsController(getAllNotificationsUseCase)
	getNotificationByIDController := controllers.NewGetNotificationByIDController(getNotificationByIDUseCase)
	updateNotificationController := controllers.NewUpdateNotificationController(updateNotificationUseCase)

	return notificationController, deleteNotificationController, getAllNotificationsController, getNotificationByIDController, updateNotificationController, nil
}
