package errors

type UserProfileNotFoundError struct {
	message string
}

func (i UserProfileNotFoundError) Error() string {
	return i.message
}

func NewUserProfileNotFoundError(message string) error {
	return &UserProfileNotFoundError{
		message: message,
	}
}
