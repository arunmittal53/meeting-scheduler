package service_test

/*
import (
	"errors"
	"meeting-scheduler/internal/model"
	"meeting-scheduler/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ---- Mock availability repo ----

type mockAvailabilityRepo struct {
	store map[string]map[string]model.Availability
}

func newMockAvailabilityRepo() *mockAvailabilityRepo {
	return &mockAvailabilityRepo{store: make(map[string]map[string]model.Availability)}
}

func (m *mockAvailabilityRepo) Get(eventID, userID string) (model.Availability, error) {
	if m.store[eventID] == nil || m.store[eventID][userID].UserID == "" {
		return model.Availability{}, errors.New("not found")
	}
	return m.store[eventID][userID], nil
}
func (m *mockAvailabilityRepo) Create(av model.Availability) error {
	if m.store[av.EventID] == nil {
		m.store[av.EventID] = make(map[string]model.Availability)
	}
	m.store[av.EventID][av.UserID] = av
	return nil
}
func (m *mockAvailabilityRepo) Update(av model.Availability) error {
	if m.store[av.EventID] == nil || m.store[av.EventID][av.UserID].UserID == "" {
		return errors.New("not found")
	}
	m.store[av.EventID][av.UserID] = av
	return nil
}
func (m *mockAvailabilityRepo) Delete(eventID, userID string) error {
	if m.store[eventID] == nil || m.store[eventID][userID].UserID == "" {
		return errors.New("not found")
	}
	delete(m.store[eventID], userID)
	return nil
}

// ---- Minimal mocks for user and event repo ----

type dummyUserRepo struct {
	users map[string]*model.User
}

func (m *dummyUserRepo) Get(id string) (*model.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}
func (m *dummyUserRepo) GetAll() (map[string]*model.User, error) { return m.users, nil }
func (m *dummyUserRepo) Create(u *model.User) error {
	m.users[u.ID] = u
	return nil
}

type dummyEventRepo struct {
	events map[string]*model.Event
}

func (m *dummyEventRepo) Get(id string) (*model.Event, error) {
	e, ok := m.events[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return e, nil
}
func (m *dummyEventRepo) Create(e *model.Event) error               { m.events[e.ID] = e; return nil }
func (m *dummyEventRepo) Update(e *model.Event) error               { return nil }
func (m *dummyEventRepo) Delete(id string) error                    { return nil }
func (m *dummyEventRepo) List() []*model.Event                      { return nil }
func (m *dummyEventRepo) AllEventIds() (map[string]struct{}, error) { return nil, nil }

func TestAddAvailability_Success(t *testing.T) {
	userRepo := &dummyUserRepo{users: map[string]*model.User{"u1": {ID: "u1", Name: "Alice"}}}
	eventRepo := &dummyEventRepo{events: map[string]*model.Event{"e1": {ID: "e1"}}}
	availabilityRepo := newMockAvailabilityRepo()

	svc := service.NewSchedulerService(userRepo, eventRepo, availabilityRepo)

	av := model.Availability{
		EventID: "e1",
		UserID:  "u1",
		Slots:   []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
	}

	err := svc.AddAvailability(av)
	assert.NoError(t, err)
}

func TestAddAvailability_InvalidUser(t *testing.T) {
	userRepo := &dummyUserRepo{users: map[string]*model.User{}}
	eventRepo := &dummyEventRepo{events: map[string]*model.Event{"e1": {ID: "e1"}}}
	availabilityRepo := newMockAvailabilityRepo()

	svc := service.NewSchedulerService(userRepo, eventRepo, availabilityRepo)

	av := model.Availability{EventID: "e1", UserID: "uX"}
	err := svc.AddAvailability(av)
	assert.Error(t, err)
}

func TestGetAvailability_Success(t *testing.T) {
	userRepo := &dummyUserRepo{users: map[string]*model.User{"u1": {ID: "u1"}}}
	eventRepo := &dummyEventRepo{events: map[string]*model.Event{"e1": {ID: "e1"}}}
	availabilityRepo := newMockAvailabilityRepo()

	svc := service.NewSchedulerService(userRepo, eventRepo, availabilityRepo)

	av := model.Availability{EventID: "e1", UserID: "u1"}
	_ = availabilityRepo.Create(av)

	got, err := svc.GetAvailability("e1", "u1")
	assert.NoError(t, err)
	assert.Equal(t, "u1", got.UserID)
}

func TestUpdateAvailability_NotExists(t *testing.T) {
	userRepo := &dummyUserRepo{users: map[string]*model.User{"u1": {ID: "u1"}}}
	eventRepo := &dummyEventRepo{events: map[string]*model.Event{"e1": {ID: "e1"}}}
	availabilityRepo := newMockAvailabilityRepo()

	svc := service.NewSchedulerService(userRepo, eventRepo, availabilityRepo)

	av := model.Availability{EventID: "e1", UserID: "u1"}
	err := svc.UpdateAvailability(av)
	assert.Error(t, err)
}

func TestDeleteAvailability_Success(t *testing.T) {
	userRepo := &dummyUserRepo{users: map[string]*model.User{"u1": {ID: "u1"}}}
	eventRepo := &dummyEventRepo{events: map[string]*model.Event{"e1": {ID: "e1"}}}
	availabilityRepo := newMockAvailabilityRepo()

	svc := service.NewSchedulerService(userRepo, eventRepo, availabilityRepo)

	_ = availabilityRepo.Create(model.Availability{EventID: "e1", UserID: "u1"})
	err := svc.DeleteAvailability("e1", "u1")
	assert.NoError(t, err)
}
*/
