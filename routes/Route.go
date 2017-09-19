package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct{}

func (r *Route) Init() *http.ServeMux {
	route := mux.NewRouter()
	route.HandleFunc("/{text}", GetEnv)

	routeMiddleware := http.NewServeMux()
	routeMiddleware.Handle("/", route)
	return routeMiddleware
}

func GetEnv(res http.ResponseWriter, req *http.Request) {
	input := mux.Vars(req)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Hello %v \n", input["text"])
}
