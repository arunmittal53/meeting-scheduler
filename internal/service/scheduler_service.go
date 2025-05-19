package service

import (
	"meeting-scheduler/internal/repository"
)

type SchedulerService struct {
	userRepo         repository.UserRepository
	eventRepo        repository.EventRepository
	availabilityRepo repository.AvailabilityRepository
}

func NewSchedulerService(u repository.UserRepository, e repository.EventRepository, a repository.AvailabilityRepository) *SchedulerService {
	return &SchedulerService{userRepo: u, eventRepo: e, availabilityRepo: a}
}
