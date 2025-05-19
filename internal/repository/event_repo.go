package repository

import (
	"errors"
	"meeting-scheduler/internal/model"
	"sync"
)

type inMemoryEventRepo struct {
	data map[string]*model.Event
	mu   sync.RWMutex
}

func NewInMemoryEventRepository() EventRepository {
	return &inMemoryEventRepo{data: make(map[string]*model.Event)}
}
func (r *inMemoryEventRepo) Create(e *model.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[e.ID] = e
	return nil
}
func (r *inMemoryEventRepo) Get(id string) (*model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	event, ok := r.data[id]
	if !ok {
		return nil, errors.New("event not found")
	}
	return event, nil
}
func (r *inMemoryEventRepo) Update(e *model.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[e.ID]; !ok {
		return errors.New("event not found")
	}
	r.data[e.ID] = e
	return nil
}
func (r *inMemoryEventRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return errors.New("event not found")
	}
	delete(r.data, id)
	return nil
}
func (r *inMemoryEventRepo) List() []*model.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := []*model.Event{}
	for _, e := range r.data {
		list = append(list, e)
	}
	return list
}
func (r *inMemoryEventRepo) AllEventIds() (map[string]struct{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	set := make(map[string]struct{}, len(r.data))
	for id := range r.data {
		set[id] = struct{}{}
	}
	return set, nil
}
