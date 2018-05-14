package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
)

// GetValueAsInt parses integer values from HTTP request
func GetValueAsInt(r *http.Request, key string) int {
	value, err := strconv.Atoi(bone.GetValue(r, key))
	// TODO(claudio): do proper err handling
	if err != nil {
		panic(err)
	}

	return value
}

func respondNotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, message)
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
