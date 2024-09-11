package messagesService

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TextName        string `json:"nameD"`
	TextDescription string `json:"descriptionD"` // Наш сервер будет ожидать json c полем text
}
