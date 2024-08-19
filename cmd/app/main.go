package main

import (
	"net/http"

	data "first/internal/database"
	hand "first/internal/handlers"
	ms "first/internal/messagesService"

	"github.com/gorilla/mux"
)

var DB = data.DB

func main() {

	//инициализация переменной DB как базы данных
	DB = data.InitDB()
	DB.AutoMigrate(ms.Message{})
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/postjson", hand.PostHandler).Methods("POST")
	router.HandleFunc("/api/hello", hand.GetHandler).Methods("GET")
	router.HandleFunc("/patch", hand.PatchHandler).Methods("PATCH")
	router.HandleFunc("/delete", hand.DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
