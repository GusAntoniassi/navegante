package user

import (
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
	return mgr.repo.Create(user)
}

func (mgr *manager) Get(id ID) (*User, error) {
	return mgr.repo.Get(id)
}

func (mgr *manager) Search(query string) ([]*User, error) {
	return mgr.repo.Search(strings.ToLower(query))
}

func (mgr *manager) Delete(id ID) error {
	return mgr.repo.Delete(id)
}

func (mgr *manager) Update(user *User) error {
	user.UpdatedAt = time.Now()
	return mgr.repo.Update(user)
}
