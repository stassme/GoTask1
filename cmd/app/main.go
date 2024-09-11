package main

import (
	"github.com/gorilla/mux"
	"github.com/stassme/GoTask1/internal/database"
	"github.com/stassme/GoTask1/internal/handlers"
	"github.com/stassme/GoTask1/internal/messagesService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/post", handler.PostMessageHandler).Methods("POST")

	router.HandleFunc("/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")

	router.HandleFunc("/patch/{id}", handler.PatchMessageHandler).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}
