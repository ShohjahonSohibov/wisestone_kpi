package models

import (
	"time"
)

type KPIProgressStatus struct {
	ID         string     `json:"id,omitempty" bson:"_id,omitempty"`
	Team       *TeamShort `bson:"team,omitempty" json:"team,omitempty"`
	TeamId     string     `bson:"team_id,omitempty" json:"team_id,omitempty"`
	Employee   *ShortUser `bson:"employee,omitempty" json:"employee,omitempty"`
	EmployeeId string     `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	Date       string     `bson:"date,omitempty" json:"date,omitempty"`
	Status     string     `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPIProgressStatus struct {
	TeamId     string `bson:"team_id,omitempty" json:"team_id,omitempty"`
	EmployeeId string `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	Date       string `bson:"date,omitempty" json:"date,omitempty"`
	Status     string `bson:"status,omitempty" json:"status,omitempty"`
}

type UpdateKPIProgressStatus struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
}

type ListKPIProgressStatusResponse struct {
	Count int                  `json:"count,omitempty"`
	Items []*KPIProgressStatus `json:"items,omitempty"`
}

type ListKPIProgressStatusRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Page   int    `json:"page"`
	Date   string `json:"date"`
	Type   string `json:"type"`
	Status string `son:"status"`
}

func (k *KPIProgressStatus) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}

// create team status for kpi progress, approved, pending, rejected.
