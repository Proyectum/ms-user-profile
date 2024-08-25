package usecases

import (
	"fmt"
	"github.com/proyectum/ms-user-profile/internal/adapters/out/persistence"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/errors"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
)

type updateUserProfileUseCaseImpl struct {
	saveRepository   repository.SaveUserProfileRepository
	getRepository    repository.GetUserProfileRepository
	existsRepository repository.ExistsUserProfileRepository
}

func (u *updateUserProfileUseCaseImpl) Update(username string, update entities.UpdateProfile) error {
	exists, err := u.existsRepository.ExistsByUsername(username)

	if err != nil {
		return err
	}

	if !exists {
		return errors.NewUserProfileNotFoundError(fmt.Sprintf("user not found %s", username))
	}

	profile, err := u.getRepository.GetByUsername(username)

	if err != nil {
		return err
	}

	if update.Bio != nil {
		profile.Bio = *update.Bio
	}

	if update.LastName != nil {
		profile.LastName = *update.LastName
	}

	if update.Locale != nil {
		profile.Locale = *update.Locale
	}

	if update.FirstName != nil {
		profile.FirstName = *update.FirstName
	}

	return u.saveRepository.Save(profile)
}

func NewUpdateUserProfileUseCase() usecases.UpdateUserProfileUseCase {
	return &updateUserProfileUseCaseImpl{
		saveRepository:   persistence.NewSaveUserProfileRepository(),
		getRepository:    persistence.NewGetUserProfileRepository(),
		existsRepository: persistence.NewExistsUserProfileRepository(),
	}
}
