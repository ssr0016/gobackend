package errors

const (
	GeneralErrorCode = "E000"
)

type ErrorStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func New(code, message string) ErrorStatus {
	return ErrorStatus{
		Code:    code,
		Message: message,
	}
}

func (e ErrorStatus) Error() string {
	return e.Message
}

func (s ErrorStatus) MessageError() string {
	return s.Message
}
