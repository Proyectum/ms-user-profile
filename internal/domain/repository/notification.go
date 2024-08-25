package repository

import "github.com/proyectum/ms-user-profile/internal/domain/entities"

type GetNotificationTypeRepository interface {
	GetNotificationTypes() ([]entities.NotificationType, error)
}
