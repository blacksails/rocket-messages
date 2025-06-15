package rocket_test

import (
	"testing"

	"github.com/blacksails/rocket-messages/pkg/api"
	"github.com/blacksails/rocket-messages/pkg/message"
	"github.com/blacksails/rocket-messages/pkg/rocket"
	"github.com/stretchr/testify/assert"
)

func TestGetRocket(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc        string
		messages    []*api.Message
		channel     string
		expected    *api.Rocket
		expectedErr error
	}{
		{
			desc:        "not found",
			channel:     "test",
			expectedErr: rocket.ErrNotFound,
		},
		{
			desc:    "speed is set from launch message",
			channel: "test",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketLaunched,
						Channel:       "test",
						MessageNumber: 1,
					},
					Message: api.MessageRocketLaunched{
						Type:        "testtype",
						Mission:     "testmission",
						LaunchSpeed: 100,
					},
				},
			},
			expected: &api.Rocket{
				Type:    "testtype",
				Mission: "testmission",
				Speed:   100,
			},
		},
		{
			desc:    "speed is increased from increase message",
			channel: "test",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketLaunched,
						Channel:       "test",
						MessageNumber: 1,
					},
					Message: api.MessageRocketLaunched{
						Type:        "testtype",
						Mission:     "testmission",
						LaunchSpeed: 100,
					},
				},
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketSpeedIncreased,
						Channel:       "test",
						MessageNumber: 2,
					},
					Message: api.MessageRocketSpeedChanged{
						By: 50,
					},
				},
			},
			expected: &api.Rocket{
				Type:    "testtype",
				Mission: "testmission",
				Speed:   150,
			},
		},
		{
			desc:    "speed is decreased from decrease message",
			channel: "test",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketLaunched,
						Channel:       "test",
						MessageNumber: 1,
					},
					Message: api.MessageRocketLaunched{
						Type:        "testtype",
						Mission:     "testmission",
						LaunchSpeed: 100,
					},
				},
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketSpeedDecreased,
						Channel:       "test",
						MessageNumber: 2,
					},
					Message: api.MessageRocketSpeedChanged{
						By: 50,
					},
				},
			},
			expected: &api.Rocket{
				Type:    "testtype",
				Mission: "testmission",
				Speed:   50,
			},
		},
		{
			desc:    "explosion",
			channel: "test",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketLaunched,
						Channel:       "test",
						MessageNumber: 1,
					},
					Message: api.MessageRocketLaunched{
						Type:        "testtype",
						Mission:     "testmission",
						LaunchSpeed: 100,
					},
				},
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketExploded,
						Channel:       "test",
						MessageNumber: 2,
					},
					Message: api.MessageRocketExploded{
						Reason: "test",
					},
				},
			},
			expected: &api.Rocket{
				Type:            "testtype",
				Mission:         "testmission",
				Speed:           0,
				Exploded:        true,
				ExplosionReason: "test",
			},
		},
		{
			desc:    "mission changed",
			channel: "test",
			messages: []*api.Message{
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketLaunched,
						Channel:       "test",
						MessageNumber: 1,
					},
					Message: api.MessageRocketLaunched{
						Type:        "testtype",
						Mission:     "testmission",
						LaunchSpeed: 100,
					},
				},
				{
					Metadata: api.MessageMetadata{
						MessageType:   api.MessageTypeRocketMissionChanged,
						Channel:       "test",
						MessageNumber: 2,
					},
					Message: api.MessageRocketMissionChanged{
						NewMission: "newmission",
					},
				},
			},
			expected: &api.Rocket{
				Type:    "testtype",
				Mission: "newmission",
				Speed:   100,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			store := message.NewInMemoryStore()
			for _, m := range test.messages {
				assert.NoError(t, store.Save(m))
			}

			s := rocket.NewService(store)

			actual, actualErr := s.GetRocket(test.channel)
			assert.Equal(t, test.expectedErr, actualErr)
			assert.Equal(t, test.expected, actual)
		})
	}
}
