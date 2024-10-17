package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown       = ErrorType{"unknown"}
	ErrorTypeValidation    = ErrorType{"invalid-input"}
	ErrorTypeNotFound      = ErrorType{"not-found"}
	ErrorTypeConflict      = ErrorType{"conflict"}
	ErrorTypeCommunication = ErrorType{"communication"}
)

type Visitor interface {
	VisitForUnknown(errorType ErrorType)
	VisitForValidation(errorType ErrorType)
	VisitForNotFound(errorType ErrorType)
	VisitForConflict(errorType ErrorType)
	VisitForCommunication(errorType ErrorType)
}

func (et ErrorType) Accept(v Visitor) {
	switch et {
	case ErrorTypeValidation:
		v.VisitForValidation(et)
	case ErrorTypeNotFound:
		v.VisitForNotFound(et)
	case ErrorTypeConflict:
		v.VisitForConflict(et)
	case ErrorTypeCommunication:
		v.VisitForCommunication(et)
	default:
		v.VisitForUnknown(et)
	}
}
