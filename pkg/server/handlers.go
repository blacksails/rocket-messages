package server

import (
	"errors"
	"net/http"

	"github.com/blacksails/rocket-messages/pkg/api"
	"github.com/blacksails/rocket-messages/pkg/rocket"
)

func (s *Server) postMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := &api.Message{}
		if err := decode(r, m); err != nil {
			respondErr(w, http.StatusBadRequest, err)
			return
		}
		if err := s.messageStore.Save(m); err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}
		respond(w, http.StatusAccepted, nil)
	}
}

func (s *Server) getRocketHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ro, err := s.rocketService.GetRocket(r.PathValue("{id}"))
		if err != nil {
			if errors.Is(err, rocket.ErrNotFound) {
				respondErr(w, http.StatusNotFound, nil)
				return
			}
			respondErr(w, http.StatusInternalServerError, err)
			return
		}
		respond(w, http.StatusOK, ro)
	}
}

func (s *Server) listRocketsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rockets, err := s.rocketService.ListRockets(&api.ListRocketsRequest{})
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}
		respond(w, http.StatusOK, rockets)
	}
}
