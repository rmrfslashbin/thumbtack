package thumbtack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack/internal/constants"
	"github.com/rs/zerolog"
)

func TestTagsGet(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	tagsGetResp := `{"api":1,"books":1,"custom":1,"haproxy":2,"homepage":1,"logging":1}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != constants.TagsGet {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, tagsGetResp)
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

	tagsGet, err := client.TagsGet()
	if err != nil {
		t.Fatalf("failed to get tags: %v", err)
	}
	if tagsGet == nil {
		t.Fatalf("expected tagsGet to be non-nil")
	}
	if tagsGet.Count != 6 {
		t.Errorf("expected Count to be 6, got %d", tagsGet.Count)
	}
}

func TestTagsDelete(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	tagsDeleteResp := `{"result":"done"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != constants.TagsDelete {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, tagsDeleteResp)
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

	tagsDelete, err := client.TagsDelete("api")
	if err != nil {
		spew.Dump(tagsDelete)
		t.Fatalf("failed to delete tag: %v", err)
	}
	if tagsDelete == nil {
		t.Fatalf("expected tagsDelete to be non-nil")
	}
	if tagsDelete.Result != "done" {
		t.Errorf("expected ResultCode to be 'done', got '%s'", tagsDelete.Result)
	}
}

func TestTagsRename(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	tagsRenameResp := `{"result":"done"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != constants.TagsRename {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, tagsRenameResp)
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

	old := "old"
	new := "new"
	tagsRename, err := client.TagsRename(&TagsRenameInput{Old: &old, New: &new})
	if err != nil {
		spew.Dump(tagsRename)
		t.Fatalf("failed to reanme tag: %v", err)
	}
	if tagsRename == nil {
		t.Fatalf("expected tagsRename to be non-nil")
	}
	if tagsRename.Result != "done" {
		t.Errorf("expected ResultCode to be 'done', got '%s'", tagsRename.Result)
	}
}
