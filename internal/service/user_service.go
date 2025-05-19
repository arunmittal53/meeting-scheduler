package service

import (
	"fmt"
	"meeting-scheduler/internal/model"
)

func (s *SchedulerService) GetUser(id string) (*model.User, error) {
	user, err := s.userRepo.Get(id)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (s *SchedulerService) GetAllUsers() ([]*model.User, error) {
	userMap, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0, len(userMap))
	for _, user := range userMap {
		users = append(users, user)
	}
	return users, nil
}

func (s *SchedulerService) CreateUser(u *model.User) error {
	if existing, _ := s.userRepo.Get(u.ID); existing != nil {
		return fmt.Errorf("user with ID %s already exists", u.ID)
	}
	return s.userRepo.Create(u)
}
