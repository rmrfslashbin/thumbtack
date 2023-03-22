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

// TestNotesById tests the NotesById method
func TestNotesById(t *testing.T) {
	config := NewConfig()

	token := "test:abc123"
	useragent := "test/1.0"
	notesByIdResp := `{"id":"xxxx67e342662e6c239c","title":"Test Note 01","created_at":"2023-03-19 14:35:16","updated_at":"2023-03-19 14:35:16","length":40,"text":"This is my test note to see how it works","hash":"xxxx910a03859fd9e80a"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("NotesById")
		if err != nil {
			t.Fatalf("failed to get NotesById api: %v", err)
		}

		if r.URL.Path != api+"/xxxx67e342662e6c239c" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		fmt.Fprint(w, notesByIdResp)
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

	notesById, err := client.NotesById("xxxx67e342662e6c239c")
	if err != nil {
		t.Fatalf("failed to get notes by id: %v", err)
	}

	if notesById == nil {
		t.Fatalf("expected notesById to not be nil")
	}

	if notesById.Id != "xxxx67e342662e6c239c" {
		t.Errorf("expected Id to be 'xxxx67e342662e6c239c', got '%s'", notesById.Id)
	}
}

// NotesByIdBadAPICall tests the NotesById method with a bad API call
func TestNotesByIdBadAPICall(t *testing.T) {
	token := "foo"
	config := NewConfig()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("NotesById", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.NotesById("xxxx67e342662e6c239c")

	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

// TestNotesByIdBadHttpResponse tests the NotesById method with bad data
func TestNotesByIdBadHttpResponse(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	notesByIdResp := `garbage`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("NotesById")
		if err != nil {
			t.Fatalf("failed to get NotesById api: %v", err)
		}
		if r.URL.Path != api+"/xxxx67e342662e6c239c" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		fmt.Fprint(w, notesByIdResp)
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

	if _, err := client.NotesById("xxxx67e342662e6c239c"); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestNotesList(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	notesListResp := `{"count":1,"notes":[{"0":"xxxx67e342662e6c239c","id":"xxxx67e342662e6c239c","1":"xxxx910a03859fd9e80a","hash":"xxxx910a03859fd9e80a","2":"Test Note 01","title":"Test Note 01","3":40,"length":40,"4":"2023-03-19 14:35:16","created_at":"2023-03-19 14:35:16","5":"2023-03-19 14:35:16","updated_at":"2023-03-19 14:35:16"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("NotesList")
		if err != nil {
			t.Fatalf("failed to get NotesList api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, notesListResp)
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

	notesList, err := client.NotesList()
	if err != nil {
		t.Fatalf("failed to get notes by list: %v", err)
	}

	if notesList == nil {
		t.Fatalf("expected notesList to not be nil")
	}
	if notesList.Count != 1 {
		t.Errorf("expected Count to be 1, got %d", notesList.Count)
	}
}

func TestNotesListBadAPICall(t *testing.T) {
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

	if _, err := client.NotesList(); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestNotesListBadConfig(t *testing.T) {
	useragent := "test/1.0"
	token := "foo"
	config := NewConfig()
	config.SetAPI("NotesList", "")

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

	_, err = client.NotesList()

	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected error to be of type ErrApiNotSet, got %T", err)
	}
}

func TestNotesListWithBadData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	notesListResp := `garbage`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("NotesList")
		if err != nil {
			t.Fatalf("failed to get NotesList api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, notesListResp)
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

	if _, err := client.NotesList(); err == nil {
		t.Fatalf("expected error to not be nil")
	}
}
