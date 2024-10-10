package errors

import (
	"errors"
)

func IsConflictError(err error) bool {
	return IsCustomErrorOfSpecificType(err, ErrorTypeConflict)
}

func IsUnavailableForLegalReasonsError(err error) bool {
	return IsCustomErrorOfSpecificType(err, ErrorTypeUnavailableForLegalReasons)
}

func IsValidationError(err error) bool {
	return IsCustomErrorOfSpecificType(err, ErrorTypeValidation)
}

func IsUnprocessableEntityError(err error) bool {
	return IsCustomErrorOfSpecificType(err, ErrorTypeUnprocessableEntity)
}

func IsNotFoundError(err error) bool {
	return IsCustomErrorOfSpecificType(err, ErrorTypeNotFound)
}

func IsCustomErrorOfSpecificType(err error, errorType ErrorType) bool {
	ce := CustomError{}
	return errors.As(err, &ce) && ce.ErrorType() == errorType
}
