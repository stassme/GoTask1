package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestMsg struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var msg RequestMsg
	decoder.Decode(&msg)

	sendToDB := Message{TextName: msg.Name,
		TextDescription: msg.Description}

	DB.Create(&sendToDB)

	fmt.Fprintf(w, "The message %v %v is saved", msg.Name, msg.Description)
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	var messages []Message
	DB.Find(&messages)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var updatedInfo RequestMsg
	decoder.Decode(&updatedInfo)
	vars := mux.Vars(r)
	id := vars["id"]
	DB.Model(&Message{}).Where("id = ?", id).Updates(Message{
		TextName:        updatedInfo.Name,
		TextDescription: updatedInfo.Description,
	})

	fmt.Fprintf(w, "Message with id %v was update for %v %v", id, updatedInfo.Name, updatedInfo.Description)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	DB.Delete(&Message{}, id)
	fmt.Fprintf(w, "The resource was deleted")
}

func main() {
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/post", postHandler).Methods("POST")
	router.HandleFunc("/get", getHandler).Methods("GET")
	router.HandleFunc("/patch/{id}", patchHandler).Methods("PATCH")
	router.HandleFunc("/delete/{id}", deleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
