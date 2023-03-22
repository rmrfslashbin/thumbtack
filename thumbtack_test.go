package thumbtack

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestThumbtackNoLogger(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"

	//log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	//zerolog.SetGlobalLevel(zerolog.PanicLevel)
	url, _ := url.Parse("https://example.com")

	client, err := New(
		WithEndpoint(url),
		WithToken(&token),
		//WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}
	if client == nil {
		t.Fatalf("expected client to not be nil")
	}
}

func TestThumbtackNoToken(t *testing.T) {
	//token := "test:abc123"
	useragent := "test/1.0"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	url, _ := url.Parse("https://example.com")

	_, err := New(
		WithEndpoint(url),
		//WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if _, ok := err.(*ErrNoToken); !ok {
		t.Fatalf("expected error to be ErrNoToken, got %v", err)
	}
}

func TestThumbtackNoEndpoint(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	//url, _ := url.Parse("https://example.com")

	_, err := New(
		WithEndpoint(nil),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}
}

func TestThumbtackNoUseragent(t *testing.T) {
	token := "test:abc123"
	//useragent := "test/1.0"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	url, _ := url.Parse("https://example.com")

	_, err := New(
		WithEndpoint(url),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(nil),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}
}

func TestThumbtackNillUseragent(t *testing.T) {
	token := "test:abc123"
	//useragent := "test/1.0"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	url, _ := url.Parse("https://example.com")

	_, err := New(
		WithEndpoint(url),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(nil),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}
}

func TestThumbtackBadHttpReply(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	url, _ := url.Parse(ts.URL)

	client, err := New(
		WithEndpoint(url),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.UserSecret()
	if err == nil {
		t.Fatalf("expected error to not be nil")
	}
	if v, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected error to be ErrBadStatusCode, got %v", v)
	}
}

func TestThumbtackWithConfig(t *testing.T) {
	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config := NewConfig()

	client, err := New(
		WithToken(&token),
		WithLogger(&log),
		WithConfigs(config),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.UserSecret()
	if err == nil {
		t.Fatalf("expected error to not be nil")
	}
	if v, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected error to be ErrBadStatusCode, got %v", v)
	}
}

func TestThumbtackBadHttpMethod(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetMethod("FOO")

	client, err := New(
		WithConfigs(config),
		WithEndpoint(&url.URL{}),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.UserSecret()
	if _, ok := err.(*url.Error); !ok {
		t.Fatalf("expected error to be of type url.Error, got %T", err)
	}
}
