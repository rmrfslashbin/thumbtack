package thumbtack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

// TestUserSecret tests the UserSecret method
func TestUserSecret(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	userSecretResp := `{"result":"a6131d72761167f08be4"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("UserSecret")
		if err != nil {
			t.Fatalf("failed to get UserSecret api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, userSecretResp)
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

	userSecret, err := client.UserSecret()
	if err != nil {
		t.Fatalf("failed to get user secret: %v", err)
	}
	if userSecret == nil {
		t.Fatalf("expected userSecret to not be nil")
	}
	if userSecret.Result != "a6131d72761167f08be4" {
		t.Errorf("expected Result to be 'a6131d72761167f08be4', got '%s'", userSecret.Result)
	}
}

// TestUserSecretBadAPICall tests the UserSecret method with a bad api call
func TestUsersBadAPICall(t *testing.T) {
	config := NewConfig()
	useragent := "test/1.0"
	token := "foo"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("UserSecret", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.UserSecret()
	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

func TestUserSecretBadHttpResponse(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	userSecretResp := "garbage"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("UserSecret")
		if err != nil {
			t.Fatalf("failed to get UserSecret api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, userSecretResp)
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
	if _, ok := err.(*ErrUnmarshalResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnmarshalResponse, got %T", err)
	}
}

func TestUserSecretBadHttpStatus(t *testing.T) {
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
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected error to be of type ErrBadStatusCode, got %T", err)
	}
}
