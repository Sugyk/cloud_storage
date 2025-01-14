package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fir1/rest-api/pkg/erru"
	"github.com/gorilla/mux"
)

func (s service) GetFile() http.HandlerFunc {
	type response struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		getResponse, err := s.cloudStorageService.GetFile(r.Context(), id)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			Id: getResponse.Id,
		}, http.StatusOK)
	}
}
