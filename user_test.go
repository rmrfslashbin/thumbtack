package thumbtack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/rmrfslashbin/thumbtack/internal/constants"
	"github.com/rs/zerolog"
)

func TestUserSecret(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	userSecretResp := `{"result":"a6131d72761167f08be4"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != constants.UserSecret {
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
		WithToken(token),
		WithLogger(&log),
		WithUserAgent(useragent),
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
