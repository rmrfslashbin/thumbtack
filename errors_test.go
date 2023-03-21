package thumbtack

import (
	"errors"
	"testing"
)

func TestErrBadEndpoint(t *testing.T) {
	err := ErrBadEndpoint{
		Err: errors.New("Testing subError"),
		Msg: "Testing ErrBadEndpoint",
	}
	errorOutput := err.Error()
	expectedOutput := "Testing ErrBadEndpoint: Testing subError"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrBadEndpointNoInput(t *testing.T) {
	err := ErrBadEndpoint{}
	errorOutput := err.Error()
	expectedOutput := "endpoint is not valid"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrBadStatusCode(t *testing.T) {
	err := ErrBadStatusCode{
		Err:        errors.New("Testing subError"),
		Msg:        "Testing ErrBadStatusCode",
		Status:     "test status",
		StatusCode: 800,
	}
	errorOutput := err.Error()
	expectedOutput := "Testing ErrBadStatusCode: 800 (test status): Testing subError"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrBadStatusCodeNoInput(t *testing.T) {
	err := ErrBadStatusCode{}
	errorOutput := err.Error()
	expectedOutput := "bad status code"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrInvalidInput(t *testing.T) {
	err := ErrInvalidInput{
		Err: errors.New("Testing subError"),
		Msg: "Testing ErrInvalidInput",
	}
	errorOutput := err.Error()
	expectedOutput := "Testing ErrInvalidInput: Testing subError"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrInvalidInputNoInput(t *testing.T) {
	err := ErrInvalidInput{}
	errorOutput := err.Error()
	expectedOutput := "input is not valid or nil"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrMissingInputField(t *testing.T) {
	err := ErrMissingInputField{
		Err: errors.New("Testing subError"),
		Msg: "Testing ErrMissingInputField",
	}
	errorOutput := err.Error()
	expectedOutput := "Testing ErrMissingInputField: Testing subError"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrMissingInputFieldNoInput(t *testing.T) {
	err := ErrMissingInputField{}
	errorOutput := err.Error()
	expectedOutput := "missing input"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrMissingInputFieldWithField(t *testing.T) {
	err := ErrMissingInputField{Field: "testField"}
	errorOutput := err.Error()
	expectedOutput := "missing input: testField"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrNoToken(t *testing.T) {
	err := ErrNoToken{
		Err: errors.New("Testing subError"),
		Msg: "Testing ErrNoToken",
	}
	errorOutput := err.Error()
	expectedOutput := "Testing ErrNoToken: Testing subError"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrNoTokenNoInput(t *testing.T) {
	err := ErrNoToken{}
	errorOutput := err.Error()
	expectedOutput := "no provided. use WithToken()"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrUnexpectedResponse(t *testing.T) {
	err := ErrUnexpectedResponse{
		Err: errors.New("Testing subError"),
		Msg: "Testing ErrUnexpectedResponse",
	}
	errorOutput := err.Error()
	expectedOutput := "Testing ErrUnexpectedResponse: Testing subError"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrUnexpectedResponseNoInput(t *testing.T) {
	err := ErrUnexpectedResponse{}
	errorOutput := err.Error()
	expectedOutput := "unexpected response"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}

func TestErrUnexpectedResponseWithResultCode(t *testing.T) {
	err := ErrUnexpectedResponse{ResultCode: "testCode"}
	errorOutput := err.Error()
	expectedOutput := "unexpected response: testCode"
	if errorOutput != expectedOutput {
		t.Errorf("Error() = %v, want %v", errorOutput, expectedOutput)
	}
}
