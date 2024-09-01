package main

import (
	"log"

	data "first/internal/database"
	hand "first/internal/handlers"
	ms "first/internal/messagesService"

	"first/internal/web/messages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DB = data.DB

// rdhdrrth sdfgazsdfgz
func main() {

	//инициализация переменной DB как базы данных
	DB = data.InitDB()
	//создание струтуры этой базы
	err := DB.AutoMigrate(ms.Message{})
	if err != nil {
		log.Fatalf("Ошибка при миграции: %v", err)
	}

	//создание струтуры и методов работы с этой струтурой репозитория
	repo := ms.NewMessageRepository(data.DB)
	//Создание методов работы с репозиторием
	service := ms.NewService(repo)

	//создание методов работы с сервисом
	handler := hand.NewHandler(service)

	// создем эхо для маршрутизации http запросов
	e := echo.New()

	//Логирует каждый входящий HTTP-запрос и его параметры
	e.Use(middleware.Logger())

	//Восстанавливает приложение после паники, предотвращая падение сервера
	//и возвращая ответ с ошибкой.
	e.Use(middleware.Recover())

	strictHandler := messages.NewStrictHandler(handler, nil)
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("faoled to start with err: %v", err)
	}
}
