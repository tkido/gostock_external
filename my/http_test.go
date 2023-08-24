package my

import (
	"errors"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "http://tkido.com/blog/"
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	if s := resp.StatusCode; s != 200 {
		t.Error(s)
	}
}

func TestGet200(t *testing.T) {
	url := "http://tkido.com/blog/"
	resp, err := Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != 200 {
		t.Errorf("http error status %d", s)
	}
}

func TestGet503(t *testing.T) {
	url := "http://ozuma.sakura.ne.jp/httpstatus/503"
	_, got := Get(url)
	want := errors.New(`"http://ozuma.sakura.ne.jp/httpstatus/503" returns http error status 503`)
	if got.Error() != want.Error() {
		t.Errorf("got %v want %v", got, want)
	}
}
