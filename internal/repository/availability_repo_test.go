package repository_test

import (
	"meeting-scheduler/internal/model"
	"meeting-scheduler/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInMemoryAvailabilityRepo_CreateGet tests the Create and Get methods of InMemoryAvailabilityRepository
func TestInMemoryAvailabilityRepo_CreateGet(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
			{
				Start: time.Date(2025, time.May, 21, 14, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 21, 15, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	gotAvailability, err := repo.Get("event1", "user1")
	require.NoError(t, err)
	assert.Equal(t, availability, gotAvailability)
}

// TestInMemoryAvailabilityRepo_Update tests the Update method of InMemoryAvailabilityRepository
func TestInMemoryAvailabilityRepo_Update(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	// Update the availability
	updatedAvailability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 21, 14, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 21, 15, 0, 0, 0, time.UTC),
			},
		},
	}

	err = repo.Update(updatedAvailability)
	require.NoError(t, err)

	gotAvailability, err := repo.Get("event1", "user1")
	require.NoError(t, err)
	assert.Equal(t, updatedAvailability.Slots[0], gotAvailability.Slots[0])
}

// TestInMemoryAvailabilityRepo_GetByEvent tests the GetByEvent method of InMemoryAvailabilityRepository
func TestInMemoryAvailabilityRepo_GetByEvent(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability1 := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability1)
	require.NoError(t, err)

	availability2 := model.Availability{
		EventID: "event1",
		UserID:  "user2",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 21, 14, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 21, 15, 0, 0, 0, time.UTC),
			},
		},
	}

	err = repo.Create(availability2)
	require.NoError(t, err)

	gotAvailability := repo.GetByEvent("event1")
	assert.Len(t, gotAvailability, 2)
}

// TestInMemoryAvailabilityRepo_GetByEvent_NoData tests the GetByEvent method of InMemoryAvailabilityRepository when no data is present
func TestInMemoryAvailabilityRepo_GetByEvent_NoData(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	gotAvailability := repo.GetByEvent("event1")
	assert.Empty(t, gotAvailability)
}

// TestInMemoryAvailabilityRepo_Update_NonExistent tests the Update method of InMemoryAvailabilityRepository when the availability does not exist
func TestInMemoryAvailabilityRepo_Update_NonExistent(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Update(availability)
	assert.EqualError(t, err, "event not found: event1")
}

// TestInMemoryAvailabilityRepo_Get_NonExistent tests the Get method of InMemoryAvailabilityRepository when the availability does not exist
func TestInMemoryAvailabilityRepo_Get_NonExistent(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	_, err := repo.Get("event1", "user1")
	assert.EqualError(t, err, "event not found: event1")
}

// TestInMemoryAvailabilityRepo_Get_NonExistentUser tests the Get method of InMemoryAvailabilityRepository when the user does not exist
func TestInMemoryAvailabilityRepo_Get_NonExistentUser(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	_, err = repo.Get("event1", "user2")
	assert.EqualError(t, err, "availability not found for user user2 in event event1")
}

// TestInMemoryAvailabilityRepo_GetByEvent_NonExistent tests the GetByEvent method of InMemoryAvailabilityRepository when the event does not exist
func TestInMemoryAvailabilityRepo_GetByEvent_NonExistent(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	gotAvailability := repo.GetByEvent("event1")
	assert.Empty(t, gotAvailability)
}

// TestInMemoryAvailabilityRepo_GetByEvent_Empty tests the GetByEvent method of InMemoryAvailabilityRepository when the event has no availability
func TestInMemoryAvailabilityRepo_GetByEvent_Empty(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	gotAvailability := repo.GetByEvent("event1")
	assert.Empty(t, gotAvailability)
}

// TestInMemoryAvailabilityRepo_GetByEvent_OneUser tests the GetByEvent method of InMemoryAvailabilityRepository when there is one user
func TestInMemoryAvailabilityRepo_GetByEvent_OneUser(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	gotAvailability := repo.GetByEvent("event1")
	assert.Len(t, gotAvailability, 1)
}

// TestInMemoryAvailabilityRepo_GetByEvent_MultipleUsers tests the GetByEvent method of InMemoryAvailabilityRepository when there are multiple users
func TestInMemoryAvailabilityRepo_GetByEvent_MultipleUsers(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability1 := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability1)
	require.NoError(t, err)

	availability2 := model.Availability{
		EventID: "event1",
		UserID:  "user2",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 21, 14, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 21, 15, 0, 0, 0, time.UTC),
			},
		},
	}

	err = repo.Create(availability2)
	require.NoError(t, err)

	gotAvailability := repo.GetByEvent("event1")
	assert.Len(t, gotAvailability, 2)
}

// TestInMemoryAvailabilityRepo_Delete tests the Delete method of InMemoryAvailabilityRepository
func TestInMemoryAvailabilityRepo_Delete(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	err = repo.Delete("event1", "user1")
	require.NoError(t, err)

	_, err = repo.Get("event1", "user1")
	assert.EqualError(t, err, "availability not found for user user1 in event event1")
}

// TestInMemoryAvailabilityRepo_Delete_NonExistent tests the Delete method of InMemoryAvailabilityRepository when the event does not exist
func TestInMemoryAvailabilityRepo_Delete_NonExistent(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	err := repo.Delete("event1", "user1")
	assert.EqualError(t, err, "event not found: event1")
}

// TestInMemoryAvailabilityRepo_Delete_NonExistentUser tests the Delete method of InMemoryAvailabilityRepository when the user does not exist
func TestInMemoryAvailabilityRepo_Delete_NonExistentUser(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	err = repo.Delete("event1", "user2")
	assert.EqualError(t, err, "availability not found for user user2 in event event1")
}

// TestInMemoryAvailabilityRepo_Delete_Empty tests the Delete method of InMemoryAvailabilityRepository when there is no data
func TestInMemoryAvailabilityRepo_Delete_Empty(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	err := repo.Delete("event1", "user1")
	assert.EqualError(t, err, "event not found: event1")
}

// TestInMemoryAvailabilityRepo_Delete_OneUser tests the Delete method of InMemoryAvailabilityRepository when there is one user
func TestInMemoryAvailabilityRepo_Delete_OneUser(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability)
	require.NoError(t, err)

	err = repo.Delete("event1", "user1")
	require.NoError(t, err)

	gotAvailability := repo.GetByEvent("event1")
	assert.Empty(t, gotAvailability)
}

// TestInMemoryAvailabilityRepo_Delete_MultipleUsers tests the Delete method of InMemoryAvailabilityRepository when there are multiple users
func TestInMemoryAvailabilityRepo_Delete_MultipleUsers(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability1 := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability1)
	require.NoError(t, err)

	availability2 := model.Availability{
		EventID: "event1",
		UserID:  "user2",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 21, 14, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 21, 15, 0, 0, 0, time.UTC),
			},
		},
	}

	err = repo.Create(availability2)
	require.NoError(t, err)

	err = repo.Delete("event1", "user1")
	require.NoError(t, err)

	gotAvailability := repo.GetByEvent("event1")
	assert.Len(t, gotAvailability, 1)
}

// TestInMemoryAvailabilityRepo_Delete_All tests the Delete method of InMemoryAvailabilityRepository when all data is deleted
func TestInMemoryAvailabilityRepo_Delete_All(t *testing.T) {
	repo := repository.NewInMemoryAvailabilityRepository()

	availability1 := model.Availability{
		EventID: "event1",
		UserID:  "user1",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 20, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	err := repo.Create(availability1)
	require.NoError(t, err)

	availability2 := model.Availability{
		EventID: "event1",
		UserID:  "user2",
		Slots: []model.Slot{
			{
				Start: time.Date(2025, time.May, 21, 14, 0, 0, 0, time.UTC),
				End:   time.Date(2025, time.May, 21, 15, 0, 0, 0, time.UTC),
			},
		},
	}

	err = repo.Create(availability2)
	require.NoError(t, err)

	err = repo.Delete("event1", "user1")
	require.NoError(t, err)

	err = repo.Delete("event1", "user2")
	require.NoError(t, err)

	gotAvailability := repo.GetByEvent("event1")
	assert.Empty(t, gotAvailability)
}
