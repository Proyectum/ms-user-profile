package entities

import (
	"github.com/google/uuid"
	"time"
)

type UserProfile struct {
	ID                   uuid.UUID
	FirstName            string
	LastName             string
	Username             string
	Email                string
	Bio                  string
	Locale               string
	Initials             string
	NotificationSettings []NotificationSetting
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

type UpdateProfile struct {
	Bio       *string
	FirstName *string
	LastName  *string
	Locale    *string
}
