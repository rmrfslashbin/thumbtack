package thumbtack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
)

func TestPostsAdd(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsAddResp := `{"result_code":"done"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsAdd")
		if err != nil {
			t.Fatalf("failed to get PostsAdd api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsAddResp)
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

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()
	postsAdd, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	})
	if err != nil {
		spew.Dump(postsAdd)
		t.Fatalf("failed to add post: %v", err)
	}

	if postsAdd == nil {
		t.Fatalf("expected postsAdd to not be nil")
	}
	if postsAdd.ResultCode != "done" {
		t.Errorf("expected ResultCode to be 'done', got '%s'", postsAdd.ResultCode)
	}
}

func TestPostsAll(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsAllResp := `[{"href":"https:\/\/example.com","description":"example post","extended":"this is the test post\/bookmark","meta":"258002234f7274ed91cd4c50ff2f65e7","hash":"c984d06aafbecf6bc55569f964148ea3","time":"2023-03-20T16:30:35Z","shared":"no","toread":"no","tags":"test example"},
	{"href":"https:\/\/rmrfslashbin.io","description":"my homepage","extended":"This is the descr of the bookmark","meta":"36b8648620031b4805d6caa45a5e6a1d","hash":"ed59c2e5b3eb284b178ffae1dcfffc08","time":"2023-03-19T22:53:10Z","shared":"no","toread":"no","tags":"homepage personal test example"}]`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsAll")
		if err != nil {
			t.Fatalf("failed to get PostsAll api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsAllResp)
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

	allTrue := true
	allResults := 2
	allStart := 0

	fromTime := time.Now().Add(-1 * time.Hour * 24 * 7)
	toTime := time.Now()
	postsAdd, err := client.PostsAll(&PostsAllInput{
		FromDT:  &fromTime,
		Meta:    &allTrue,
		Results: &allResults,
		Start:   &allStart,
		Tags:    []string{"test", "example"},
		ToDT:    &toTime,
	})
	if err != nil {
		spew.Dump(postsAdd)
		t.Fatalf("failed to get all posts: %v", err)
	}

	if postsAdd == nil {
		t.Fatalf("expected postsAdd to not be nil")
	}
	if len(*postsAdd) != 2 {
		t.Errorf("expected two posts, got '%d'", len(*postsAdd))
	}
}

func TestPostsDates(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsDatesResp := `{"user":"test","tag":"","dates":{"2023-04-20":1,"2023-04-21":4}}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsDates")
		if err != nil {
			t.Fatalf("failed to get PostsDates api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsDatesResp)
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

	postsDates, err := client.PostsDates([]string{"test", "example"})
	if err != nil {
		spew.Dump(postsDates)
		t.Fatalf("failed to get all posts: %v", err)
	}

	if postsDates == nil {
		t.Fatalf("expected postsDates to not be nil")
	}
	if postsDates.User != "test" {
		t.Errorf("expected User to be 'test', got '%s'", postsDates.User)
	}
}

func TestPostsDelete(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsDeleteResp := `{"result_code":"done"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsDelete")
		if err != nil {
			t.Fatalf("failed to get PostsDelete api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsDeleteResp)
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

	postsDelete, err := client.PostsDelete("https://example.com")
	if err != nil {
		spew.Dump(postsDelete)
		t.Fatalf("failed to delete post: %v", err)
	}

	if postsDelete == nil {
		t.Fatalf("expected postsDelete to not be nil")
	}
	if postsDelete.ResultCode != "done" {
		t.Errorf("expected ResultCode to be 'done', got '%s'", postsDelete.ResultCode)
	}
}

func TestPostsGet(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsGetResp := `{"date":"2023-03-20T16:30:35Z","user":"test","posts":[{"href":"https:\/\/example.com","description":"example post","extended":"this is the test post\/bookmark","meta":"258002234f7274ed91cd4c50ff2f65e7","hash":"c984d06aafbecf6bc55569f964148ea3","time":"2023-03-20T16:30:35Z","shared":"no","toread":"no","tags":"test example"}]}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsGet")
		if err != nil {
			t.Fatalf("failed to get PostsGet api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsGetResp)
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

	timestamp, _ := time.Parse(time.RFC3339, "2023-03-20T16:30:35Z")
	getTrue := true
	getUrl := "https://example.com"
	postsGet, err := client.PostsGet(&PostsGetInput{
		Date: &timestamp,
		Meta: &getTrue,
		Tags: []string{"test", "example"},
		URL:  &getUrl,
	})
	if err != nil {
		spew.Dump(postsGet)
		t.Fatalf("failed to get posts: %v", err)
	}

	if postsGet == nil {
		t.Fatalf("expected postsGet to not be nil")
	}
	if postsGet.User != "test" {
		t.Errorf("expected user 'test', got '%s'", postsGet.User)
	}
}

func TestPostsRecent(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsRecentResp := `{"date":"2023-03-20T16:30:35Z","user":"test","posts":[{"href":"https:\/\/example.com","description":"example post","extended":"this is the test post\/bookmark","meta":"258002234f7274ed91cd4c50ff2f65e7","hash":"c984d06aafbecf6bc55569f964148ea3","time":"2023-03-20T16:30:35Z","shared":"no","toread":"no","tags":"test example"}]}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsRecent")
		if err != nil {
			t.Fatalf("failed to get PostsRecent api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsRecentResp)
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

	count := 2
	tags := []string{"example"}
	postsRecent, err := client.PostsRecent(&PostsRecentInput{Count: &count, Tags: tags})
	if err != nil {
		spew.Dump(postsRecent)
		t.Fatalf("failed to get recent posts: %v", err)
	}

	if postsRecent == nil {
		t.Fatalf("expected postsRecent to not be nil")
	}
	if postsRecent.User != "test" {
		t.Errorf("expected user 'test', got '%s'", postsRecent.User)
	}
}

func TestPostsSuggest(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsSuggestResp := `[{"popular":["Unread","pinboard","bookmarks"]},{"recommended":["Unread","pinboard","bookmarks"]}]`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsSuggest")
		if err != nil {
			t.Fatalf("failed to get PostsSuggest api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsSuggestResp)
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

	postsSuggestions, err := client.PostsSuggest("https://example.com")
	if err != nil {
		spew.Dump(postsSuggestions)
		t.Fatalf("failed to get suggestions: %v", err)
	}

	if postsSuggestions == nil {
		t.Fatalf("expected postsSuggestions to not be nil")
	}
	if len(postsSuggestions.Popular) != 3 {
		t.Errorf("expected '3' suggetions, got '%d'", len(postsSuggestions.Popular))
	}
}

func TestPostsUpdate(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsUpdateResp := `{"update_time":"2023-03-20T16:44:37Z"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api, err := config.GetAPI("PostsUpdate")
		if err != nil {
			t.Fatalf("failed to get PostsUpdate api: %v", err)
		}
		if r.URL.Path != api {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		fmt.Fprint(w, postsUpdateResp)
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

	postsUpdate, err := client.PostsUpdate()
	if err != nil {
		spew.Dump(postsUpdate)
		t.Fatalf("failed to get posts update: %v", err)
	}

	if postsUpdate == nil {
		t.Fatalf("expected postsUpdate to not be nil")
	}
	expectedUpdateTime, _ := time.Parse(time.RFC3339, "2023-03-20T16:44:37Z")
	if postsUpdate.UpdateTime != expectedUpdateTime {
		t.Errorf("expected update time to be '%s', got '%s'", expectedUpdateTime, postsUpdate.UpdateTime)
	}
}
