package usecases

import (
	"fmt"
	"github.com/proyectum/ms-user-profile/internal/adapters/out/persistence"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/errors"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
)

type getNotificationUseCaseImpl struct {
	getRepository    repository.GetUserProfileRepository
	existsRepository repository.ExistsUserProfileRepository
}

func (uc *getNotificationUseCaseImpl) GetNotificationSettings(username string) ([]entities.NotificationSetting, error) {
	exists, err := uc.existsRepository.ExistsByUsername(username)

	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.NewUserProfileNotFoundError(fmt.Sprintf("user not found %s", username))
	}

	profile, err := uc.getRepository.GetByUsername(username)

	if err != nil {
		return nil, err
	}

	return profile.NotificationSettings, nil
}

func NewGetNotificationUseCase() usecases.GetNotificationUseCase {
	return &getNotificationUseCaseImpl{
		getRepository:    persistence.NewGetUserProfileRepository(),
		existsRepository: persistence.NewExistsUserProfileRepository(),
	}
}
