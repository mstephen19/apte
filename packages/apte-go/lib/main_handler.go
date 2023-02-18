package lib

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var STREAM_HEADERS = map[string]string{
	"Content-Type":      "text/event-stream",
	"Cache-Control":     "no-cache",
	"Connection":        "keep-alive",
	"Transfer-Encoding": "chunked",
}

func handleStream(w http.ResponseWriter, r *http.Request, namespace *namespace) {
	streamer := namespace.Streamer()
	if streamer == nil {
		return
	}

	// Initialize the stream
	for name, value := range STREAM_HEADERS {
		w.Header().Set(name, value)
	}
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	streamer(&StreamHandlerContext{
		BaseHandlerContext: &BaseHandlerContext{
			Writer:  w,
			Request: r,
		},
	})
}

func handleReceive(w http.ResponseWriter, r *http.Request, namespace *namespace) {
	messageType := r.URL.Query().Get("type")
	receiver, ok := namespace.Receiver(messageType)
	if !ok {
		return
	}

	receiver(&ReceiveHandlerContext{
		BaseHandlerContext: &BaseHandlerContext{
			Writer:  w,
			Request: r,
		},
	})

	w.WriteHeader(http.StatusAccepted)
}

func CreateHandler(endpoint string, h2 *http2.Server, routes Routes) http.Handler {
	if h2 == nil {
		h2 = &http2.Server{}
	}

	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignore paths that don't match the desired endpoint
		if r.URL.Path != endpoint {
			return
		}

		name := r.URL.Query().Get("namespace")
		// Do nothing if the namespace doesn't exist
		namespace, ok := routes.getNamespace(name)
		if !ok {
			return
		}

		switch r.Method {
		case http.MethodGet:
			handleStream(w, r, namespace)
			return
		case http.MethodPost:
			handleReceive(w, r, namespace)
			return
		default:
			http.NotFound(w, r)
		}
	}), h2)
}
