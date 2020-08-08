package user

type Reader interface {
	Get(id ID) (*User, error)
	Search(query string) ([]*User, error)
	List() ([]*User, error)
}

type Writer interface {
	Create(e *User) (ID, error)
	Update(e *User) error
	Delete(id ID) error
}

type repository interface {
	Reader
	Writer
}

type Manager interface {
	repository
}
