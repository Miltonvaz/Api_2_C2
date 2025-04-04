package entities

type Notification struct {
	ID      int64
	UserID  string `json:"user_id"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
