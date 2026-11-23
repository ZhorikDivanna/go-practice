package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *GameHandlers
}

func NewHTTPServer(httpHandler *GameHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/miners").Methods("POST").HandlerFunc(s.handlers.HandleHire)
	router.Path("/equipment").Methods("POST").HandlerFunc(s.handlers.HandleBuy)
	router.Path("/status").Methods("GET").HandlerFunc(s.handlers.HandleStatus)
	router.Path("/finish").Methods("POST").HandlerFunc(s.handlers.HandleFinish)
	router.Path("/miners/prices").Methods("GET").HandlerFunc(s.handlers.HandleMinerPrices)
	router.Path("/equipment/prices").Methods("GET").HandlerFunc(s.handlers.HandleEquipmentPrices)
	router.Path("/miners/active").Methods("GET").HandlerFunc(s.handlers.HandleActiveMiners)

	if err := http.ListenAndServe(":9090", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
