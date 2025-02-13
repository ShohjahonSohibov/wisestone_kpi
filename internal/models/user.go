package models

import (
	"time"
)

type User struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	FullNameUz string    `bson:"full_name_uz" json:"full_name_uz" binding:"required"`
	FullNameEn string    `bson:"full_name_en" json:"full_name_en" binding:"required"`
	FullNameKr string    `bson:"full_name_kr" json:"full_name_kr" binding:"required"`
	Email      string    `json:"email" binding:"required,email" bson:"email"`
	Password   string    `json:"password,omitempty" binding:"required" bson:"password"`
	RoleId     string    `bson:"role_id" json:"role_id,omitempty"`
	Position   string    `bson:"position" json:"position,omitempty"`
	Username   string    `bson:"username" json:"username,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateUser struct {
	FullNameUz string    `bson:"full_name_uz" json:"full_name_uz" binding:"required"`
	FullNameEn string    `bson:"full_name_en" json:"full_name_en" binding:"required"`
	FullNameKr string    `bson:"full_name_kr" json:"full_name_kr" binding:"required"`
	Email      string    `json:"email" binding:"required,email" bson:"email"`
	Password   string    `json:"password,omitempty" binding:"required" bson:"password"`
	RoleId     string    `bson:"role_id" json:"role_id,omitempty"`
	Position   string    `bson:"position" json:"position,omitempty"`
	Username   string    `bson:"username" json:"username,omitempty"`
}

type UpdateUser struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	FullNameUz string `bson:"full_name_uz" json:"full_name_uz" binding:"required"`
	FullNameEn string `bson:"full_name_en" json:"full_name_en" binding:"required"`
	FullNameKr string `bson:"full_name_kr" json:"full_name_kr" binding:"required"`
	Email      string `json:"email" binding:"required,email" bson:"email"`
	Password   string `json:"password,omitempty" binding:"required" bson:"password"`
	RoleId     string `bson:"role_id" json:"role_id,omitempty"`
	Position   string `bson:"position" json:"position,omitempty"`
	Username   string    `bson:"username" json:"username,omitempty"`
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Items []*User `json:"items"`
}

type ListUsersRequest struct {
	Filter
}

// BeforeCreate sets timestamps before creating a record
func (u *User) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
