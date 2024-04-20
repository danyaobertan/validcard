package errorops

type Error struct { //nolint:govet // benefits readability
	Code        int      `json:"code,omitempty"`
	Message     string   `json:"message,omitempty"`
	Description any      `json:"description,omitempty"`
	Value       []string `json:"value,omitempty"`
}

func NewError(code int, message string, description any, value ...string) *Error {
	return &Error{
		Code:        code,
		Message:     message,
		Description: description,
		Value:       value,
	}
}

func (e *Error) Error() string {
	return e.Message
}
