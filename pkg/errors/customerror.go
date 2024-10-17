package errors

import (
	"runtime/debug"
)

type CustomError struct {
	code       string
	message    string
	errorType  ErrorType
	stackTrace []byte
}

func NewCustomError(code string, message string, errorType ErrorType) CustomError {
	return CustomError{
		code:       code,
		message:    message,
		errorType:  errorType,
		stackTrace: debug.Stack(),
	}
}

func (s CustomError) Code() string {
	return s.code
}

func (s CustomError) Error() string {
	return s.message
}

func (s CustomError) ErrorType() ErrorType {
	return s.errorType
}

func (s CustomError) StackTrace() []byte {
	return s.stackTrace
}

func NewUnknownError(code string, message string) CustomError {
	return NewCustomError(code, message, ErrorTypeUnknown)
}

func NewValidationError(code string, message string) CustomError {
	return NewCustomError(code, message, ErrorTypeValidation)
}

func NewNotFoundError(message string) CustomError {
	return NewCustomError("not_found", message, ErrorTypeNotFound)
}

func NewConflictError(message string) CustomError {
	return NewCustomError("conflict", message, ErrorTypeConflict)
}

func NewCommunicationError(message string) CustomError {
	return NewCustomError("communication_error", message, ErrorTypeCommunication)
}
