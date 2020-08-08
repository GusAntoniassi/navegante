package entity

import (
	"encoding/json"
	"github.com/gusantoniassi/navegante/core/entity/user"
	"time"
)

type User struct {
	ID        user.ID   `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type usr User // prevent recursion
	safeUser := usr(u)
	safeUser.Password = ""
	return json.Marshal(safeUser)
}
