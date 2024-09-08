package handlers

import (
	//"fmt"
	"context"
	//"fmt"
	///"io"
	//"net/http"

	data "first/internal/database"

	"first/internal/web/messages"
	"first/internal/web/users"

	"encoding/json"
	ms "first/internal/messagesService"
	"log"

	//"strings"
	//"github.com/gorilla/mux"
	us "first/internal/userService"
)

type Handler struct {
	Service *ms.MessageService
}

func NewHandler(service *ms.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

type UHandler struct {
	uService *us.UserService
}

// DeleteUsers implements users.StrictServerInterface.
func (h *UHandler) DeleteUsers(ctx context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	ans, err := h.uService.DeleteUserById(*request.Body)

	if err != nil {
		log.Fatal(err)
	}
	return ans, nil
}

// PatchUsers implements users.StrictServerInterface.
func (h *UHandler) PatchUsers(ctx context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	ans, err := h.uService.PatchUser(*request.Body)

	if err != nil {
		log.Fatal(err)
	}
	return ans, nil
}

// PostUsers implements users.StrictServerInterface.
func (h *UHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

	err := h.uService.CreateNewUser(*request.Body)
	if err != nil {
		log.Fatal(err)
	}

	status := true

	ans := users.Ans{
		Id:      request.Body.Id,
		Changed: &status,
	}

	ansf := users.PostUsers201JSONResponse{
		Id:      ans.Id,
		Changed: ans.Changed,
	}

	return ansf, nil

}

func NewUHandler(service *us.UserService) *UHandler {
	return &UHandler{
		uService: service,
	}
}

func (h *UHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {

	ans, err := h.uService.GetAllMessages()
	if err != nil {
		log.Fatal(err)
	}

	var ansf users.GetUsers200JSONResponse

	ansf = ans

	return ansf, err

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

var DB = data.DB
