package user

import (
	"fmt"
	"strings"
)

type InMemoryRepo struct {
	users map[ID]*User
}

func NewInMemRepository() *InMemoryRepo {
	db := map[ID]*User{}

	return &InMemoryRepo{
		users: db,
	}
}

func (r *InMemoryRepo) Create(user *User) (ID, error) {
	r.users[user.ID] = user
	return user.ID, nil
}

func (r *InMemoryRepo) Get(id ID) (*User, error) {
	if r.users[id] == nil {
		return nil, fmt.Errorf("user not found") // @TODO: Implement not found error for the entire domain
	}

	return r.users[id], nil
}

func (r *InMemoryRepo) Update(user *User) error {
	_, err := r.Get(user.ID)

	if err != nil {
		return err
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryRepo) Search(query string) ([]*User, error) {
	var matches []*User
	for _, user := range r.users {
		if strings.Contains(strings.ToLower(user.Name), query) || strings.Contains(strings.ToLower(user.Email), query) {
			matches = append(matches, user)
		}
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("users not found")
	}

	return matches, nil
}

func (r *InMemoryRepo) List() ([]*User, error) {
	var users []*User
	for _, j := range r.users {
		users = append(users, j)
	}

	return users, nil
}

func (r *InMemoryRepo) Delete(id ID) error {
	if r.users[id] == nil {
		return fmt.Errorf("user not found")
	}

	r.users[id] = nil
	return nil
}
