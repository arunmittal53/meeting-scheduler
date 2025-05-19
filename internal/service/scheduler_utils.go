package service

import (
	"fmt"
	"meeting-scheduler/internal/model"
)

func (s *SchedulerService) validateUserAndEventExist(av model.Availability) error {
	if err := s.ensureUsersExist(av.UserID); err != nil {
		return err
	}
	if _, err := s.ensureEventExists(av.EventID); err != nil {
		return err
	}
	return nil
}

func (s *SchedulerService) ensureUsersExist(userIDs ...string) error {
	missing, ok := s.validateUsersExistenceInSystem(userIDs...)
	if !ok {
		return fmt.Errorf("missing users: %v", missing)
	}
	return nil
}

func (s *SchedulerService) validateUsersExistenceInSystem(userIDs ...string) ([]string, bool) {
	allUsers, err := s.userRepo.GetAll()
	if err != nil {
		return nil, false
	}
	var missing []string
	for _, id := range userIDs {
		if _, ok := allUsers[id]; !ok {
			missing = append(missing, id)
		}
	}
	return missing, len(missing) == 0
}

func (s *SchedulerService) ensureEventExists(eventId string) (*model.Event, error) {
	event, _ := s.eventRepo.Get(eventId)
	if event == nil {
		return nil, fmt.Errorf("event with ID %s does not exist", eventId)
	}
	return event, nil
}

func (s *SchedulerService) validateUserAndEventExistByIDs(eventID, userID string) error {
	if _, err := s.GetEvent(eventID); err != nil {
		return err
	}
	if _, err := s.GetUser(userID); err != nil {
		return err
	}
	return nil
}

func isUserAvailableForExactWindow(target model.Slot, slots []model.Slot) bool {
	for _, s := range slots {
		if !s.Start.After(target.Start) && !s.End.Before(target.End) {
			return true
		}
	}
	return false
}

func getMissingUsers2(all []string, present []string) []string {
	set := make(map[string]struct{}, len(present))
	for _, u := range present {
		set[u] = struct{}{}
	}
	var missing []string
	for _, u := range all {
		if _, ok := set[u]; !ok {
			missing = append(missing, u)
		}
	}
	return missing
}
