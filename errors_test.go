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
