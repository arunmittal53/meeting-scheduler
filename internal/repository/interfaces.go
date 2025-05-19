package repository

import "meeting-scheduler/internal/model"

type UserRepository interface {
	Get(id string) (*model.User, error)
	GetAll() (map[string]*model.User, error)
	Create(user *model.User) error
}

type EventRepository interface {
	Create(event *model.Event) error
	Get(id string) (*model.Event, error)
	Update(event *model.Event) error
	Delete(id string) error
	List() []*model.Event
	AllEventIds() (map[string]struct{}, error)
}

type AvailabilityRepository interface {
	Get(eventID, userID string) (model.Availability, error)
	Create(av model.Availability) error
	Update(av model.Availability) error
	GetByEvent(eventID string) map[string]model.Availability // user -> Availability
	Delete(eventID, userID string) error
}
