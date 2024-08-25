package entities

import (
	"github.com/google/uuid"
	"time"
)

type NotificationType struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
	DeletedAt   *time.Time
	DeletedBy   *string
}

type NotificationSetting struct {
	ID                 uuid.UUID
	UserProfileID      uuid.UUID
	NotificationTypeID uuid.UUID
	Active             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
