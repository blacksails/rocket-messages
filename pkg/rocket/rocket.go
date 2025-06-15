package rocket

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/blacksails/rocket-messages/pkg/api"
	"github.com/blacksails/rocket-messages/pkg/message"
)

type Service interface {
	GetRocket(channel string) (*api.Rocket, error)
	ListRockets(request *api.ListRocketsRequest) ([]*api.Rocket, error)
}

type service struct {
	messageStore message.Store
}

var (
	ErrNotFound = errors.New("not found")
)

// GetRocket implements Service.
func (s *service) GetRocket(channel string) (*api.Rocket, error) {
	messages, err := s.messageStore.List(channel)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, ErrNotFound
	}

	r := &api.Rocket{ID: channel}

	for _, m := range messages {
		r.LastMessage = m.Metadata.MessageTime
		switch m.Metadata.MessageType {
		case api.MessageTypeRocketLaunched:
			var mes api.MessageRocketLaunched
			if err := json.Unmarshal(m.Message, &mes); err != nil {
				return nil, fmt.Errorf("could not unmarshal rocket launched message: %w", err)
			}
			r.Type = mes.Type
			r.Mission = mes.Mission
			r.Speed = mes.LaunchSpeed
		case api.MessageTypeRocketSpeedIncreased:
			var mes api.MessageRocketSpeedChanged
			if err := json.Unmarshal(m.Message, &mes); err != nil {
				return nil, fmt.Errorf("could not unmarshal rocket speed increase message: %w", err)
			}
			r.Speed += mes.By
		case api.MessageTypeRocketSpeedDecreased:
			var mes api.MessageRocketSpeedChanged
			if err := json.Unmarshal(m.Message, &mes); err != nil {
				return nil, fmt.Errorf("could not unmarshal rocket speed decrease message: %w", err)
			}
			r.Speed -= mes.By
		case api.MessageTypeRocketExploded:
			var mes api.MessageRocketExploded
			if err := json.Unmarshal(m.Message, &mes); err != nil {
				return nil, fmt.Errorf("could not unmarshal rocket explosion message: %w", err)
			}
			r.Speed = 0
			r.Exploded = true
			r.ExplosionReason = mes.Reason
		case api.MessageTypeRocketMissionChanged:
			var mes api.MessageRocketMissionChanged
			if err := json.Unmarshal(m.Message, &mes); err != nil {
				return nil, fmt.Errorf("could not unmarshal rocket mission changed message: %w", err)
			}
			r.Mission = mes.NewMission
		}
	}

	return r, nil
}

// ListRockets implements Service.
func (s *service) ListRockets(request *api.ListRocketsRequest) ([]*api.Rocket, error) {
	channels, err := s.messageStore.ListChannels()
	if err != nil {
		return nil, err
	}

	rockets := make([]*api.Rocket, len(channels))
	for i, c := range channels {
		rocket, err := s.GetRocket(c)
		if err != nil {
			return nil, err
		}
		rockets[i] = rocket
	}

	// TODO: implement sorting

	return rockets, nil
}

func NewService(store message.Store) Service {
	return &service{
		messageStore: store,
	}
}
