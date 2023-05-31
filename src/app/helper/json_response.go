package helper

import "net/http"

type BaseResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseMessageResponse struct {
	Message string `json:"message"`
}

// SuccessResponse format on response success
func SuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// SuccessCreatedResponse format on response success created
func SuccessCreatedResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: message,
	}
}

// BadRequestResponse format on response error bad request
func BadRequestResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusBadRequest,
		Status:  "error",
		Message: message,
	}
}

// NotFoundResponse format on response error not found
func NotFoundResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusNotFound,
		Status:  "error",
		Message: message,
	}
}

// UnauthorizedResponse format on response error unauthorized
func UnauthorizedResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusUnauthorized,
		Status:  "error",
		Message: message,
	}
}

// ForbiddenResponse format on response error forbidden
func ForbiddenResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusForbidden,
		Status:  "error",
		Message: message,
	}
}

// ConflictResponse format on response error conflict
func ConflictResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusConflict,
		Status:  "error",
		Message: message,
	}
}

// InternalServerErrorResponse format on response internal server error
func InternalServerErrorResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusInternalServerError,
		Status:  "error",
		Message: message,
	}
}

// MessageResponse format on response error
func MessageErrorResponse(message string) BaseMessageResponse {
	return BaseMessageResponse{
		Message: message,
	}
}
