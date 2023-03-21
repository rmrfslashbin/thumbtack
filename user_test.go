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

func TestUsersBadAPICall(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithEndpoint(&url.URL{}),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	if _, err := client.UserSecret(); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestUserSecretBadResults(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	userSecretResp := "bad json"
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

	if _, err := client.UserSecret(); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}
