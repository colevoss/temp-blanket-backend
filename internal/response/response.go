package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/colevoss/temperature-blanket-backend/internal/log"
)

type Map = map[string]any

func Ok(w http.ResponseWriter, r *http.Request, payload interface{}) error {
	return JSON(w, r, http.StatusOK, payload)
}

func Error(w http.ResponseWriter, r *http.Request, err error) error {
	var resError *ErrorResponse

	switch {
	case errors.As(err, &resError):
		JSON(w, r, resError.Code, resError)

	default:
		ServerError(w, r, err)
	}

	return nil
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.C(r.Context()).Error(err.Error())

	JSON(w, r, http.StatusInternalServerError, Map{
		"error": "Internal server error",
	})
}

func JSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) error {
	response, err := json.Marshal(payload)

	if err != nil {
		ServerError(w, r, err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}
