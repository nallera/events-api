package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown                    = ErrorType{"unknown"}
	ErrorTypeValidation                 = ErrorType{"invalid-input"}
	ErrorTypeUnprocessableEntity        = ErrorType{"unprocessable-entity"}
	ErrorTypeNotFound                   = ErrorType{"not-found"}
	ErrorTypeRequestTimeout             = ErrorType{"request-timeout"}
	ErrorTypeRateLimit                  = ErrorType{"rate-limit-reached"}
	ErrorTypeConflict                   = ErrorType{"conflict"}
	ErrorTypeUnavailableForLegalReasons = ErrorType{"unavailable-for-legal-reasons"}
	ErrorTypeCommunication              = ErrorType{"communication"}
)

type Visitor interface {
	VisitForUnknown(errorType ErrorType)
	VisitForValidation(errorType ErrorType)
	VisitForUnprocessableEntity(errorType ErrorType)
	VisitForNotFound(errorType ErrorType)
	VisitForRequestTimeout(errorType ErrorType)
	VisitForRateLimit(errorType ErrorType)
	VisitForConflict(errorType ErrorType)
	VisitForUnavailableForLegalReasons(errorType ErrorType)
	VisitForCommunication(errorType ErrorType)
}

func (et ErrorType) Accept(v Visitor) {
	switch et {
	case ErrorTypeValidation:
		v.VisitForValidation(et)
	case ErrorTypeUnprocessableEntity:
		v.VisitForUnprocessableEntity(et)
	case ErrorTypeNotFound:
		v.VisitForNotFound(et)
	case ErrorTypeRequestTimeout:
		v.VisitForRequestTimeout(et)
	case ErrorTypeRateLimit:
		v.VisitForRateLimit(et)
	case ErrorTypeConflict:
		v.VisitForConflict(et)
	case ErrorTypeUnavailableForLegalReasons:
		v.VisitForUnavailableForLegalReasons(et)
	case ErrorTypeCommunication:
		v.VisitForCommunication(et)
	default:
		v.VisitForUnknown(et)
	}
}
