package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Surname  string    `json:"surname,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	RoleId   uuid.UUID `json:"role_id,omitempty"`
}

type InsertUserRequest struct {
	Name     string    `json:"name,omitempty"`
	Surname  string    `json:"surname,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	RoleName string    `json:"role_name,omitempty"`
}
