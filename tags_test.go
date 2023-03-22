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

// TestTagsDelete tests the TagsDelete method
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

// TestTagsDeleteBadAPICall tests the TagsDelete method with a bad api call
func TestTagsDeleteBadAPICall(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("TagsDelete", "")

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

// TestTagsDeleteBadHttpResponse tests the TagsDelete method with a bad http response
func TestTagsDeleteBadHttpResponse(t *testing.T) {
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

	_, err = client.TagsDelete("api")
	if _, ok := err.(*ErrUnmarshalResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnmarshalResponse, got %T", err)
	}
}

// TestTagsDeleteBadHttpStatus tests the TagsDelete method with a bad http status
func TestTagsDeleteBadHttpStatus(t *testing.T) {
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

	_, err = client.TagsDelete("api")
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected error to be of type ErrBadStatusCode, got %T", err)
	}
}

// TestTagsDeleteUnexpectedResponse tests the TagsDelete method with an unexpected response
func TestTagsDeleteUnexpectedResponse(t *testing.T) {
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

	_, err = client.TagsDelete("api")
	if _, ok := err.(*ErrUnexpectedResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnexpectedResponse, got %T", err)
	}
}

// TestTagsGet tests the TagsGet method
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

// TestTagsGetBadAPICall tests the TagsGet method with a bad api call
func TestTagsGetBadAPICall(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("TagsGet", "")

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

	_, err = client.TagsGet()

	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

// TestTagsGetBadHttpResponse tests the TagsGet method with a bad http response
func TestTagsGetBadHttpResponse(t *testing.T) {
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

	_, err = client.TagsGet()
	if _, ok := err.(*ErrUnmarshalResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnmarshalResponse, got %T", err)
	}
}

// TestTagsGetBadHttpStatus tests the TagsGet method with a bad http status
func TestTagsGetBadHttpStatus(t *testing.T) {
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

	_, err = client.TagsGet()
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected error to be of type ErrBadStatusCode, got %T", err)
	}
}

// TestTagsRename tests the TagsRename method
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

// TestTagsRenameBadAPICall tests the TagsRename method with a bad api call
func TestTagsRenameBadAPICall(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("TagsRename", "")

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

// TestTagsRenameBadHttpResponse tests the TagsRename method with a bad http response
func TestTagsRenameBadHttpResponse(t *testing.T) {
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
	_, err = client.TagsRename(&TagsRenameInput{Old: &old, New: &new})
	if _, ok := err.(*ErrUnmarshalResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnmarshalResponse, got %T", err)
	}
}

// TestTagsRenameBadHttpStatus tests the TagsRename method with a bad http status
func TestTagsRenameBadHttpStatus(t *testing.T) {
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

	old := "old"
	new := "new"
	_, err = client.TagsRename(&TagsRenameInput{Old: &old, New: &new})
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected error to be of type ErrBadStatusCode, got %T", err)
	}
}

// TestTagsRenameUnexpectedResponse tests the TagsRename method with an unexpected response
func TestTagsRenameUnexpectedResponse(t *testing.T) {
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

// TestTagsRenameInputNil tests the TagsRename method with a nil input
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

	_, err = client.TagsRename(nil)
	if _, ok := err.(*ErrInvalidInput); !ok {
		t.Fatalf("expected error to be of type ErrInvalidInput, got %T", err)
	}
}

// TestTagsRenameInputOldNil tests the TagsRename method with a nil input
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
	_, err = client.TagsRename(&TagsRenameInput{Old: nil, New: &new})
	if _, ok := err.(*ErrMissingInputField); !ok {
		if err.(*ErrMissingInputField).Field != "Old" {
			t.Fatalf("expected error field to be Old, got %s", err.(*ErrMissingInputField).Field)
		}
		t.Fatalf("expected error to be of type ErrMissingInputField, got %T", err)
	}
}

// TestTagsRenameInputNewNil tests the TagsRename method with a nil input
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
	_, err = client.TagsRename(&TagsRenameInput{Old: &old, New: nil})
	if _, ok := err.(*ErrMissingInputField); !ok {
		if err.(*ErrMissingInputField).Field != "New" {
			t.Fatalf("expected error field to be New, got %s", err.(*ErrMissingInputField).Field)
		}
		t.Fatalf("expected error to be of type ErrMissingInputField, got %T", err)
	}
}
