package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

type User struct {
	ID        int64     `json:"id"`
	Topic     string    `json:"topic"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *ApiServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/user", s.HandleUserRegister).Methods("POST")

	log.Printf("starting the api at %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, subrouter))
}

func (s *ApiServer) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ending mein error hain")
		return
	}
}
