package usecases

import (
	"github.com/proyectum/ms-user-profile/internal/adapters/out/persistence"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/repository"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
)

type getNotificationTypeUseCaseImpl struct {
	getNotificationTypeRepository repository.GetNotificationTypeRepository
}

func (uc *getNotificationTypeUseCaseImpl) GetNotificationTypes() ([]entities.NotificationType, error) {
	return uc.getNotificationTypeRepository.GetNotificationTypes()
}

func NewGetNotificationTypeUseCase() usecases.GetNotificationTypeUseCase {
	return &getNotificationTypeUseCaseImpl{
		getNotificationTypeRepository: persistence.NewGetNotificationTypeRepository(),
	}
}
