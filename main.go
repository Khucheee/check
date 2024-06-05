package main

import (
	"encoding/json"
	"net/http"
)

type Answer struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", checkHealth)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		println("ну что теперь, не запустился твой сервер")
	}
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	answer := Answer{"UP"}
	resp, err := json.Marshal(answer)
	if err != nil {
		println("что-то отвалилось при создании json")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
