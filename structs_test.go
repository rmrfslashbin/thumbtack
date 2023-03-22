package thumbtack

import (
	"encoding/json"
	"testing"
)

// TestBookmarkStructBadTimestamp tests the Bookmark struct with a bad timestamp
func TestBookmarkStructBadTimestamp(t *testing.T) {
	data := []byte(`{"href":"https:\/\/example.com","description":"example post","extended":"this is the test post\/bookmark","meta":"258002234f7274ed91cd4c50ff2f65e7","hash":"c984d06aafbecf6bc55569f964148ea3","time":"20T16:30:35Z","shared":"no","toread":"no","tags":"test example"}`)
	bookmark := Bookmark{}
	if err := json.Unmarshal(data, &bookmark); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestBookmarkStructJsonUnmarshelFail tests the Bookmark struct with a bad json unmarshal
func TestBookmarkStructJsonUnmarshelFail(t *testing.T) {
	bookmark := Bookmark{}
	if err := json.Unmarshal([]byte("garbage"), &bookmark); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestBookmarkStructSharedAndToread tests the Bookmark struct with shared and toread set to yes
func TestBookmarkStructSharedAndToread(t *testing.T) {
	data := []byte(`{"href":"https:\/\/example.com","description":"example post","extended":"this is the test post\/bookmark","meta":"258002234f7274ed91cd4c50ff2f65e7","hash":"c984d06aafbecf6bc55569f964148ea3","time":"2023-03-20T16:30:35Z","shared":"yes","toread":"yes","tags":"test example"}`)
	bookmark := Bookmark{}
	if err := json.Unmarshal(data, &bookmark); err != nil {
		t.Error(err)
	}
}

// TestNoteStructBadCreatedAt tests the Note struct with a bad created_at timestamp
func TestNoteStructBadCreatedAt(t *testing.T) {
	notesByIdResp := `{"id":"xxxx67e342662e6c239c","title":"Test Note 01","created_at":"03-19 14:35:16","updated_at":"2023-03-19 14:35:16","length":40,"text":"This is my test note to see how it works","hash":"xxxx910a03859fd9e80a"}`
	notesById := Note{}
	if err := json.Unmarshal([]byte(notesByIdResp), &notesById); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestNoteStructBadUpdatedAt tests the Note struct with a bad updated_at timestamp
func TestNoteStructBadUpdatedAt(t *testing.T) {
	notesByIdResp := `{"id":"xxxx67e342662e6c239c","title":"Test Note 01","created_at":"2023-03-19 14:35:16","updated_at":"03-19 14:35:16","length":40,"text":"This is my test note to see how it works","hash":"xxxx910a03859fd9e80a"}`
	notesById := Note{}
	if err := json.Unmarshal([]byte(notesByIdResp), &notesById); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestNoteStructJsonUnmarshalFail tests the Bookmark struct with shared and toread set to yes
func TestNoteStructJsonUnmarshalFail(t *testing.T) {
	notesById := Note{}
	if err := json.Unmarshal([]byte(`{garbage}`), &notesById); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestSuggestionsStructJsonUnmarshalFail tests the Note struct with a bad updated_at timestamp
func TestSuggestionsStructJsonUnmarshalFail(t *testing.T) {
	suggestions := Suggestions{}
	if err := json.Unmarshal([]byte(`{garbage}`), &suggestions); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestTagsStructJsonUnmarshalFail tests the Note struct with a bad updated_at timestamp
func TestTagsStructJsonUnmarshalFail(t *testing.T) {
	tags := Tags{}
	if err := json.Unmarshal([]byte(`{garbage}`), &tags); err == nil {
		t.Error("Expected error, got nil")
	}
}
