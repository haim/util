package util

import (
	. "github.com/aandryashin/matchers"
	"net/http"
	"testing"
	"time"
)

func TestSecondsSince(t *testing.T) {
	instant := time.Now()
	time.Sleep(100 * time.Millisecond)
	AssertThat(t, SecondsSince(instant) > 0, Is{true})
}

func TestRequestInfoBasicAuthXForwardedFor(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://some-url.example.com/", nil)
	req.Header.Add("X-Forwarded-For", "some-addr.example.com")
	req.SetBasicAuth("some-user", "any-password")
	user, remote := RequestInfo(req)
	AssertThat(t, user, EqualTo{"some-user"})
	AssertThat(t, remote, EqualTo{"some-addr.example.com"})
}

func TestRequestInfoAnonymous(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://some-url.example.com/", nil)
	req.RemoteAddr = "some-addr.example.com:34256"
	user, remote := RequestInfo(req)
	AssertThat(t, user, EqualTo{"unknown"})
	AssertThat(t, remote, EqualTo{"some-addr.example.com"})
}
