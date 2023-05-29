package main

import (
	"log"
	"net/http"
	"project/internal/routers"
)

func main() {
	router := routers.SetupRouter()

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}


