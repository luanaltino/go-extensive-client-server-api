package main

import (
	handler "go-extensive-client-server-api/server/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/quotation", handler.FindExchange)

	// inicializando o servidor
	http.ListenAndServe(":8080", nil)
}
