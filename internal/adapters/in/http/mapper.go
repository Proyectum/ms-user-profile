package http

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/api"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
)

type UserProfileApiMapper interface {
	ToNotificationSettingDTOs([]entities.NotificationSetting) api.NotificationSettings
	ToNotificationSettingDTO(*entities.NotificationSetting) *api.NotificationSetting
	ToGeneralProfile(*entities.UserProfile) *api.GeneralProfile
	ToUpdateProfileDomain(*api.UpdateProfile) *entities.UpdateProfile
}

type userProfileApiMapperImpl struct {
}

func (m *userProfileApiMapperImpl) ToUpdateProfileDomain(src *api.UpdateProfile) *entities.UpdateProfile {
	if src == nil {
		return nil
	}

	return &entities.UpdateProfile{
		LastName:  src.LastName,
		FirstName: src.FirstName,
		Locale:    src.Locale,
		Bio:       src.Bio,
	}
}

func (m *userProfileApiMapperImpl) ToNotificationSettingDTO(src *entities.NotificationSetting) *api.NotificationSetting {
	if src == nil {
		return nil
	}

	return &api.NotificationSetting{
		TypeId: src.NotificationTypeID,
		Active: src.Active,
	}
}

func (m *userProfileApiMapperImpl) ToGeneralProfile(src *entities.UserProfile) *api.GeneralProfile {
	if src == nil {
		return nil
	}

	return &api.GeneralProfile{
		Bio:       &src.Bio,
		Email:     openapi_types.Email(src.Email),
		FirstName: src.FirstName,
		Initials:  src.Initials,
		LastName:  src.LastName,
		Locale:    src.Locale,
		Username:  src.Username,
	}
}

func (m *userProfileApiMapperImpl) ToNotificationSettingDTOs(src []entities.NotificationSetting) api.NotificationSettings {
	if src == nil {
		return nil
	}

	results := make([]api.NotificationSetting, len(src))

	for i, _ := range src {
		results[i] = *m.ToNotificationSettingDTO(&src[i])
	}

	return results
}

func NewUserProfileApiMapper() UserProfileApiMapper {
	return &userProfileApiMapperImpl{}
}
