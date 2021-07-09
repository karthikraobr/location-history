package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/karthikrao/location-history/internal/store"
)

type server struct {
	store  *store.Store
	logger *log.Logger
}

func New(store *store.Store, logger *log.Logger) *server {
	return &server{
		store:  store,
		logger: logger,
	}
}

func (s *server) Router() *http.ServeMux {
	mux := http.ServeMux{}
	mux.Handle("/location/", s.handleHistory())
	return &mux
}

func (s *server) handleHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/location/")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodPut:
			s.processPut(id, w, r)
		case http.MethodGet:
			s.processGet(id, w, r)
		case http.MethodDelete:
			s.processDelete(id, w)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func (s *server) processPut(id string, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var location store.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, state, err := s.store.UpdateHistory(id, location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(getHttpStatus(state))
}

func (s *server) processGet(id string, w http.ResponseWriter, r *http.Request) {
	max := r.URL.Query().Get("max")
	var limit int
	limit, _ = strconv.Atoi(max)
	locations, err := s.store.GetHistory(id, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := struct {
		OrderID string           `json:"order_id"`
		History []store.Location `json:"history"`
	}{
		OrderID: id,
		History: locations,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (s *server) processDelete(id string, w http.ResponseWriter) {
	if err := s.store.DeleteHistory(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getHttpStatus(s store.Status) int {
	if s == store.Created {
		return http.StatusCreated
	}
	return http.StatusOK
}
