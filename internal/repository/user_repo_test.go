package repository_test

import (
	"meeting-scheduler/internal/model"
	"meeting-scheduler/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryUserRepo_CreateGet(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	user := &model.User{ID: "u1", Name: "Alice"}
	err := repo.Create(user)
	assert.NoError(t, err)

	gotUser, err := repo.Get("u1")
	assert.NoError(t, err)
	assert.Equal(t, user, gotUser)
}

func TestInMemoryUserRepo_GetAll(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	user1 := &model.User{ID: "u1", Name: "Alice"}
	user2 := &model.User{ID: "u2", Name: "Bob"}

	_ = repo.Create(user1)
	_ = repo.Create(user2)

	allUsers, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, allUsers, 2)
	assert.Contains(t, allUsers, "u1")
	assert.Contains(t, allUsers, "u2")
}
