package repository

import (
	"meeting-scheduler/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestInMemoryEventRepo_CreateGet tests the Create and Get methods of InMemoryEventRepository
func TestInMemoryEventRepo_CreateGet(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	gotEvent, err := repo.Get("e1")
	assert.NoError(t, err)
	assert.Equal(t, event, gotEvent)
}

// TestInMemoryEventRepo_GetAll tests the GetAll method of InMemoryEventRepository
func TestInMemoryEventRepo_AllEventIds(t *testing.T) {
	repo := NewInMemoryEventRepository()

	event1 := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}
	event2 := &model.Event{
		ID:           "e2",
		Title:        "Conference",
		DurationMin:  120,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(2 * time.Hour)}},
		Participants: []string{"u3", "u4"},
	}

	err := repo.Create(event1)
	assert.NoError(t, err)
	err = repo.Create(event2)
	assert.NoError(t, err)

	allEvents, err := repo.AllEventIds()
	assert.NoError(t, err)
	assert.Len(t, allEvents, 2)
	assert.Contains(t, allEvents, "e1")
	assert.Contains(t, allEvents, "e2")
}

// TestInMemoryEventRepo_Update tests the Update method of InMemoryEventRepository
func TestInMemoryEventRepo_Update(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	event.Title = "Updated Meeting"
	err = repo.Update(event)
	assert.NoError(t, err)

	gotEvent, err := repo.Get("e1")
	assert.NoError(t, err)
	assert.Equal(t, event, gotEvent)
}

// TestInMemoryEventRepo_Delete tests the Delete method of InMemoryEventRepository
func TestInMemoryEventRepo_Delete(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	err = repo.Delete("e1")
	assert.NoError(t, err)

	gotEvent, err := repo.Get("e1")
	assert.Error(t, err)
	assert.Nil(t, gotEvent)
}

// TestInMemoryEventRepo_List tests the List method of InMemoryEventRepository
func TestInMemoryEventRepo_List(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event1 := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}
	event2 := &model.Event{
		ID:           "e2",
		Title:        "Conference",
		DurationMin:  120,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(2 * time.Hour)}},
		Participants: []string{"u3", "u4"},
	}

	err := repo.Create(event1)
	assert.NoError(t, err)
	err = repo.Create(event2)
	assert.NoError(t, err)

	gotEvents := repo.List()
	assert.Len(t, gotEvents, 2)
	assert.Contains(t, gotEvents, event1)
	assert.Contains(t, gotEvents, event2)
}

// TestInMemoryEventRepo_DeleteNonExistent tests the Delete method of InMemoryEventRepository for a non-existent event
func TestInMemoryEventRepo_DeleteNonExistent(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	err = repo.Delete("non-existent-id")
	assert.Error(t, err)
}

// TestInMemoryEventRepo_UpdateNonExistent tests the Update method of InMemoryEventRepository for a non-existent event
func TestInMemoryEventRepo_UpdateNonExistent(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	event.ID = "non-existent-id"
	err = repo.Update(event)
	assert.Error(t, err)
}

// TestInMemoryEventRepo_GetNonExistent tests the Get method of InMemoryEventRepository for a non-existent event
func TestInMemoryEventRepo_GetNonExistent(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	gotEvent, err := repo.Get("non-existent-id")
	assert.Error(t, err)
	assert.Nil(t, gotEvent)
}

// TestInMemoryEventRepo_ListEmpty tests the List method of InMemoryEventRepository when no events are present
func TestInMemoryEventRepo_ListEmpty(t *testing.T) {
	repo := NewInMemoryEventRepository()
	gotEvents := repo.List()
	assert.Empty(t, gotEvents)
}

// TestInMemoryEventRepo_AllEventIdsEmpty tests the AllEventIds method of InMemoryEventRepository when no events are present
func TestInMemoryEventRepo_AllEventIdsEmpty(t *testing.T) {
	repo := NewInMemoryEventRepository()
	allEvents, err := repo.AllEventIds()
	assert.NoError(t, err)
	assert.Empty(t, allEvents)
}

// TestInMemoryEventRepo_AllEventIdsNonExistent tests the AllEventIds method of InMemoryEventRepository when no events are present
func TestInMemoryEventRepo_AllEventIdsNonExistent(t *testing.T) {
	repo := NewInMemoryEventRepository()
	event := &model.Event{
		ID:           "e1",
		Title:        "Meeting",
		DurationMin:  60,
		Slots:        []model.Slot{{Start: time.Now(), End: time.Now().Add(1 * time.Hour)}},
		Participants: []string{"u1", "u2"},
	}

	err := repo.Create(event)
	assert.NoError(t, err)

	allEvents, err := repo.AllEventIds()
	assert.NoError(t, err)
	assert.Len(t, allEvents, 1)
}
