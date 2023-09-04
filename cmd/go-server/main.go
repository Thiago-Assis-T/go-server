package main

import (
	"fmt"
	"net/http"

	"github.com/Thiago-Assis-T/go-server/internal/routes"
)

func main() {
	router := routes.NewRouter()
	port := 42069
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server listening on https://localhost%s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
