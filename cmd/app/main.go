package main

import (
	"net/http"

	"first/internal/database"
	data "first/internal/database"
	hand "first/internal/handlers"
	ms "first/internal/messagesService"

	"github.com/gorilla/mux"
)

var DB = data.DB

func main() {

	//инициализация переменной DB как базы данных
	DB = data.InitDB()
	//создание струтуры этой базы
	DB.AutoMigrate(ms.Message{})

	//создание струтуры и методов работы с этой струтурой репозитория
	repo := ms.NewMessageRepository(database.DB)

	//Создание методов работы с репозиторием
	service := ms.NewService(repo)

	//создание методов работы с сервисом
	handler := hand.NewHandler(service)

	//создаем роутер
	router := mux.NewRouter()

	//при заданном адресе. запускаем метод PostMessageHandler из струтуры хендлер, он вызывает метод CreateMessage из
	router.HandleFunc("/postjson", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/hello", handler.GetMessageHandler).Methods("GET")
	router.HandleFunc("/patch", handler.UpdateMessageByID).Methods("PATCH")
	router.HandleFunc("/delete", handler.DeleteMessageByIDHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
