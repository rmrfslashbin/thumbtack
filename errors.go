package thumbtack

// ErrBadEndpoint is returned when the endpoint is not valid
type ErrBadEndpoint struct {
	Err error
	Msg string
}

// Error returns the error message
func (e *ErrBadEndpoint) Error() string {
	if e.Msg == "" {
		e.Msg = "endpoint is not valid"
	}
	if e.Err != nil {
		e.Msg += ": " + e.Err.Error()
	}
	return e.Msg
}

// ErrBadStatusCode is returned when the status code is not valid
type ErrBadStatusCode struct {
	StatusCode int
	Status     string
	Err        error
	Msg        string
}

// Error returns the error message
func (e *ErrBadStatusCode) Error() string {
	if e.Msg == "" {
		e.Msg = "endpoint returned a bad status code"
	}
	if e.StatusCode != 0 {
		e.Msg += ": " + e.Status
	}
	if e.Status != "" {
		e.Msg += ": " + e.Status
	}
	if e.Err != nil {
		e.Msg += ": " + e.Err.Error()
	}
	return e.Msg
}

// ErrInvalidInput is returned when the input is not valid
type ErrInvalidInput struct {
	Err error
	Msg string
}

// Error returns the error message
func (e *ErrInvalidInput) Error() string {
	if e.Msg == "" {
		e.Msg = "input is not valid or nil"
	}
	if e.Err != nil {
		e.Msg += ": " + e.Err.Error()
	}
	return e.Msg
}

// ErrMissingInputField is returned when the input is missing
type ErrMissingInputField struct {
	Err   error
	Msg   string
	Field string
}

// Error returns the error message
func (e *ErrMissingInputField) Error() string {
	if e.Msg == "" {
		e.Msg = "missing input"
	}
	if e.Field != "" {
		e.Msg += ": " + e.Field
	}
	if e.Err != nil {
		e.Msg += ": " + e.Err.Error()
	}
	return e.Msg
}

// ErrNoToken is returned when no token is provided
type ErrNoToken struct {
	Err error
	Msg string
}

// Error returns the error message
func (e *ErrNoToken) Error() string {
	if e.Msg == "" {
		e.Msg = "no provided. use WithToken()"
	}
	if e.Err != nil {
		e.Msg += ": " + e.Err.Error()
	}
	return e.Msg
}

// ErrUnexpectedResponse is returned when the response is not valid
type ErrUnexpectedResponse struct {
	Err        error
	Msg        string
	ResultCode string
}

// Error returns the error message
func (e *ErrUnexpectedResponse) Error() string {
	if e.Msg == "" {
		e.Msg = "unexpected response"
	}
	if e.ResultCode != "" {
		e.Msg += ": " + e.ResultCode
	}
	if e.Err != nil {
		e.Msg += ": " + e.Err.Error()
	}
	return e.Msg
}
