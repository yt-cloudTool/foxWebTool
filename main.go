package main

import (
	"net/http"
)

func main() {
	handler := Http_handler{}
	handler.ROUTES()

	http.ListenAndServe(server_url, nil)
}
