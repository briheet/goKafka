package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("POST /user", s.HandleUserRegister)

	server := http.Server{
		Addr:    s.port,
		Handler: router,
	}

	log.Printf("Server has started at %s", s.port)

	return server.ListenAndServe()
}

func (s *ApiServer) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
