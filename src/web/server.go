package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
}

func NewServer() *http.Server {
	r := mux.NewRouter()
	s := Server{}

	r.HandleFunc("/", s.Home).Methods(http.MethodGet)

	return &http.Server{
		Addr:    ":8087",
		Handler: r,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Cool, the server is up and running!")
}
