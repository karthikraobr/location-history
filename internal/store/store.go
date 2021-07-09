package store

import (
	"errors"
	"log"
	"sync"
)

type Location struct {
	Lat float64
	Lng float64
}

type Status int

const (
	Created Status = iota
	Updated
)

type Store struct {
	data map[string][]Location
	l    sync.RWMutex
	log  *log.Logger
}

func New(log *log.Logger) *Store {
	return &Store{
		data: map[string][]Location{},
		log:  log,
	}
}

func (s *Store) UpdateHistory(orderId string, location Location) ([]Location, Status, error) {
	s.l.Lock()
	defer s.l.Unlock()
	s.data[orderId] = append([]Location{location}, s.data[orderId]...)
	status := Created
	if len(s.data[orderId]) > 1 {
		status = Updated
	}
	return s.data[orderId], status, nil
}

func (s *Store) GetHistory(orderId string, limit int) ([]Location, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	r, ok := s.data[orderId]
	if !ok {
		return nil, errors.New("invalid orderid")
	}
	if limit > 0 && limit < len(r) {
		return r[:limit], nil
	}
	return r, nil
}

func (s *Store) DeleteHistory(orderId string) error {
	if _, ok := s.data[orderId]; !ok {
		return errors.New("invalid orderid")
	}
	s.l.Lock()
	defer s.l.Unlock()
	delete(s.data, orderId)
	return nil
}
