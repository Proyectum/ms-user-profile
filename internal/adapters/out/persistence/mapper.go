package persistence

import "github.com/proyectum/ms-user-profile/internal/domain/entities"

type UserProfileMapper interface {
	ToUserProfile(*UserProfileEntity) *entities.UserProfile
	ToUserProfileEntity(*entities.UserProfile) *UserProfileEntity
	ToNotificationSetting(*NotificationSettingEntity) *entities.NotificationSetting
	ToNotificationSettingEntity(*entities.NotificationSetting) *NotificationSettingEntity
	ToNotificationSettings([]NotificationSettingEntity) []entities.NotificationSetting
	ToNotificationSettingEntities([]entities.NotificationSetting) []NotificationSettingEntity
	ToNotificationTypeEntity(*entities.NotificationType) *NotificationTypeEntity
	ToNotificationType(*NotificationTypeEntity) *entities.NotificationType
	ToNotificationTypes([]NotificationTypeEntity) []entities.NotificationType
}

type userProfileMapperImpl struct {
}

func (u *userProfileMapperImpl) ToNotificationTypes(src []NotificationTypeEntity) []entities.NotificationType {
	if src == nil {
		return nil
	}

	result := make([]entities.NotificationType, len(src))

	for i, _ := range src {
		result[i] = *u.ToNotificationType(&src[i])
	}

	return result
}

func (u *userProfileMapperImpl) ToNotificationSettings(src []NotificationSettingEntity) []entities.NotificationSetting {
	if src == nil {
		return nil
	}

	result := make([]entities.NotificationSetting, len(src))

	for i, _ := range src {
		result[i] = *u.ToNotificationSetting(&src[i])
	}

	return result
}

func (u *userProfileMapperImpl) ToNotificationSettingEntities(src []entities.NotificationSetting) []NotificationSettingEntity {
	if src == nil {
		return nil
	}

	result := make([]NotificationSettingEntity, len(src))

	for i, _ := range src {
		result[i] = *u.ToNotificationSettingEntity(&src[i])
	}

	return result
}

func (u *userProfileMapperImpl) ToUserProfile(src *UserProfileEntity) *entities.UserProfile {
	if src == nil {
		return nil
	}

	return &entities.UserProfile{
		ID:                   src.ID,
		FirstName:            src.FirstName,
		LastName:             src.LastName,
		Username:             src.Username,
		Email:                src.Email,
		Bio:                  src.Bio,
		Locale:               src.Locale,
		NotificationSettings: u.ToNotificationSettings(src.NotificationSettings),
		CreatedAt:            src.CreatedAt,
		UpdatedAt:            src.UpdatedAt,
		DeletedAt:            src.DeletedAt,
	}
}

func (u *userProfileMapperImpl) ToUserProfileEntity(src *entities.UserProfile) *UserProfileEntity {
	if src == nil {
		return nil
	}

	return &UserProfileEntity{
		ID:                   src.ID,
		FirstName:            src.FirstName,
		LastName:             src.LastName,
		Username:             src.Username,
		Email:                src.Email,
		Bio:                  src.Bio,
		Locale:               src.Locale,
		NotificationSettings: u.ToNotificationSettingEntities(src.NotificationSettings),
		CreatedAt:            src.CreatedAt,
		UpdatedAt:            src.UpdatedAt,
		DeletedAt:            src.DeletedAt,
	}
}

func (u *userProfileMapperImpl) ToNotificationSetting(src *NotificationSettingEntity) *entities.NotificationSetting {
	if src == nil {
		return nil
	}

	return &entities.NotificationSetting{
		ID:                 src.ID,
		UserProfileID:      src.UserProfileID,
		NotificationTypeID: src.NotificationTypeID,
		Active:             src.Active,
		CreatedAt:          src.CreatedAt,
		UpdatedAt:          src.UpdatedAt,
		DeletedAt:          src.DeletedAt,
	}
}

func (u *userProfileMapperImpl) ToNotificationSettingEntity(src *entities.NotificationSetting) *NotificationSettingEntity {
	if src == nil {
		return nil
	}

	return &NotificationSettingEntity{
		ID:                 src.ID,
		UserProfileID:      src.UserProfileID,
		NotificationTypeID: src.NotificationTypeID,
		Active:             src.Active,
		CreatedAt:          src.CreatedAt,
		UpdatedAt:          src.UpdatedAt,
		DeletedAt:          src.DeletedAt,
	}
}

func (u *userProfileMapperImpl) ToNotificationTypeEntity(src *entities.NotificationType) *NotificationTypeEntity {
	if src == nil {
		return nil
	}

	return &NotificationTypeEntity{
		ID:          src.ID,
		Name:        src.Name,
		Description: src.Description,
		CreatedAt:   src.CreatedAt,
		CreatedBy:   src.CreatedBy,
		UpdatedAt:   src.UpdatedAt,
		UpdatedBy:   src.UpdatedBy,
		DeletedAt:   src.DeletedAt,
		DeletedBy:   src.DeletedBy,
	}
}

func (u *userProfileMapperImpl) ToNotificationType(src *NotificationTypeEntity) *entities.NotificationType {
	if src == nil {
		return nil
	}

	return &entities.NotificationType{
		ID:          src.ID,
		Name:        src.Name,
		Description: src.Description,
		CreatedAt:   src.CreatedAt,
		CreatedBy:   src.CreatedBy,
		UpdatedAt:   src.UpdatedAt,
		UpdatedBy:   src.UpdatedBy,
		DeletedAt:   src.DeletedAt,
		DeletedBy:   src.DeletedBy,
	}
}

func NewUserProfileMapper() UserProfileMapper {
	return &userProfileMapperImpl{}
}
