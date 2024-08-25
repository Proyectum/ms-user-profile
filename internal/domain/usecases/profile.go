package usecases

import (
	"context"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
)

type GetUserProfileUseCase interface {
	GetUserProfile(ctx context.Context) (*entities.UserProfile, error)
}
