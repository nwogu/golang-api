package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (hanlder *Handler) SetupRoutes() {
	fmt.Println("-------Setting up Routes-------")
	hanlder.Router = mux.NewRouter()
	hanlder.Router.HandleFunc("/api/health-check", func(response http.ResponseWriter, request *http.Request){
		fmt.Fprintf(response, "Healthy")
	})
}