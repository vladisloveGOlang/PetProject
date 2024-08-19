package handlers

import (
	"fmt"
	"io"
	"net/http"

	data "first/internal/database"
	ms "first/internal/messagesService"
	"log"

	"encoding/json"

	"strings"
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

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

var DB = data.DB

var msg ms.Message

func GetHandler(w http.ResponseWriter, r *http.Request) {

	//создаем слайс куда будем сохранять содержимое нашеё БД
	var texts []string

	//устанавливем с что хотим работать с моделью мессадж в DB,
	//Планком из столбца текст вытаскиваем значения в наш слайс
	result := DB.Model(&ms.Message{}).Pluck("Text", &texts)
	if result.Error != nil {
		fmt.Println("Error reading from database:", result.Error)
		return
	}

	//склеиваем слайс через новую строку
	text := strings.Join(texts, "\n")

	fmt.Fprintln(w, text)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	JsonFromRtoMSG(w, r)

	// Добавляем запись в базу данных
	result := DB.Create(&msg)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании записи: %v", result.Error)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Сообщение успешно добавлено"))
}

var lastValue ms.Message

func PatchHandler(w http.ResponseWriter, r *http.Request) {

	JsonFromRtoMSG(w, r)

	lastValue.ID = msg.ID

	result := DB.First(&lastValue, msg.ID)

	messLast := fmt.Sprintf("Текущие значения:\nID : %v, Text :%s", lastValue.ID, lastValue.Text)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(messLast))
	if result.Error != nil {
		fmt.Println("Error reading from database:", result.Error)
		return
	}

	result = DB.Updates(&msg)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Изменнений не произошло"))
		return
	}
	if result.Error != nil {
		log.Fatalf("Ошибка при обновлении записи: %v", result.Error)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	// Возвращаем успешный ответ
	status := fmt.Sprintf("\nОбновление совершено, Текущие значения:\nID : %v, Text :%s", msg.ID, msg.Text)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(status))

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

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	JsonFromRtoMSG(w, r)
	result := DB.Delete(&msg)
	if result.RowsAffected == 0 {
		errorText := fmt.Sprintf("Изменнений не произошло так как строка c id %d не существует", msg.ID)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(errorText))
		return
	}
	if result.Error != nil {
		log.Fatalf("Ошибка при удалении записи: %v", result.Error)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)

		return
	}
	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Обновление совершено"))

}
