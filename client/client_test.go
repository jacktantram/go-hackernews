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
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, expectedOutput)
	}))

}

func TestHNClient_GetTopStories(t *testing.T) {
	tests := []struct {
		name     string
		want     []Item
		expected string
	}{
		{"first test for 1 stories", []Item{}, `<table class="hnmain">
		<table class="itemlist">
		<tr class='athing' id='20325395'>
				<td align="right" valign="top" class="title"><span class="rank">1.</span></td>      <td valign="top" class="votelinks"><center><a id='up_20325395' href='vote?id=20325395&amp;how=up&amp;goto=news'><div class='votearrow' title='upvote'></div></a></center></td><td class="title"><a href="https://opensource.googleblog.com/2019/07/googles-robotstxt-parser-is-now-open.html" class="storylink">Google’s robots.txt parser is now open source</a><span class="sitebit comhead"> (<a href="from?site=googleblog.com"><span class="sitestr">googleblog.com</span></a>)</span></td></tr><tr><td colspan="2"></td><td class="subtext">
				  <span class="score" id="score_20325395">440 points</span> by <a href="user?id=dankohn1" class="hnuser">dankohn1</a> <span class="age"><a href="item?id=20325395">6 hours ago</a></span> <span id="unv_20325395"></span> | <a href="hide?id=20325395&amp;goto=news">hide</a> | <a href="item?id=20325395">128&nbsp;comments</a>              </td></tr>
		</table>
		</table>
		`},
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
		name     string
		fields   fields
		want     Item
		expected string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := NewTestServer(tt.expected)
			defer testServer.Close()
			hnc := &HNClient{
				BaseURL: testServer.URL,
			}
			if got := hnc.GetItem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HNClient.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
