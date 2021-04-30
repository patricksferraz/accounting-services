package model

type Response struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

func NewResponse(id, status, _error string) *Response {
	return &Response{
		ID:     id,
		Status: status,
		Error:  _error,
	}
}
