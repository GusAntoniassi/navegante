package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getMockUsers() []*User {
	return []*User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@doe.com",
			Password:  "hash123456",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			ID:        2,
			Name:      "Jane Doe",
			Email:     "jane@doe.com",
			Password:  "hash123456",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
}

func TestManager_Create(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	user := getMockUsers()[0]

	id, err := mgr.Create(user)
	assert.Nilf(t, err, "Create user should not return any errors")
	assert.Equal(t, user.ID, id)
	assert.False(t, user.CreatedAt.IsZero())
	assert.True(t, user.UpdatedAt.IsZero())
}

func TestManager_Get(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	expectedUser := getMockUsers()[0]
	_, _ = mgr.Create(expectedUser)

	user, err := mgr.Get(expectedUser.ID)
	assert.Nilf(t, err, "Get user should not return any errors")
	assert.Equal(t, user.ID, expectedUser.ID)
	assert.Equal(t, expectedUser, user)
}

func TestManager_GetNotFound(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)

	user, err := mgr.Get(0)
	assert.Nilf(t, user, "Get with nonexistent ID should not return any users")
	assert.Error(t, err)
}

func TestManager_Search(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	expectedUser := getMockUsers()[0]
	_, _ = mgr.Create(expectedUser)

	users, err := mgr.Search("John")
	assert.Nilf(t, err, "Search users should not return any errors")
	assert.Equal(t, 1, len(users), "Search should return at least one user")
	assert.Equal(t, users[0].ID, expectedUser.ID)
}

func TestManager_SearchNotFound(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)

	users, err := mgr.Search("Foobar")
	assert.Nilf(t, users, "Should not return any users")
	assert.Error(t, err)
}

func TestManager_Delete(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	expectedUser := getMockUsers()[0]
	_, _ = mgr.Create(expectedUser)

	err := mgr.Delete(expectedUser.ID)
	assert.Nilf(t, err, "Delete should not return any errors")

	user, err := mgr.Get(expectedUser.ID)
	assert.Nil(t, user)
	assert.Error(t, err, "Should return user not found")
}

func TestManager_DeleteNotFound(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	expectedUser := getMockUsers()[0]

	err := mgr.Delete(expectedUser.ID)
	assert.Errorf(t, err, "Should return user not found")
}

func TestManager_Update(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	expectedUser := getMockUsers()[0]
	_, _ = mgr.Create(expectedUser)

	expectedUser.Name = "Gus"
	err := mgr.Update(expectedUser)
	assert.Nilf(t, err, "Update user should not return any errors")

	user, err := mgr.Get(expectedUser.ID)
	assert.Nilf(t, err, "Should find the updated user")
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestManager_UpdateNotFound(t *testing.T) {
	repo := NewInMemRepository()
	mgr := NewManager(repo)
	userNotInDatabase := getMockUsers()[0]

	err := mgr.Update(userNotInDatabase)
	assert.NotNilf(t, err, "Should return user not found error")
}
