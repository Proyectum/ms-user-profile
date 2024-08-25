package usecases

import (
	"github.com/google/uuid"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
)

type UpdateNotificationUseCase interface {
	Update(username string, typeID uuid.UUID, active bool) error
}

type GetNotificationUseCase interface {
	GetNotificationSettings(username string) ([]entities.NotificationSetting, error)
}

type GetNotificationTypeUseCase interface {
	GetNotificationTypes() ([]entities.NotificationType, error)
}
