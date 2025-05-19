package repository

import (
	"maps"
	"meeting-scheduler/internal/model"
	"sync"
)

type inMemoryUserRepo struct {
	user map[string]*model.User
	mu   sync.RWMutex
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepo{user: make(map[string]*model.User)}
}

func (r *inMemoryUserRepo) Get(id string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.user[id], nil
}
func (r *inMemoryUserRepo) GetAll() (map[string]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	copy := make(map[string]*model.User, len(r.user))
	maps.Copy(copy, r.user)
	return copy, nil
}
func (r *inMemoryUserRepo) Create(e *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.user[e.ID] = e
	return nil
}
