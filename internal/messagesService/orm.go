package messagesService

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TextData string `json:"text"` // Наш сервер будет ожидать json c полем text
}
