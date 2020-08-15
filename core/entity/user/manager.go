package user

import (
	"fmt"
	"strings"
	"time"
)

type manager struct {
	repo repository
}

func NewManager(r repository) *manager {
	return &manager{
		repo: r,
	}
}

func (mgr *manager) Create(user *User) (ID, error) {
	user.CreatedAt = time.Now()
	hashedPassword, err := Hash(user.Password)

	if err != nil {
		return 0, err
	}

	user.Password = hashedPassword
	return mgr.repo.Create(user)
}

func (mgr *manager) Get(id ID) (*User, error) {
	return mgr.repo.Get(id)
}

func (mgr *manager) List() ([]*User, error) {
	return mgr.repo.List()
}

func (mgr *manager) Search(query string) ([]*User, error) {
	return mgr.repo.Search(strings.ToLower(query))
}

func (mgr *manager) Delete(id ID) error {
	return mgr.repo.Delete(id)
}

func (mgr *manager) Update(user *User) error {
	user.UpdatedAt = time.Now()

	oldUser, err := mgr.Get(user.ID)
	if err != nil {
		return fmt.Errorf("error getting the previous user configuration: %s", err)
	}

	// disallow updating e-mail
	user.Email = oldUser.Email
	user.CreatedAt = oldUser.CreatedAt

	return mgr.repo.Update(user)
}
