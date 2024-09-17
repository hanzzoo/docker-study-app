package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func setupHeader(response http.ResponseWriter, request *http.Request) {
	origin := request.Header.Get("Origin")
	if origin == "http://localhost" || origin == "http://localhost:3000" {
		response.Header().Set("Access-Control-Allow-Origin", origin)
	}
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "application/json")
}

type SampleHandler struct{}

func (h *SampleHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	setupHeader(response, request)
	var sampleData = map[string]string{
		"message": "Hello, World!",
	}
	result, err := json.Marshal(sampleData)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if _, err := response.Write(result); err != nil {
		fmt.Println("Error: ", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	const PORT = 4000
	server := http.Server{
		Addr: fmt.Sprintf(":%d", PORT),
	}

	http.Handle("/sample", &SampleHandler{})

	if server.ListenAndServe() != nil {
		os.Exit(1)
	}

	fmt.Println("Server is running on port", PORT)
}
