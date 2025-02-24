package models

import (
	"time"
)

type KPIProgress struct {
	ID                string       `json:"id,omitempty" bson:"_id,omitempty"`
	Kpi               KPIPrgParent `bson:"kpi,omitempty" json:"kpi,omitempty"`
	FactorId          string       `bson:"factor_id,omitempty" json:"factor_id,omitempty"`
	FactorIndicatorId string       `bson:"factor_indicator_id,omitempty" json:"factor_indicator_id,omitempty"`
	Ratio             int32        `bson:"ratio,omitempty" json:"ratio,omitempty"`
	TeamId            string       `bson:"team_id,omitempty" json:"team_id,omitempty"`
	EmployeeId        string       `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	CreatedById       string       `bson:"created_by_id,omitempty" json:"created_by_id,omitempty"`
	Date              string       `bson:"date,omitempty" json:"date,omitempty"`
	CreatedAt         time.Time    `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt         time.Time    `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPIProgress struct {
	FactorId          string `bson:"factor_id,omitempty" json:"factor_id,omitempty"`
	FactorIndicatorId string `bson:"factor_indicator_id,omitempty" json:"factor_indicator_id,omitempty"`
	TeamId            string `bson:"team_id,omitempty" json:"team_id,omitempty"`
	EmployeeId        string `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	CreatedById       string `bson:"created_by_id,omitempty" json:"created_by_id,omitempty"`
	Ratio             int32  `bson:"ratio,omitempty" json:"ratio,omitempty"`
	Date              string `bson:"date,omitempty" json:"date,omitempty"`
}

type ListKPIProgressResponse struct {
	Progress *KPIProgress `json:"progress,omitempty"`
}

type ListKPIProgressRequest struct {
	Date       string `json:"date"`
	EmployeeId string `json:"employee_id"`
	TeamId     string `json:"team_id"`
}

type KPIProgressTeamFilter struct {
	Date   string `json:"date"`
	TeamId string `json:"team_id"`
}

type KPIProgressEmployeeFilter struct {
	Date       string `json:"date"`
	EmployeeId string `json:"employee_id"`
}

type KPIPrgParent struct {
	ID            string                `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string                `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string                `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string                `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string                `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Year          string                `bson:"year,omitempty" json:"year,omitempty"`
	Status        string                `bson:"status,omitempty" json:"status,omitempty"`
	Type          string                `bson:"type,omitempty" json:"type,omitempty"`
	TotalRatio    float64               `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	Divisions     []ShortKPIPrgDivision `bson:"divisions,omitempty" json:"divisions,omitempty"`
}

type ShortKPIPrgDivision struct {
	ID            string                 `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string                 `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string                 `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string                 `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string                 `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	TotalRatio    float64                `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	Criterions    []ShortKPIPrgCriterion `bson:"criterions,omitempty" json:"criterions,omitempty"`
}

type ShortKPIPrgCriterion struct {
	ID            string               `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string               `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string               `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string               `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string               `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	TotalRatio    float64              `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	Factors       []*ShortKPIPrgFactor `bson:"factors,omitempty" json:"factors,omitempty"`
}

type ShortKPIPrgFactor struct {
	ID               string                        `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn           string                        `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr           string                        `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	Ratio            float64                       `bson:"ratio,omitempty" json:"ratio,omitempty"`
	DescriptionEn    string                        `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr    string                        `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	TotalRatio       float64                       `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	FactorIndicators []*ShortKPIPrgFactorIndicator `bson:"factor_indicators,omitempty" json:"factor_indicators,omitempty"`
}

type ShortKPIPrgFactorIndicator struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	ProgressRange int32  `bson:"progress_range,omitempty" json:"progress_range,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Ratio         int32  `bson:"ratio,omitempty" json:"ratio,omitempty"`
}

func (k *KPIProgress) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}
