package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	port string
}

func NewApiServer(port string) *ApiServer {
	return &ApiServer{
		port: port,
	}
}

func (s *ApiServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	user := &User{}
	user.RegisterRoute(*subrouter)

	log.Printf("starting the api at %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, subrouter))
}
