package repository

import (
	"fmt"
	"maps"
	"meeting-scheduler/internal/model"
	"sync"
)

// [event -> [user -> Availability]]
type inMemoryAvailabilityRepo struct {
	data map[string]map[string]model.Availability
	mu   sync.RWMutex
}

func NewInMemoryAvailabilityRepository() AvailabilityRepository {
	return &inMemoryAvailabilityRepo{data: make(map[string]map[string]model.Availability)}
}

func (r *inMemoryAvailabilityRepo) Get(eventID, userID string) (model.Availability, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	eventData, ok := r.data[eventID]
	if !ok {
		return model.Availability{}, fmt.Errorf("event not found: %s", eventID)
	}
	availability, ok := eventData[userID]
	if !ok {
		return model.Availability{}, fmt.Errorf("availability not found for user %s in event %s", userID, eventID)
	}
	return availability, nil
}
func (r *inMemoryAvailabilityRepo) Create(av model.Availability) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[av.EventID]; !ok {
		r.data[av.EventID] = make(map[string]model.Availability)
	} else {
		if _, ok := r.data[av.EventID][av.UserID]; ok {
			return fmt.Errorf("availability already exists for user %s in event %s", av.UserID, av.EventID)
		}
	}
	r.data[av.EventID][av.UserID] = av
	return nil
}
func (r *inMemoryAvailabilityRepo) Update(av model.Availability) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[av.EventID]; !ok {
		return fmt.Errorf("event not found: %s", av.EventID)
	}
	if _, ok := r.data[av.EventID][av.UserID]; !ok {
		return fmt.Errorf("availability not found in event: %s for user : %s", av.EventID, av.UserID)
	}
	r.data[av.EventID][av.UserID] = av
	return nil
}
func (r *inMemoryAvailabilityRepo) GetByEvent(eventID string) map[string]model.Availability {
	r.mu.RLock()
	defer r.mu.RUnlock()
	src, ok := r.data[eventID]
	if !ok {
		return nil
	}
	copy := make(map[string]model.Availability, len(src))
	maps.Copy(copy, src)
	return copy
}
func (r *inMemoryAvailabilityRepo) Delete(eventID string, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[eventID]; !ok {
		return fmt.Errorf("event not found: %s", eventID)
	}
	if _, ok := r.data[eventID][userID]; !ok {
		return fmt.Errorf("availability not found for user %s in event %s", userID, eventID)
	}
	delete(r.data[eventID], userID)
	return nil
}
