package scrape

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestServer(expectedOutput string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, expectedOutput)
	}))

}

func TestHNClient_ScrapeHNFeed(t *testing.T) {

}
