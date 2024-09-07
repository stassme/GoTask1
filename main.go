package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&requestBody)

	newMessage := Message{Text: requestBody.Message}

	DB.Create(&newMessage)

	fmt.Fprintf(w, "Сообщение сохранено %v", newMessage.Text)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/post", postHandler)
	router.HandleFunc("/get", getHandler)

	http.ListenAndServe(":8080", router)
}
