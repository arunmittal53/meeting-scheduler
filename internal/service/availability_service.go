package service

import (
	"meeting-scheduler/internal/model"
)

func (s *SchedulerService) GetAvailability(eventID, userID string) (model.Availability, error) {
	if err := s.validateUserAndEventExistByIDs(eventID, userID); err != nil {
		return model.Availability{}, err
	}
	return s.availabilityRepo.Get(eventID, userID)
}

func (s *SchedulerService) AddAvailability(av model.Availability) error {
	if err := s.validateUserAndEventExist(av); err != nil {
		return err
	}
	return s.availabilityRepo.Create(av)
}

func (s *SchedulerService) UpdateAvailability(av model.Availability) error {
	if err := s.validateUserAndEventExist(av); err != nil {
		return err
	}
	return s.availabilityRepo.Update(av)
}

func (s *SchedulerService) DeleteAvailability(eventID, userID string) error {
	return s.availabilityRepo.Delete(eventID, userID)
}
