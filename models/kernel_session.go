package models

import (
	"gorm.io/datatypes"
	"time"
)

type KernelSession struct {
	ID        string            `json:"kernel_id" gorm:"primaryKey"`
	Session   datatypes.JSONMap `json:"kernel_session" gorm:"type:json"`
	CreatedAt time.Time         `gorm:"autoCreateTime"`
	UpdateAt  time.Time         `gorm:"autoUpdateTime"`
}
