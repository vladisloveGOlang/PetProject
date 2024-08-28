package handlers

import (
	//"fmt"
	"context"
	"fmt"
	"io"
	"net/http"

	data "first/internal/database"
	"first/internal/web/messages"

	ms "first/internal/messagesService"
	//"log"

	"encoding/json"
	//"strings"
	//"github.com/gorilla/mux"
)

type Handler struct {
	Service *ms.MessageService
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

func NewHandler(service *ms.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

var DB = data.DB

var msg ms.Message

func (h *Handler) DeleteMessageByIDHandler(w http.ResponseWriter, r *http.Request) {
	JsonFromRtoMSG(w, r)
	err := h.Service.DeleteMessage(msg.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) ////////////////////////////////////////////////////////////////////// Не понимаю как писать обработку ошибок
		return
	}
	ans := fmt.Sprintf("строка с id: %v удалена", msg.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ans)
}

func (h *Handler) UpdateMessageByID(w http.ResponseWriter, r *http.Request) {
	JsonFromRtoMSG(w, r)
	Message, err := h.Service.PatchMessage(msg.ID, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) ////////////////////////////////////////////////////////////////////// Не понимаю как писать обработку ошибок
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message)
}

// Распаршивает json из запроса в msg
func JsonFromRtoMSG(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(bodyBytes, &msg)
	if err != nil {
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)

		return
	}

}
