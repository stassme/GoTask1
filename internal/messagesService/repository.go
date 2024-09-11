package messagesService

import (
	"gorm.io/gorm"
)

type MessageRepository interface {
	CreateMessage(m Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id int, m Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(m Message) (Message, error) {
	result := r.db.Create(&m)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return m, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id int, m Message) (Message, error) {
	err := r.db.Model(&Message{}).Where("id = ?", id).Updates(m).Error
	return m, err
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	return r.db.Delete(&Message{}, id).Error
}
