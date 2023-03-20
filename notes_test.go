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

func TestNotesById(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	notesByIdResp := `{"id":"xxxx67e342662e6c239c","title":"Test Note 01","created_at":"2023-03-19 14:35:16","updated_at":"2023-03-19 14:35:16","length":40,"text":"This is my test note to see how it works","hash":"xxxx910a03859fd9e80a"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != constants.NotesById+"/xxxx67e342662e6c239c" {
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
		WithToken(token),
		WithLogger(&log),
		WithUserAgent(useragent),
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

func TestNotesList(t *testing.T) {
	token := "test:abc123"
	useragent := "test/1.0"
	notesListResp := `{"count":1,"notes":[{"0":"xxxx67e342662e6c239c","id":"xxxx67e342662e6c239c","1":"xxxx910a03859fd9e80a","hash":"xxxx910a03859fd9e80a","2":"Test Note 01","title":"Test Note 01","3":40,"length":40,"4":"2023-03-19 14:35:16","created_at":"2023-03-19 14:35:16","5":"2023-03-19 14:35:16","updated_at":"2023-03-19 14:35:16"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != constants.NotesList {
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
		WithToken(token),
		WithLogger(&log),
		WithUserAgent(useragent),
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
