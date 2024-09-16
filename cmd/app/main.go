package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stassme/GoTask1/internal/database"
	"github.com/stassme/GoTask1/internal/handlers"
	"github.com/stassme/GoTask1/internal/messagesService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.GET("/get", handler.GetMessagesHandler)
	e.POST("/post", handler.PostMessageHandler)
	e.DELETE("/delete/:id", handler.DeleteMessageHandler)
	e.PATCH("/patch/:id", handler.PatchMessageHandler)

	e.Start(":8080")
}
