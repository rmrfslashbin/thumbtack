package thumbtack

import "testing"

func TestApiUnknown(t *testing.T) {
	config := NewConfig()
	if _, err := config.GetAPI("foo"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestEmptyApi(t *testing.T) {
	config := NewConfig()
	config.SetAPI("NotesById", "")
	_, err := config.GetAPI("NotesById")
	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet")
	}
	if err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestEndpoint(t *testing.T) {
	config := NewConfig()
	endpoint := "https://example.com"
	config.SetEndpoint(endpoint)
	if config.GetEndpoint() != endpoint {
		t.Errorf("expected endpoint to be '%s', got '%s'", endpoint, config.GetEndpoint())
	}
}

func TestErrApiNotSetApiMsg(t *testing.T) {
	err := &ErrApiNotSet{Api: "foo", Msg: "bar"}
	expected := "bar: foo"
	if err.Error() != expected {
		t.Fatalf("expected error message to be 'foo', got '%s'", expected)
	}
}

func TestErrApiNotSetNoMsg(t *testing.T) {
	err := &ErrApiNotSet{}
	if err.Error() != "requested api is not set or empty" {
		t.Fatalf("expected error message to be 'requested api is not set or empty', got '%s'", err.Error())
	}
}

func TestErrUnknownApiApiMsg(t *testing.T) {
	err := &ErrUnknownApi{Api: "foo", Msg: "bar"}
	expected := "bar: foo"
	if err.Error() != expected {
		t.Fatalf("expected error message to be 'foo', got '%s'", expected)
	}
}

func TestErrUnknownApiNoMsg(t *testing.T) {
	err := &ErrUnknownApi{}
	if err.Error() != "requested api is not known or defined" {
		t.Fatalf("expected error message to be 'requested api is not known or defined', got '%s'", err.Error())
	}
}

func TestSetUnknownApi(t *testing.T) {
	config := NewConfig()
	if err := config.SetAPI("foo", "bar"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestUserAgent(t *testing.T) {
	config := NewConfig()
	useragent := "testagent/1.0"
	config.SetUserAgent(useragent)
	if config.GetUserAgent() != useragent {
		t.Errorf("expected useragent to be '%s', got '%s'", useragent, config.GetUserAgent())
	}
}

func TestVersion(t *testing.T) {
	config := NewConfig()
	version := "testversion"
	config.SetVersion(version)
	if config.GetVersion() != version {
		t.Errorf("expected version to be '%s', got '%s'", version, config.GetVersion())
	}
}
