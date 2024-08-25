package persistence

import (
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"gorm.io/gorm"
)

type getUserProfileRepositoryImpl struct {
	ds     *gorm.DB
	mapper UserProfileMapper
}

func (g *getUserProfileRepositoryImpl) GetByUsername(username string) (*entities.UserProfile, error) {
	var entity UserProfileEntity
	err := g.ds.Preload("NotificationSettings").
		First(&entity, "username = ? and deleted_at is null", username).
		Error

	if err != nil {
		return nil, err
	}

	return g.mapper.ToUserProfile(&entity), nil
}

func NewGetUserProfileRepository() repository.GetUserProfileRepository {
	return &getUserProfileRepositoryImpl{
		ds:     getDatasource(),
		mapper: NewUserProfileMapper(),
	}
}
