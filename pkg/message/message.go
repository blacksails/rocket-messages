package message

import (
	"slices"

	"github.com/blacksails/rocket-messages/pkg/api"
)

type Store interface {
	Save(message *api.Message) error
	List(channel string) ([]*api.Message, error)
	ListChannels() ([]string, error)
}

type InMemoryStore struct {
	messages map[string][]*api.Message
}

func NewInMemoryStore() Store {
	return &InMemoryStore{messages: map[string][]*api.Message{}}
}

func (s *InMemoryStore) Save(message *api.Message) error {
	ms := s.messages[message.Metadata.Channel]

	// If we already have the message, discard the new one.
	for _, m := range ms {
		if m.Metadata.MessageNumber == message.Metadata.MessageNumber {
			return nil
		}
	}

	// Append new message and ensure that stored messages are sorted
	ms = append(ms, message)
	slices.SortFunc(ms, func(a, b *api.Message) int {
		if a.Metadata.MessageNumber < b.Metadata.MessageNumber {
			return -1
		}
		if a.Metadata.MessageNumber > b.Metadata.MessageNumber {
			return 1
		}
		return 0
	})
	s.messages[message.Metadata.Channel] = ms

	return nil
}

func (s InMemoryStore) List(c string) ([]*api.Message, error) {
	return s.messages[c], nil
}

func (s InMemoryStore) ListChannels() ([]string, error) {
	if len(s.messages) == 0 {
		return nil, nil
	}
	channels := make([]string, len(s.messages))
	i := 0
	for channel := range s.messages {
		channels[i] = channel
		i++
	}
	slices.Sort(channels)
	return channels, nil
}
