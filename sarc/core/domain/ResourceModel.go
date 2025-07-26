package domain

import (
	"github.com/lib/pq"
)

// Resource represents a resource in the system.
// swagger:model
type Resource struct {
	ResourceID  uint           `gorm:"primaryKey" json:"resourceId"`
	Description string         `json:"description"`
	Status      ResourceStatus `json:"status"`
	// Characteristics is an array of strings stored as Postgres text[].
	// For Swagger, treat as []string.
	Characteristics pq.StringArray `gorm:"type:text[]" json:"characteristics" swaggertype:"array,string"`
	ResourceTypeID  uint           `json:"resourceTypeId"`
}

// ResourceType represents the type of a resource.
// swagger:model
type ResourceType struct {
	ResourceTypeID uint   `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
}

// ResourceStatus represents the status of a resource.
// swagger:model
type ResourceStatus string

const (
	ResourceStatusAvailable   ResourceStatus = "available"
	ResourceStatusUnavailable ResourceStatus = "unavailable"
	ResourceStatusReserved    ResourceStatus = "reserved"
)
