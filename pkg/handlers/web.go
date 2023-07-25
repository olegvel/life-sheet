package handlers

import "net/http"

func NewWebServer() http.Handler {
	return http.FileServer(http.Dir("./web"))
}
