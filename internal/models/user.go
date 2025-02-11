package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	FullName  string    `bson:"full_name" json:"full_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email" bson:"email"`
	Password  string    `json:"password,omitempty" binding:"required" bson:"password"`
	RoleId    string    `bson:"role_id" json:"role_id,omitempty"`
	Position  string    `bson:"position" json:"position,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// BeforeCreate sets timestamps before creating a record
func (u *User) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

// BeforeUpdate sets the updated timestamp before updating a record
func (u *User) BeforeUpdate() {
	u.UpdatedAt = time.Now()
}

// MarshalJSON customizes the JSON output of timestamps
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias:     Alias(u),
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	})
}

type CreateUser struct {
	FullName  string    `bson:"full_name" json:"full_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email" bson:"email"`
	Password  string    `json:"password,omitempty" binding:"required" bson:"password"`
	RoleId    string    `bson:"role_id" json:"role_id,omitempty"`
	Position  string    `bson:"position" json:"position,omitempty"`
}

type ListUsersResponse struct {
	Users []*User `json:"users"`
	Count int     `json:"count"`
}

type ListUsersRequest struct {
	Filter Filter `json:"filter"`
}
