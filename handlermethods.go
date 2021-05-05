package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"sync"
)

// ---
// handlers as server methods &
// handler-specific setup

// Route handlers as methods on server give access to server's dependencies.
// The downside is that sharing of common dependencies can lead to data races.

// handleSomething decorator returns a HandlerFunc closure to allow
// handler-specific setup
func (s *server) handleSomething() http.HandlerFunc {
	thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		thing() // use  thing in HandlerFunc
	}
}

// name handlers to play nicely with alphabetical docs & autocomplete:
func (s *server) handleTasksCreate() (h http.HandlerFunc) { return h }
func (s *server) handleTasksDone() (h http.HandlerFunc)   { return h }
func (s *server) handleTasksGet() (h http.HandlerFunc)    { return h }

func (s *server) handleAuthLogin() (h http.HandlerFunc)  { return h }
func (s *server) handleAuthLogout() (h http.HandlerFunc) { return h }

// ---
// handler method arguments &

// arguments for handler-specific dependencies:
func (s *server) handleGreeting(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, format, r.FormValue("name"))
	}
}
func (s *server) examplePassingFormats() {
	// router.HandleFunc likely from https://github.com/gorilla/mux#examples
	s.router.HandleFunc("/one", s.handleGreeting("Hello %s"))
	s.router.HandleFunc("/two", s.handleGreeting("Hola %s"))
	// ^ formats provided as arguments
}

// more examples of handler methods with arguments
func handleTemplate(template *template.Template) (h http.HandlerFunc) { return h }
func handleRandomQuote(q Quoter, r *rand.Rand) (h http.HandlerFunc)   { return h }
func handleSendMagicLinkEmail(e EmailSender) (h http.HandlerFunc)     { return h }

// ---
// inner types

func (s *server) handleGreet() (h http.HandlerFunc) {
	// declaring types inside functions:
	// * tells devs that these are function-only types
	// * declutters package space
	// * does not require long unique names for types
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}
	// ... more stuff
	return h
}

// ---
// lazy setup with sync.Once

func (s *server) handleTemplate(files ...string) http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			// setup only happens the handler is first called
			tpl, tplerr = template.ParseFiles(files...) // expensive
			// if handler never called, parsing is never done
			// ...not sure why this wouldn't be true anyway
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, "data") // use tpl
	}
}

// ---
// ANCILLARY STUBS (to make compiler happy)

// some prep in a handler method
func prepareThing() (thing func()) { return thing }

// Quoter quotes.
type Quoter struct{}
