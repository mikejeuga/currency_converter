package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/currency_converter/specifications"
	"net/http"
)

type Server struct {
	converter specifications.Converter
}

func NewServer(converter specifications.Converter) *http.Server {
	r := mux.NewRouter()
	s := Server{
		converter: converter,
	}

	r.HandleFunc("/", s.Home).Methods(http.MethodGet)
	r.HandleFunc("/rate", s.GetRate).Methods(http.MethodGet)

	return &http.Server{
		Addr:    ":8077",
		Handler: r,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Cool, the server is up and running!")
}

func (s *Server) GetRate(w http.ResponseWriter, r *http.Request) {
	baseCurrency := r.URL.Query().Get("have")
	fxCurrency := r.URL.Query().Get("want")

	if baseCurrency == "" || fxCurrency == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rate, err := s.converter.GetFXRate(baseCurrency, fxCurrency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(rate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
