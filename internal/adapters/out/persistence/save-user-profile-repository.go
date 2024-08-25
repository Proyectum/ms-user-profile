package persistence

import (
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"gorm.io/gorm"
)

type saveUserProfileRepositoryImpl struct {
	ds     *gorm.DB
	mapper UserProfileMapper
}

func (s *saveUserProfileRepositoryImpl) Save(profile *entities.UserProfile) error {
	entity := s.mapper.ToUserProfileEntity(profile)
	return s.ds.Save(&entity).Error
}

func NewSaveUserProfileRepository() repository.SaveUserProfileRepository {
	return &saveUserProfileRepositoryImpl{
		ds:     getDatasource(),
		mapper: NewUserProfileMapper(),
	}
}
