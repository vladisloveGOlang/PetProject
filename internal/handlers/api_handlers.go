package handlers

import (
	//"fmt"
	"fmt"
	"io"
	"net/http"

	data "first/internal/database"

	ms "first/internal/messagesService"
	//"log"

	"encoding/json"
	//"strings"
	//"github.com/gorilla/mux"
)

type Handler struct {
	Service *ms.MessageService
}

func NewHandler(service *ms.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

var DB = data.DB

var msg ms.Message

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message ms.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

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
