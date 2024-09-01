package handlers

import (
	//"fmt"
	"context"
	//"fmt"
	///"io"
	//"net/http"

	data "first/internal/database"
	"first/internal/web/messages"

	"encoding/json"
	ms "first/internal/messagesService"
	"log"
	//"strings"
	//"github.com/gorilla/mux"
)

type Handler struct {
	Service *ms.MessageService
}

// DeleteMessages implements messages.StrictServerInterface.
func (h *Handler) DeleteMessages(ctx context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	messageRequest := request.Body

	err := h.Service.DeleteMessage(*messageRequest.Id)
	if err != nil {
		return nil, err
	}

	response := messages.DeleteMessages204JSONResponse{
		Id:      messageRequest.Id,
		Message: messageRequest.Message,
	}
	return response, nil

}

// PatchMessages implements messages.StrictServerInterface.
func (h *Handler) PatchMessages(ctx context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	messageRequest := request.Body

	messageToPatch := messages.Message{
		Id:      messageRequest.Id,
		Message: messageRequest.Message,
	}
	_, err := h.Service.PatchMessage(*messageRequest.Id, messageToPatch)
	if err != nil {
		return nil, err
	}

	var response messages.PatchMessages200JSONResponse

	data, err := json.Marshal(messageRequest)
	if err != nil {
		log.Fatal("json не раскодируется")
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal("json не закодирывается")
	}

	return response, nil
}

// PostMessages implements messages.StrictServerInterface.
func (h *Handler) PostMessages(ctx context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := ms.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	// Просто возвращаем респонс!
	return response, nil
}

// GetMessages implements messages.StrictServerInterface.
func (h *Handler) GetMessages(ctx context.Context, request messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {

	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}
	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}
	return response, nil
}

func NewHandler(service *ms.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

var DB = data.DB

//var msg ms.Message
