package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

// ---
// testing servers
// https://youtu.be/rWBSMsLG8po?t=2011

func TestHandleAbout(t *testing.T) {
	is := is.New(t)
	srv := newServer() // create a new server inside each unit test
	db, cleanup := connectTestDatabase()
	defer cleanup()
	srv.db = db // only set the dependency you need
	r := httptest.NewRequest("GET", "/about", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusOK)
	// ^ Mat has just w.StatusCode (typo or maybe older version?)
}

// ---
// testing inner types
// https://youtu.be/rWBSMsLG8po?t=2229

func TestGreet(t *testing.T) {
	is := is.New(t)
	// handleGreet (in handlermethods.go) had an inner type request with a
	// Name property, so we make an anonymous struct with a Name and encode
	// as json
	p := struct {
		Name string `json:"name"`
	}{
		Name: "Mat Ryer",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	is.NoErr(err) // json.NewEncoder
	req := httptest.NewRequest(http.MethodPost, "/greet", &buf)
	// ... more test code here

	// ANCILLARY (to make compiler happy)
	is.Equal(req.Method, http.MethodPost)
}

// ---
// integration vs. unit tests
// https://youtu.be/rWBSMsLG8po?t=2324

func twoOrThreeWaysToTestServer(srv *server, w http.ResponseWriter, r *http.Request) {

	// test the whole stack (integration test)
	srv.ServeHTTP(w, r)

	// test just this handler (unit test)
	srv.handleGreet()(w, r)

	// third way: e2e tests with http.Server below:
}

// ---
// e2e tests
// https://youtu.be/rWBSMsLG8po?t=2408

// from https://pkg.go.dev/net/http/httptest?utm_source=gopls#Server
// An httptest.Server is an HTTP server listening on a system-chosen port on the
// local loopback interface, for use in end-to-end HTTP tests.

func TestTips(t *testing.T) {
	h := newFakeRemoteService()

	srv := httptest.NewServer(h)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/tips")

	// ANCILLARY (to make compiler happy)
	is := is.New(t)
	is.NoErr(err)
	is.Equal(200, resp.StatusCode)
}

// ---
// testing middleware
// https://youtu.be/rWBSMsLG8po?t=2474

func TestAdminOnly(t *testing.T) {
	var (
		n int
		w *responseWriter
		r *http.Request
	)
	is := is.New(t)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n++ // dummy handler increments n when called
	})
	srv := newServer()
	h = srv.adminOnly(h)

	// test: not admin
	r = httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Authorization", "Bearer no mate")
	h(w, r)
	is.Equal(n, 0) // should not get through

	// test: yes, admin!
	r = httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Authorization", "Bearer VALID")
	h(w, r)
	is.Equal(n, 1) // should get through
}

// ---
// ANCILLARY STUBS (to make compiler happy)

func connectTestDatabase() (db *dbConn, cleanup func()) { return db, cleanup }

func newFakeRemoteService() (h http.HandlerFunc) { return h }
