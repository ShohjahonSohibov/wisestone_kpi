package models

type KPIStatus string

const (
	KPIStatusDraft     KPIStatus = "draft"
	KPIStatusPending   KPIStatus = "pending"
	KPIStatusApproved  KPIStatus = "approved"
	KPIStatusRejected  KPIStatus = "rejected"
	KPIStatusCancelled KPIStatus = "cancelled"
)