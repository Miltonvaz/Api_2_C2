package a_rabbit

import (
	"API_notification/src/notification/application/Use_case/repositories"
	"API_notification/src/notification/domain/entities"
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var _ repositories.NotificationPort = (*RabbitMQAdapter)(nil)

// Crear una única instancia de RabbitMQAdapter
var rabbitInstance *RabbitMQAdapter

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	if rabbitInstance != nil {
		return rabbitInstance, nil
	}

	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Printf("Error conectando a RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error abriendo canal: %v", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"notificaciones",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error declarando la cola: %v", err)
		return nil, err
	}

	rabbitInstance = &RabbitMQAdapter{conn: conn, ch: ch}
	return rabbitInstance, nil
}

func (r *RabbitMQAdapter) PublishNotification(notification entities.Notification) error {
	body, err := json.Marshal(notification)
	if err != nil {
		log.Printf("Error convirtiendo evento a JSON: %v", err)
		return err
	}

	// Usar contexto con timeout para evitar bloqueos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.ch.PublishWithContext(
		ctx,
		"",
		"notificaciones",
		true,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Error enviando mensaje a RabbitMQ: %v", err)
		return err
	}

	log.Println("Evento publicado:", notification.UserID)
	return nil
}

// Cerrar conexión de RabbitMQ correctamente
func (r *RabbitMQAdapter) Close() {
	if r.ch != nil {
		if err := r.ch.Close(); err != nil {
			log.Printf("Error cerrando canal RabbitMQ: %v", err)
		}
	}
	if r.conn != nil {
		if err := r.conn.Close(); err != nil {
			log.Printf("Error cerrando conexión RabbitMQ: %v", err)
		}
	}
}
