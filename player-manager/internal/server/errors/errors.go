package errors

import "encoding/json"

func NewError(s int, m string, e error) *CommonError {
	return &CommonError{s, m, e}
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int
	Status int `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
	// Reason of the error
	// in: string
	Error error `json:"error"`
}

func (ce CommonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}{
		Status:  ce.Status,
		Message: ce.Message,
		Error:   ce.Error.Error(),
	})
}
