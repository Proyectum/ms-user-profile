package persistence

import (
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"gorm.io/gorm"
)

type existsUserProfileRepositoryImpl struct {
	ds *gorm.DB
}

func (e *existsUserProfileRepositoryImpl) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := e.ds.Model(&UserProfileEntity{}).
		Where("username = ?", username).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func NewExistsUserProfileRepository() repository.ExistsUserProfileRepository {
	return &existsUserProfileRepositoryImpl{
		ds: getDatasource(),
	}
}
