package thumbtack

import "testing"

func TestVersion(t *testing.T) {
	config := NewConfig()
	version := "testversion"
	config.SetVersion(version)
	if config.GetVersion() != version {
		t.Errorf("expected version to be '%s', got '%s'", version, config.GetVersion())
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

func TestUserAgent(t *testing.T) {
	config := NewConfig()
	useragent := "testagent/1.0"
	config.SetUserAgent(useragent)
	if config.GetUserAgent() != useragent {
		t.Errorf("expected useragent to be '%s', got '%s'", useragent, config.GetUserAgent())
	}
}

func TestApiUnknown(t *testing.T) {
	config := NewConfig()
	if _, err := config.GetAPI("foo"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestSetUnknownApi(t *testing.T) {
	config := NewConfig()
	if err := config.SetAPI("foo", "bar"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestEmptyApi(t *testing.T) {
	config := NewConfig()
	config.SetAPI("NotesById", "")
	if _, err := config.GetAPI("NotesById"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}
