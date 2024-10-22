package pubsub

import (
	"sync"

	inj "github.com/lucasvillarinho/plumber/internal/injector"
)

type PubSub[T any] struct {
	subscribers []chan T

	mu sync.RWMutex

	closed bool
}

func NewPubSub[T any](_ *inj.Injector) (*PubSub[T], error) {
	return &PubSub[T]{
		subscribers: make([]chan T, 0),
	}, nil
}

func (s *PubSub[T]) Subscribe() <-chan T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}

	r := make(chan T)
	s.subscribers = append(s.subscribers, r)

	return r
}

func (s *PubSub[T]) Publish(value T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return

	}

	for _, ch := range s.subscribers {
		ch <- value
	}
}
