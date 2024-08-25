package persistence

import (
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"gorm.io/gorm"
)

type getNotificationTypeRepositoryImpl struct {
	ds     *gorm.DB
	mapper UserProfileMapper
}

func (r *getNotificationTypeRepositoryImpl) GetNotificationTypes() ([]entities.NotificationType, error) {
	var types []NotificationTypeEntity

	err := r.ds.Find(&types).Error

	if err != nil {
		return nil, err
	}

	return r.mapper.ToNotificationTypes(types), nil
}

func NewGetNotificationTypeRepository() repository.GetNotificationTypeRepository {
	return &getNotificationTypeRepositoryImpl{
		ds:     getDatasource(),
		mapper: NewUserProfileMapper(),
	}
}
