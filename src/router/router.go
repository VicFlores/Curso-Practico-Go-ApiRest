package router

import (
	"net/http"
)

type Router struct {
	Rules map[string]map[string]http.HandlerFunc // metodo/url/handler
}

// Procesar las rutas
func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.Rules[path]                     // ruta
	handler, methodExist := r.Rules[path][method] // metodo

	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	handler, method, exist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if !method {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(response, request)
}
