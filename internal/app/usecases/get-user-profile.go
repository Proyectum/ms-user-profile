package usecases

import (
	"context"
	"github.com/google/uuid"
	"github.com/proyectum/ms-user-profile/internal/adapters/out/persistence"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
	"strings"
	"time"
)

type getUserProfileUseCaseImpl struct {
	existsRepository repository.ExistsUserProfileRepository
	getRepository    repository.GetUserProfileRepository
	saveRepository   repository.SaveUserProfileRepository
}

func (g *getUserProfileUseCaseImpl) GetUserProfile(ctx context.Context) (*entities.UserProfile, error) {
	username := ctx.Value("username").(string)
	email := ctx.Value("email").(string)
	exists, err := g.existsRepository.ExistsByUsername(username)

	if err != nil {
		return nil, err
	}

	if !exists {
		return g.createProfile(username, email)
	}

	profile, err := g.getRepository.GetByUsername(username)

	if err != nil {
		return nil, err
	}

	profile.Initials = g.calculateInitials(profile.FirstName, profile.LastName)

	return profile, nil
}

func (g *getUserProfileUseCaseImpl) calculateInitials(firstName, lastName string) string {
	firstNames := strings.Fields(firstName)
	lastNames := strings.Fields(lastName)

	if len(firstNames) == 0 || len(lastNames) == 0 {
		return ""
	}

	if len(firstNames) > 1 {
		firstName = firstNames[0]
	}

	if len(lastNames) > 1 {
		lastName = lastNames[0]
	}

	f := strings.ToUpper(string(firstName[0]))
	s := strings.ToUpper(string(lastName[0]))
	return strings.Join([]string{f, s}, "")
}

func (g *getUserProfileUseCaseImpl) createProfile(username string, email string) (*entities.UserProfile, error) {
	profile := &entities.UserProfile{
		ID:        uuid.New(),
		Email:     email,
		Username:  username,
		Initials:  g.calculateInitials("", ""),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := g.saveRepository.Save(profile)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func NewUserProfileUseCase() usecases.GetUserProfileUseCase {
	return &getUserProfileUseCaseImpl{
		existsRepository: persistence.NewExistsUserProfileRepository(),
		getRepository:    persistence.NewGetUserProfileRepository(),
		saveRepository:   persistence.NewSaveUserProfileRepository(),
	}
}
