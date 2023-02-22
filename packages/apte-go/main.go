package apte

import (
	"log"
	"net/http"

	"github.com/mstephen19/apte/lib"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
)

func DefaultStartConfig() *StartConfig {
	return &StartConfig{
		Addr: ":3000",
		Cors: cors.Options{},
	}
}

// Options for creating an event server with NewEventServer().
type EventServerOptions struct {
	Endpoint string
	Http2    *http2.Server
}

type sse struct {
	// The options initially provided when creating the
	// event server with NewEventServer().
	Options *EventServerOptions
	lib.Routes
}

// The single handler which will be used when starting
// a server with .Start().
// Can also be used in a custom server.
func (server *sse) Handler() http.Handler {
	return lib.CreateHandler(server.Options.Endpoint, server.Options.Http2, server.Routes)
}

// Create a new event server that can be used for sending and
// receiving data to/from clients.
func NewEventServer(opts *EventServerOptions) *sse {
	router := make(lib.Routes)

	return &sse{
		Options: opts,
		Routes:  router,
	}
}

type StartConfig struct {
	Addr string
	Cors cors.Options
}

func (server *sse) Start(config *StartConfig) {
	handler := cors.New(config.Cors).Handler(server.Handler())
	h1 := http.Server{Addr: config.Addr, Handler: handler}
	log.Fatal(h1.ListenAndServe())
}
