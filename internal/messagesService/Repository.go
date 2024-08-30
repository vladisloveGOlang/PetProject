package messagesService

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"first/internal/web/messages"
)

type MessageRepository interface {

	// CreateMessage - Передаем в функцию message типа Message из orm.go
	// возвращаем созданный Message и ошибку
	CreateMessage(message Message) (Message, error)

	// GetAllMessages - Возвращаем массив из всех писем в БД и ошибку
	GetAllMessages() ([]Message, error)

	// UpdateMessageByID - Передаем id и Message, возвращаем обновленный Message
	// и ошибку
	UpdateMessageByID(id uint, message messages.Message) (messages.Message, error)

	// DeleteMessageByID - Передаем id для удаления, возвращаем только ошибку
	DeleteMessageByID(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

// (r *messageRepository) привязывает данную функцию к нашему репозиторию
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id uint, message messages.Message) (messages.Message, error) {
	result := r.db.Model(&message).Omit("id").Updates(map[string]interface{}{
		"id":   id,
		"text": message.Message,
	})
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id uint) error {
	var mess Message

	result := r.db.Where("id = ? AND deleted_at is NULL", id).Delete(&mess)
	if result.RowsAffected == 0 {
		return errors.New("ошибка при удалении, 0 строк")
	}
	if result.Error != nil {
		fmt.Print("Ошибка при удалении")
		return result.Error
	}
	return result.Error
}
