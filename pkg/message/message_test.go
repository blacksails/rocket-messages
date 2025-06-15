package message_test

import (
	"testing"

	"github.com/blacksails/rocket-messages/pkg/api"
	"github.com/blacksails/rocket-messages/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryStore(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc             string
		messages         []*api.Message
		channel          string
		expectedList     []*api.Message
		expectedChannels []string
	}{
		{
			desc: "no messages",
		},
		{
			desc:    "no messages with channel",
			channel: "test",
		},
		{
			desc:    "unordered messages",
			channel: "test",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						Channel:       "test",
						MessageNumber: 2,
					},
				},
				{
					Metadata: api.MessageMetadata{
						Channel:       "test",
						MessageNumber: 1,
					},
				},
			},
			expectedList: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						Channel:       "test",
						MessageNumber: 1,
					},
				},
				{
					Metadata: api.MessageMetadata{
						Channel:       "test",
						MessageNumber: 2,
					},
				},
			},
			expectedChannels: []string{"test"},
		},
		{
			desc:    "different channels",
			channel: "test A",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						Channel:       "test A",
						MessageNumber: 1,
					},
				},
				{
					Metadata: api.MessageMetadata{
						Channel:       "test B",
						MessageNumber: 1,
					},
				},
			},
			expectedList: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						Channel:       "test A",
						MessageNumber: 1,
					},
				},
			},
			expectedChannels: []string{"test A", "test B"},
		},
		{
			desc: "channels sorted",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						Channel:       "test B",
						MessageNumber: 1,
					},
				},
				{
					Metadata: api.MessageMetadata{
						Channel:       "test A",
						MessageNumber: 1,
					},
				},
				{
					Metadata: api.MessageMetadata{
						Channel:       "test C",
						MessageNumber: 1,
					},
				},
			},
			expectedChannels: []string{"test A", "test B", "test C"},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			s := message.NewInMemoryStore()

			for _, m := range test.messages {
				assert.NoError(t, s.Save(m))
			}

			if test.channel != "" {
				actualList, err := s.List(test.channel)
				assert.NoError(t, err)
				assert.Equal(t, test.expectedList, actualList)
			}

			actualChannels, err := s.ListChannels()
			assert.NoError(t, err)
			assert.Equal(t, test.expectedChannels, actualChannels)
		})
	}
}
