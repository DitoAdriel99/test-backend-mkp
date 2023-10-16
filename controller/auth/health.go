package auth

import (
	"encoding/json"
	"go-learn/library/response"
	"net/http"
)

func (c *_ControllerLogin) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		succResponse = response.NewResponse().
			WithStatus("Success").
			WithMessage("Success")
	)

	object, _ := json.Marshal(succResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(object)
}
