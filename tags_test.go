package thumbtack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
)

func TestTagsDelete(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsDeleteResp := `{"result":"done"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsDelete")
		if err != nil {
			t.Fatalf("failed to get TagsDelete api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

func TestTagsDeleteBadAPICall(t *testing.T) {
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

	if _, err := client.TagsDelete("api"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsDeleteBadConfig(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()
	config.SetAPI("TagsDelete", "")

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

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

	_, err = client.TagsDelete("api")

	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

func TestTagsDeleteNotDone(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsDeleteResp := `{"result":"somethingelse"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsDelete")
		if err != nil {
			t.Fatalf("failed to get TagsDelete api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	if tagsDelete, _ := client.TagsDelete("api"); tagsDelete != nil {
		t.Fatalf("expected tagsDelete to be nil")
	}
}

func TestTagsDeleteWithBadData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsDeleteResp := "garbage"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsDelete")
		if err != nil {
			t.Fatalf("failed to get TagsDelete api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	if _, err := client.TagsDelete("api"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsGet(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsGetResp := `{"api":1,"books":1,"custom":1,"haproxy":2,"homepage":1,"logging":1}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsGet")
		if err != nil {
			t.Fatalf("failed to get TagsGet api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

func TestTagsGetBadAPICall(t *testing.T) {
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

	if _, err := client.TagsGet(); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsGetBadConfig(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()
	config.SetAPI("TagsGet", "")

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

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

	_, err = client.TagsDelete("api")

	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

func TestTagsGetWithBadData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsGetResp := "garbage"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsGet")
		if err != nil {
			t.Fatalf("failed to get TagsGet api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	if _, err := client.TagsGet(); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsRename(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsRenameResp := `{"result":"done"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsRename")
		if err != nil {
			t.Fatalf("failed to get TagsRename api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

func TestTagsRenameInputNil(t *testing.T) {
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

	if _, err := client.TagsRename(nil); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsRenameInputOldNil(t *testing.T) {
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

	new := "new"
	if _, err := client.TagsRename(&TagsRenameInput{Old: nil, New: &new}); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsRenameInputNewNil(t *testing.T) {
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

	old := "old"
	if _, err := client.TagsRename(&TagsRenameInput{Old: &old, New: nil}); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsRenameBadAPICall(t *testing.T) {
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

	old := "old"
	new := "new"
	if _, err := client.TagsRename(&TagsRenameInput{Old: &old, New: &new}); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsRenameBadConfig(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()
	config.SetAPI("TagsRename", "")

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

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

	old := "old"
	new := "new"
	_, err = client.TagsRename(&TagsRenameInput{Old: &old, New: &new})

	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

func TestTagsRenameWithBadData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsRenameResp := "garbage"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsRename")
		if err != nil {
			t.Fatalf("failed to get TagsRename api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	old := "old"
	new := "new"
	if _, err := client.TagsRename(&TagsRenameInput{Old: &old, New: &new}); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestTagsRenameResultNotDone(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	tagsRenameResp := `{"result":"something else"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("TagsRename")
		if err != nil {
			t.Fatalf("failed to get TagsRename api: %v", err)
		}
		if r.URL.Path != api {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	old := "old"
	new := "new"
	_, err = client.TagsRename(&TagsRenameInput{Old: &old, New: &new})
	if _, ok := err.(*ErrUnexpectedResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnexpectedResponse, got %T", err)
	}
}
