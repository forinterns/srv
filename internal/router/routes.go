package router

import (
	"github.com/forinterns/srv/internal/handlers/add"
	"github.com/forinterns/srv/internal/handlers/get"
	"net/http"
)

type Routes struct {
}

func NewRoutes() *Routes {
	return &Routes{}
}

func (r *Routes) Get() map[string]http.HandlerFunc {
	routes := make(map[string]http.HandlerFunc)
	routes["/employee/add"] = add.New().Handle
	routes["/employee/get"] = get.New().Handle
	return routes
}
