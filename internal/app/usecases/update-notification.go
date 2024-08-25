package usecases

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/proyectum/ms-user-profile/internal/adapters/out/persistence"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/errors"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
	"time"
)

type updateNotificationUseCaseImpl struct {
	saveRepository   repository.SaveUserProfileRepository
	getRepository    repository.GetUserProfileRepository
	existsRepository repository.ExistsUserProfileRepository
}

func (uc *updateNotificationUseCaseImpl) Update(username string, typeID uuid.UUID, active bool) error {
	exists, err := uc.existsRepository.ExistsByUsername(username)

	if err != nil {
		return err
	}

	if !exists {
		return errors.NewUserProfileNotFoundError(fmt.Sprintf("user not found %s", username))
	}

	profile, err := uc.getRepository.GetByUsername(username)

	if err != nil {
		return err
	}

	existsNotification := false
	settings := profile.NotificationSettings

	for i, _ := range settings {
		if settings[i].NotificationTypeID == typeID {
			settings[i].Active = active
			settings[i].UpdatedAt = time.Now()
			existsNotification = true
			break
		}
	}

	if existsNotification {
		return uc.saveRepository.Save(profile)
	}

	newSetting := entities.NotificationSetting{
		ID:                 uuid.New(),
		UserProfileID:      profile.ID,
		NotificationTypeID: typeID,
		Active:             active,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		DeletedAt:          nil,
	}

	profile.NotificationSettings = append(profile.NotificationSettings, newSetting)
	return uc.saveRepository.Save(profile)
}

func NewUpdateNotificationUseCase() usecases.UpdateNotificationUseCase {
	return &updateNotificationUseCaseImpl{
		saveRepository:   persistence.NewSaveUserProfileRepository(),
		getRepository:    persistence.NewGetUserProfileRepository(),
		existsRepository: persistence.NewExistsUserProfileRepository(),
	}
}
