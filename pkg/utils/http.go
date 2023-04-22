package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
	Meta    map[string]interface{} `json:"meta"`
}

func ResponseWithJson(w http.ResponseWriter, code int, data interface{}, err error, meta ...map[string]interface{}) {
	r := Response{
		Status:  code,
		Message: strings.ToLower(http.StatusText(code)),
	}

	r.Data = data
	if err != nil {
		r.Data = err.Error()
	}

	rJSON, _ := json.MarshalIndent(r, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(rJSON)))
	w.WriteHeader(code)

	_, _ = w.Write(rJSON)
}
