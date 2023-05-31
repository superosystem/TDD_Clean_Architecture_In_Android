package constant

import "errors"

// error message conventions
var (
	// ErrAuthenticationFailed error wrong authentication data
	ErrAuthenticationFailed = errors.New("email or password is wrong")

	// ErrAuthenticationFailed error wrong authentication data
	ErrWrongPassword = errors.New("password is wrong")

	// ErrEmailAlreadyExist error email already exist
	ErrEmailAlreadyExist = errors.New("email already used")

	// ErrUserNotFound error user does not exist
	ErrUserNotFound = errors.New("user does not exists")

	// ErrCategoryNotFound error category does not exist
	ErrCategoryNotFound = errors.New("category does not exist")

	// ErrPasswordLengthInvalid error invalid password length
	ErrPasswordLengthInvalid = errors.New("password should have min 6 characters")

	// ErrPasswordNotMatch error both password not match
	ErrPasswordNotMatch = errors.New("password does not match")

	// ErrAccessForbidden error access forbidden
	ErrAccessForbidden = errors.New("you are not have credential to access")

	// ErrUserUnauthorized error user unauthorized
	ErrUserUnauthorized = errors.New("you are not authorized to access")

	// ErrInvalidRequest error invalid request body
	ErrInvalidRequest = errors.New("invalid request body")

	// ErrInvalidJWTPayload error invalid JWT payloads
	ErrInvalidJWTPayload = errors.New("invalid JWT payloads")

	// ErrInvalidTokenHeader error invalid token header
	ErrInvalidTokenHeader = errors.New("invalid token header")

	// ErrRecordNotFound error record not found (cannot specify the error)
	ErrRecordNotFound = errors.New("records is not found")

	// ErrInternalServerError error internal server error
	ErrInternalServerError = errors.New("internal server error")
)
