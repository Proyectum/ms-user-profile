package repository

import "github.com/proyectum/ms-user-profile/internal/domain/entities"

type ExistsUserProfileRepository interface {
	ExistsByUsername(username string) (bool, error)
}

type SaveUserProfileRepository interface {
	Save(profile *entities.UserProfile) error
}

type GetUserProfileRepository interface {
	GetByUsername(username string) (*entities.UserProfile, error)
}
