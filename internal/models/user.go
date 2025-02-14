package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	FullNameEn   string    `bson:"full_name_en" json:"full_name_en"`
	FullNameKr   string    `bson:"full_name_kr" json:"full_name_kr"`
	Email        string    `json:"email" bson:"email"`
	Password     string    `json:"password,omitempty" binding:"required" bson:"password"`
	RoleId       string    `bson:"role_id" json:"role_id,omitempty"`
	Role         ShortRole `bson:"role" json:"role,omitempty"`
	TeamId       string    `bson:"team_id" json:"team_id,omitempty"`
	Team         TeamShort `bson:"team" json:"team,omitempty"`
	Position     string    `bson:"position" json:"position,omitempty"`
	IsTeamLeader bool      `bson:"is_team_leader" json:"is_team_leader,omitempty"`
	Username     string    `bson:"username" json:"username,omitempty" binding:"required"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateUser struct {
	FullNameEn string `bson:"full_name_en" json:"full_name_en"`
	FullNameKr string `bson:"full_name_kr" json:"full_name_kr"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password,omitempty" binding:"required" bson:"password"`
	Username   string `bson:"username" json:"username,omitempty" binding:"required"`
	RoleId     string `bson:"role_id" json:"role_id,omitempty"`
	Position   string `bson:"position" json:"position,omitempty"`
}

type UpdateUser struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	FullNameEn string `bson:"full_name_en" json:"full_name_en"`
	FullNameKr string `bson:"full_name_kr" json:"full_name_kr"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password,omitempty" bson:"password"`
	RoleId     string `bson:"role_id" json:"role_id,omitempty"`
	Position   string `bson:"position" json:"position,omitempty"`
	Username   string `bson:"username" json:"username,omitempty"`
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Items []*User `json:"items"`
}

type ListUsersRequest struct {
	TeamId string `json:"team_id"`
	Filter
}

// BeforeCreate sets timestamps before creating a record
func (u *User) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
