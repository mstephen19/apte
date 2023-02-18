package lib

import (
	"sync"
)

type receivers map[string]ReceiveHandler

type namespace struct {
	lock      *sync.Mutex
	stream    StreamHandler
	receivers receivers
}

// Register a function for streaming the
func (namespace *namespace) Stream(handler StreamHandler) {
	namespace.lock.Lock()
	defer namespace.lock.Unlock()

	namespace.stream = handler
}

func (namespace *namespace) Receive(messageType string, handler ReceiveHandler) {
	namespace.lock.Lock()
	defer namespace.lock.Unlock()

	namespace.receivers[messageType] = handler
}

func (namespace *namespace) Receiver(messageType string) (receiver ReceiveHandler, ok bool) {
	receiver, ok = namespace.receivers[messageType]
	return
}

func (namespace *namespace) Streamer() StreamHandler {
	return namespace.stream
}

// A map of maps where the keys pertain to a certain
// namespace.
type Routes map[string]*namespace

// Attempts to grab an existing namespace. If it doesn't exist,
// "ok" will be false.
func (routes Routes) getNamespace(name string) (namespace *namespace, ok bool) {
	namespace, ok = routes[name]
	return
}

// Attempts to grab an existing namespace. If it doesn't exist,
// a new namespace will be created.
func (routes Routes) Namespace(name string) *namespace {
	if namespace, ok := routes[name]; ok {
		return namespace
	}

	routes[name] = &namespace{
		lock:      &sync.Mutex{},
		receivers: make(receivers),
	}
	return routes[name]
}
