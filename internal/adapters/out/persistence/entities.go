package persistence

import (
	"github.com/google/uuid"
	"time"
)

type UserProfileEntity struct {
	ID                   uuid.UUID                   `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName            string                      `gorm:"column:first_name;type:varchar(50)"`
	LastName             string                      `gorm:"column:last_name;type:varchar(50)"`
	Username             string                      `gorm:"column:username;type:varchar(50);not null"`
	Email                string                      `gorm:"column:email;type:varchar(50);not null"`
	Bio                  string                      `gorm:"column:bio;type:varchar(255)"`
	Locale               string                      `gorm:"column:locale;type:varchar(5)"`
	CreatedAt            time.Time                   `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time                   `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	DeletedAt            *time.Time                  `gorm:"column:deleted_at;type:timestamp with time zone"`
	NotificationSettings []NotificationSettingEntity `gorm:"foreignKey:user_id"`
}

func (*UserProfileEntity) TableName() string {
	return "user_profiles"
}

type NotificationSettingEntity struct {
	ID                 uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	UserProfileID      uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	NotificationTypeID uuid.UUID `gorm:"column:notification_type_id;type:uuid;not null"`
	Active             bool      `gorm:"column:active;type:boolean;default:false"`
	CreatedAt          time.Time `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	DeletedAt          time.Time `gorm:"column:deleted_at;type:timestamp with time zone"`
}

func (*NotificationSettingEntity) TableName() string {
	return "notification_settings"
}

type NotificationTypeEntity struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"column:name;type:varchar(15);not null"`
	Description string    `gorm:"column:description;type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	CreatedBy   string    `gorm:"column:created_by;type:varchar(50);not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	UpdatedBy   string    `gorm:"column:updated_by;type:varchar(50);not null"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:timestamp with time zone"`
	DeletedBy   string    `gorm:"column:deleted_by;type:varchar(50)"`
}

func (*NotificationTypeEntity) TableName() string {
	return "notification_types"
}
