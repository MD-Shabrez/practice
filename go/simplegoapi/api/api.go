package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int16
	Balance int64
}

type Error struct {
	Code int16
	Message string
}

func writeError(w http.ResponseWriter, message string, code int16) {
	resp := Error {
		Code: code,
		Message: message,
	};

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(int(code));

	json.NewEncoder(w).Encode(resp);
}

var (
	RequestErrorHandler = func (w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest);
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)