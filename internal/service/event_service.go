package service

import (
	"fmt"
	"meeting-scheduler/internal/model"
)

func (s *SchedulerService) GetEvent(id string) (*model.Event, error) {
	event, err := s.eventRepo.Get(id)
	if err != nil || event == nil {
		return nil, fmt.Errorf("event with ID %s not found", id)
	}
	return event, nil
}

func (s *SchedulerService) CreateEvent(e *model.Event) error {
	if len(e.Participants) == 0 {
		return fmt.Errorf("event must have at least one participant")
	}
	if err := s.ensureUsersExist(e.Participants...); err != nil {
		return err
	}
	if existing, _ := s.eventRepo.Get(e.ID); existing != nil {
		return fmt.Errorf("event with ID %s already exists", e.ID)
	}
	return s.eventRepo.Create(e)
}

func (s *SchedulerService) UpdateEvent(e *model.Event) error {
	if len(e.Participants) == 0 {
		return fmt.Errorf("event must have at least one participant")
	}
	if err := s.ensureUsersExist(e.Participants...); err != nil {
		return err
	}
	if existing, err := s.eventRepo.Get(e.ID); err != nil || existing == nil {
		return fmt.Errorf("event with ID %s does not exist", e.ID)
	}
	return s.eventRepo.Update(e)
}

func (s *SchedulerService) DeleteEvent(id string) error {
	if id == "" {
		return fmt.Errorf("event ID cannot be empty")
	}
	if existing, err := s.eventRepo.Get(id); err != nil || existing == nil {
		return fmt.Errorf("event with ID %s does not exist", id)
	}
	return s.eventRepo.Delete(id)
}
