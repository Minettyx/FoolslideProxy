package main

import (
	"github.com/Minettyx/FoolslideProxy/pkg/server"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	if os.Getenv("SIGN_TOKEN") == "" {
		panic("Environment variable SIGN_TOKEN not found")
	}

	err := http.ListenAndServe(":"+port, server.Router())
	if err != nil {
		panic(err)
	}
}
