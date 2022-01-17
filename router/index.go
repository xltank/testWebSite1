package router

import (
	"time"
)

type Model struct {
	ID        int       `json:"ID,omitempty"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
