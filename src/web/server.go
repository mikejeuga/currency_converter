package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/web/auth"
	"net/http"
)

//go:generate moq -out mocks/gateway_moq.go -pkg=mocks . Gateway
type Gateway interface {
	GetRate(base, foreign string) (models.Rate, error)
	Convert(amount, baseCurrency, foreignCurrency string) (models.Amount, error)
}

type Server struct {
	config.Config
	converter Gateway
}

func NewServer(conf config.Config, converter Gateway) *http.Server {
	r := mux.NewRouter()
	s := Server{
		converter: converter,
	}

	r.Use(auth.NewMiddleware(conf))

	r.HandleFunc("/", s.Home).Methods(http.MethodGet)
	r.HandleFunc("/rate", s.GetRate).Methods(http.MethodGet)
	r.HandleFunc("/converted-amount", s.Convert).Methods(http.MethodGet)

	return &http.Server{
		Addr:    ":8002",
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

	rate, err := s.converter.GetRate(baseCurrency, fxCurrency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(rate.Spot)

	err = json.NewEncoder(w).Encode(rate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (s *Server) Convert(w http.ResponseWriter, r *http.Request) {
	baseAmount := r.URL.Query().Get("amount")
	baseCurrency := r.URL.Query().Get("have")
	fxCurrency := r.URL.Query().Get("want")

	if baseCurrency == "" || fxCurrency == "" || baseAmount == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convert, err := s.converter.Convert(baseAmount, baseCurrency, fxCurrency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(convert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
