package adapters

import (
	"API_notification/src/notification/domain/entities"
	"database/sql"
	"errors"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

// Save method updated to return the saved notification and error
func (m *MySQL) Save(notification entities.Notification) (entities.Notification, error) {
	query := `INSERT INTO notifications (user_id, message, status) 
              VALUES (?, ?, ?)`
	result, err := m.conn.Exec(query, notification.UserID, notification.Message, notification.Status)
	if err != nil {
		return entities.Notification{}, fmt.Errorf("failed to save notification: %v", err)
	}

	// Get the ID of the newly inserted row
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return entities.Notification{}, fmt.Errorf("failed to get last insert id: %v", err)
	}

	// Assign the generated ID to the notification and return it
	notification.ID = lastInsertID
	return notification, nil
}

// PublishNotification is now updated to handle the return values correctly
// PublishNotification is now updated to return only error
func (m *MySQL) PublishNotification(notification entities.Notification) error {
	// Save the notification and return any error
	_, err := m.Save(notification)
	return err
}

func (m *MySQL) GetByID(id string) (entities.Notification, error) {
	var notification entities.Notification
	query := `SELECT user_id, message, status FROM notifications WHERE id = ? LIMIT 1`
	err := m.conn.QueryRow(query, id).Scan(
		&notification.UserID, &notification.Message, &notification.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Notification{}, errors.New("notification not found")
		}
		return entities.Notification{}, err
	}
	return notification, nil
}

func (m *MySQL) GetAll() ([]entities.Notification, error) {
	query := "SELECT user_id, message, status FROM notifications"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve notifications: %v", err)
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		err := rows.Scan(&notification.UserID, &notification.Message, &notification.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return notifications, nil
}

func (m *MySQL) Update(id string, notification entities.Notification) (entities.Notification, error) {
	query := "UPDATE notifications SET user_id = ?, message = ?, status = ? WHERE id = ?"
	_, err := m.conn.Exec(query, notification.UserID, notification.Message, notification.Status, id)
	if err != nil {
		return entities.Notification{}, fmt.Errorf("failed to update notification: %v", err)
	}

	// Return updated notification (usually you'll want to retrieve it after updating)
	updatedNotification, err := m.GetByID(id)
	if err != nil {
		return entities.Notification{}, fmt.Errorf("failed to retrieve updated notification: %v", err)
	}

	return updatedNotification, nil
}

func (m *MySQL) Delete(id string) error {
	query := "DELETE FROM notifications WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %v", err)
	}
	return nil
}
