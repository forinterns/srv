package main

import (
	"github.com/forinterns/srv/internal/router"
	"net/http"
)

func main() {
	for url, handlerFunc := range router.NewRoutes().Get() {
		http.Handle(url, handlerFunc)
	}
	http.ListenAndServe(":8080", nil)
}
