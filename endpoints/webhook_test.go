package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"avatar": "example_avatar", "channel_id": "123", "guild_id": "456", "id": "789", "name": "example_name"}`))
	}))
	defer server.Close()

	webhook, err := Fetch(server.URL)
	if err != nil {
		t.Errorf("Fetch returned an error: %v", err)
	}

	expected := &Webhook{
		Avatar:    "example_avatar",
		ChannelID: "123",
		GuildID:   "456",
		Id:        "789",
		Name:      "example_name",
	}
	if *webhook != *expected {
		t.Errorf("Fetch returned unexpected result, got: %v, want: %v", webhook, expected)
	}
}

func TestFetch_HTTPError(t *testing.T) {
	_, err := Fetch("invalid-url")
	if err == nil {
		t.Error("Fetch did not return an error for invalid URL")
	}
}

func TestFetch_JSONError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`invalid-json`))
	}))
	defer server.Close()

	_, err := Fetch(server.URL)
	if err == nil {
		t.Error("Fetch did not return an error for invalid JSON")
	}
}
