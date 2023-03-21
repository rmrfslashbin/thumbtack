package thumbtack

import (
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
		WithToken(token),
		//WithLogger(&log),
		WithUserAgent(useragent),
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
		//WithToken(token),
		WithLogger(&log),
		WithUserAgent(useragent),
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
		WithToken(token),
		WithLogger(&log),
		WithUserAgent(useragent),
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
		WithToken(token),
		WithLogger(&log),
		WithUserAgent(""),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}
}
