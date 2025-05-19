package service_test

import (
	"errors"
	"meeting-scheduler/internal/model"
	"meeting-scheduler/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepo is a mocked implementation of UserRepository
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Get(id string) (*model.User, error) {
	args := m.Called(id)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*model.User), args.Error(1)
}

func (m *MockUserRepo) GetAll() (map[string]*model.User, error) {
	args := m.Called()
	return args.Get(0).(map[string]*model.User), args.Error(1)
}

func (m *MockUserRepo) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func setup() (*service.SchedulerService, *MockUserRepo) {
	mockRepo := new(MockUserRepo)
	// Use a constructor or exported fields to set dependencies
	svc := service.NewSchedulerService(mockRepo, nil, nil)
	return svc, mockRepo
}

func TestGetUser_Success(t *testing.T) {
	svc, mockRepo := setup()
	user := &model.User{ID: "123", Name: "Alice"}

	mockRepo.On("Get", "123").Return(user, nil)

	result, err := svc.GetUser("123")
	assert.NoError(t, err)
	assert.Equal(t, user, result)

	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	svc, mockRepo := setup()

	mockRepo.On("Get", "999").Return(nil, errors.New("not found"))

	result, err := svc.GetUser("999")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "user with ID 999 not found")

	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	svc, mockRepo := setup()
	users := map[string]*model.User{
		"1": {ID: "1", Name: "Alice"},
		"2": {ID: "2", Name: "Bob"},
	}

	mockRepo.On("GetAll").Return(users, nil)

	result, err := svc.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	svc, mockRepo := setup()
	newUser := &model.User{ID: "100", Name: "Charlie"}

	mockRepo.On("Get", "100").Return(nil, nil)
	mockRepo.On("Create", newUser).Return(nil)

	err := svc.CreateUser(newUser)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_AlreadyExists(t *testing.T) {
	svc, mockRepo := setup()
	existingUser := &model.User{ID: "1", Name: "Alice"}

	mockRepo.On("Get", "1").Return(existingUser, nil)

	err := svc.CreateUser(existingUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user with ID 1 already exists")

	mockRepo.AssertExpectations(t)
}
