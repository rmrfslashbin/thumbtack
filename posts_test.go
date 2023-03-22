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

// TestPostsAdd tests the PostsAdd method
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

// TestPostsAddBadApiEndpoint tests the PostsAdd method with a bad API endpoint
func TestPostsAddBadApiEndpoint(t *testing.T) {
	token := "test:abc123"
	config := NewConfig()
	config.SetAPI("PostsAdd", "")

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
		WithConfigs(config),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()
	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAddBadData tests the PostsAdd method with bad data
func TestPostsAddBadData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsAddResp := "garbage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()
	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsAddErrorResult tests the PostsAdd method with an error result
func TestPostsAddErrorResult(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsAddResp := `{"result_code":"error"}`
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()
	_, err = client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	})
	if _, ok := err.(*ErrUnexpectedResponse); !ok {
		t.Fatalf("expected error to be of type ErrUnexpectedResponse, got %T", err)
	}
}

// TestPostsAddInputFieldsTrue tests the PostsAdd method with input fields set to true
func TestPostsAddInputFieldsTrue(t *testing.T) {

	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	//addTrue := true
	addFalse := false
	addTime := time.Now()
	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addFalse,
		Shared:      &addFalse,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addFalse,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAddInputNil tests the PostsAdd method with nil input
func TestPostsAddInputNil(t *testing.T) {

	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	if _, err := client.PostsAdd(nil); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAddInputTagsGet100 tests the PostsAdd method with 100 tags
func TestPostsAddInputTagsGet100(t *testing.T) {

	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()

	var tags []string
	for i := 0; i < 101; i++ {
		tags = append(tags, "test")
	}

	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        tags,
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAddInputTimestampgt10 tests the PostsAdd method with a timestamp greater than 10 minutes
func TestPostsAddInputTimestampgt10(t *testing.T) {

	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now().Add(time.Hour)
	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAddInputTitleNil tests the PostsAdd method with a nil title
func TestPostsAddInputTitleNil(t *testing.T) {

	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	addUrl := "https://example.com"
	addDescr := "Example Description"
	//addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()
	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         &addUrl,
		Title:       nil, //&addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAddInputUrlNil tests the PostsAdd method with a nil url
func TestPostsAddInputUrlNil(t *testing.T) {

	token := "test:abc123"

	client, err := New(
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	//addUrl := "https://example.com"
	addDescr := "Example Description"
	addTitle := "Example Title"
	addTrue := true
	addTime := time.Now()
	if _, err := client.PostsAdd(&PostsAddInput{
		Url:         nil, //&addUrl,
		Title:       &addTitle,
		Description: &addDescr,
		Replace:     &addTrue,
		Shared:      &addTrue,
		Tags:        []string{"test", "example"},
		Timestamp:   &addTime,
		ToRead:      &addTrue,
	}); err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAll tests the PostsAll method
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

// TestPostsAllBadApiCall tests the PostsAll method with a bad api call
func TestPostsAllBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	config.SetAPI("PostsAll", "")

	client, err := New(
		WithConfigs(config),
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsAll(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsAllBadHttpResponse tests the PostsAll method with a bad http response
func TestPostsAllBadHttpResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	token := "test:abc123"

	client, err := New(
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsAll(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsAllBadReplyData tests the PostsAll method with bad reply data
func TestPostsAllBadReplyData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsAllResp := "gargage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsAll(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsAllInputNil tests the PostsAll method with a nil input
func TestPostsAllInputNil(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsAll(nil)
	if err != nil {
		t.Fatalf("failed to get all posts: %v", err)
	}
}

// TestPostsAllInputMetaFalse tests the PostsAll method with input meta false
func TestPostsAllInputMetaFalse(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	allFalse := false
	allResults := 2
	allStart := 0

	fromTime := time.Now().Add(-1 * time.Hour * 24 * 7)
	toTime := time.Now()
	postsAdd, err := client.PostsAll(&PostsAllInput{
		FromDT:  &fromTime,
		Meta:    &allFalse,
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

// TestPostsAllInputTagsgt3 tests the PostsAll method with input tags greater than 3
func TestPostsAllInputTagsgt3(t *testing.T) {
	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithLogger(&log),
		WithToken(&token),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsAll(&PostsAllInput{
		Tags: []string{"test", "example", "foo", "bar"},
	})

	if _, ok := err.(*ErrInvalidInput); !ok {
		t.Fatalf("expected ErrInvalidInput, got '%v'", err)
	}
}

// TestPostsDates tests the PostsDates method
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

// TestPostsDatesInputTagsgt3 tests the PostsDates method with input tags greater than 3
func TestPostsDatesInputTagsgt3(t *testing.T) {
	token := "test:abc123"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	client, err := New(
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDates([]string{"test", "example", "foo", "bar"})
	if _, ok := err.(*ErrInvalidInput); !ok {
		t.Fatalf("expected ErrInvalidInput, got '%v'", err)
	}

}

// TestPostsDatesBadApiCall tests the PostsDates method with a bad api call
func TestPostsDatesBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	config.SetAPI("PostsDates", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDates([]string{"test", "example"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsDatesBadHttpResponse tests the PostsDates method with a bad http response
func TestPostsDatesBadHttpResponse(t *testing.T) {
	token := "test:abc123"

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
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDates([]string{"test", "example"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsDatesBadResponseData tests the PostsDates method with bad response data
func TestPostsDatesBadResponseData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsDatesResp := "gargage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDates([]string{"test", "example"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsDates tests the PostsDates method
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

// TestPostsDeleteBadApiCall tests the PostsDelete method with a bad api call
func TestPostsDeleteBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)

	config.SetAPI("PostsDelete", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDelete("https://example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsDeleteBadHttpResponse tests the PostsDelete method with a bad http response
func TestPostsDeleteBadHttpResponse(t *testing.T) {
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

	_, err = client.PostsDelete("https://example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsDeleteBadReplyData tests the PostsDelete method with bad reply data
func TestPostsDeleteBadReplyData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsDeleteResp := "garbage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDelete("https://example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsDeleteBadApiCall tests the PostsDelete method with a bad api call
func TestPostsDeleteErrorResult(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsDeleteResp := `{"result_code":"something else"}`
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsDelete("https://example.com")
	if _, ok := err.(*ErrUnexpectedResponse); !ok {
		t.Fatalf("expected ErrorResult, got %T", err)
	}
}

// TestPostsGet tests the PostsGet method
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	timestamp, _ := time.Parse(time.RFC3339, "2023-03-20T16:30:35Z")
	getFalse := false
	getUrl := "https://example.com"
	postsGet, err := client.PostsGet(&PostsGetInput{
		Date: &timestamp,
		Meta: &getFalse,
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

// TestPostsGetBadApiCall tests the PostsGet method with a bad api call
func TestPostsGetBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("PostsGet", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsGet(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsGetBadHttpStatus tests the PostsGet method with a bad http status
func TestPostsGetBadHttpStatus(t *testing.T) {
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

	_, err = client.PostsGet(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsGetBadResponseData tests the PostsGet method with bad response data
func TestPostsGetBadResponseData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsGetResp := "garbage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsGet(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}

// TestPostsGetInputNil tests the PostsGet method with nil input
func TestPostsGetInputNil(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	postsGet, err := client.PostsGet(nil)
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

// TestPostsGetInputTagsgt3 tests the PostsGet method with input tags gt 3
func TestPostsGetInputTagsgt3(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	timestamp, _ := time.Parse(time.RFC3339, "2023-03-20T16:30:35Z")
	getFalse := false
	getUrl := "https://example.com"
	_, err = client.PostsGet(&PostsGetInput{
		Date: &timestamp,
		Meta: &getFalse,
		Tags: []string{"test", "example", "foo", "bar"},
		URL:  &getUrl,
	})
	if _, ok := err.(*ErrInvalidInput); !ok {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

// TestPostsRecent tests the PostsRecent method
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

// TestPostsRecentBadApiCall tests the PostsRecent method with a bad api call
func TestPostsRecentBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("PostsRecent", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsRecent(nil)
	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected ErrApiNotSet, got %v", err)
	}
}

// TestPostsRecentBadHttpStatus tests the PostsRecent method with a bad http status
func TestPostsRecentBadHttpStatus(t *testing.T) {
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

	_, err = client.PostsRecent(nil)
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected ErrBadStatusCode, got %v", err)
	}
}

// TestPostsRecentBadResonseData tests the PostsRecent method with bad response data
func TestPostsRecentBadResonseData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsRecentResp := "garbage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsRecent(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestPostsRecentInputNil tests the PostsRecent method with nil input
func TestPostsRecentInputNil(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	postsRecent, err := client.PostsRecent(nil)
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

// TestPostsRecentInputCountGt100 tests the PostsRecent method with input count > 100
func TestPostsRecentInputCountGt100(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	count := 200
	tags := []string{"example"}
	_, err = client.PostsRecent(&PostsRecentInput{Count: &count, Tags: tags})
	if _, ok := err.(*ErrInvalidInput); !ok {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

// TestPostsRecentInputTagsGt3 tests the PostsRecent method with input tags > 3
func TestPostsRecentInputTagsGt3(t *testing.T) {
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	count := 2
	tags := make([]string, 4)
	for i := 0; i < 4; i++ {
		tags[i] = "example"
	}
	_, err = client.PostsRecent(&PostsRecentInput{Count: &count, Tags: tags})
	if _, ok := err.(*ErrInvalidInput); !ok {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

func TestPostsSuggestBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("PostsSuggest", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsSuggest("https://example.com")
	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected ErrApiNotSet, got %v", err)
	}
}

func TestPostsSuggestBadHttpResponse(t *testing.T) {
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

	_, err = client.PostsSuggest("https://example.com")
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected ErrBadStatusCode, got %v", err)
	}
}

func TestPostsSuggestBadResponseData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsSuggestResp := "garbage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsSuggest("https://example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
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

func TestPostsUpdateBadApiCall(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	config.SetAPI("PostsUpdate", "")

	client, err := New(
		WithConfigs(config),
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsUpdate()
	if _, ok := err.(*ErrApiNotSet); !ok {
		t.Fatalf("expected ErrApiNotSet, got %v", err)
	}
}

func TestPostsUpdateBadHttpStatus(t *testing.T) {
	token := "test:abc123"
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
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsUpdate()
	if _, ok := err.(*ErrBadStatusCode); !ok {
		t.Fatalf("expected ErrBadStatusCode, got %v", err)
	}

}

func TestPostsUpdateResponseData(t *testing.T) {
	config := NewConfig()
	token := "test:abc123"
	useragent := "test/1.0"
	postsUpdateResp := "garbage"
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
		WithToken(&token),
		WithLogger(&log),
		WithUserAgent(&useragent),
	)
	if err != nil {
		t.Fatalf("failed to create thumbtask instance: %v", err)
	}

	_, err = client.PostsUpdate()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}
