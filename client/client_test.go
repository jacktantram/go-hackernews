package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func NewTestServer(expectedOutput string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, expectedOutput)
	}))

}

func TestHNClient_GetTopStories(t *testing.T) {
	tests := []struct {
		name     string
		want     []Item
		expected string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testServer := NewTestServer(tt.expected)
			defer testServer.Close()
			hnc := HNClient{BaseURL: testServer.URL}
			if got := hnc.GetTopStories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HNClient.GetTopStories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHNClient_GetItem(t *testing.T) {
	type fields struct {
		BaseURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hnc := &HNClient{
				BaseURL: tt.fields.BaseURL,
			}
			if got := hnc.GetItem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HNClient.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
