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

var message string = "Lets GOO"

func postHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&requestBody)
	fmt.Fprintf(w, "Получено сообщение %v", message)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %v", message)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/post", postHandler)
	router.HandleFunc("/get", getHandler)

	http.ListenAndServe(":8080", router)
}
