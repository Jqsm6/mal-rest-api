package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"mal-rest-api/pkg/logger"
)

func Logger(next httprouter.Handle, logger *logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		logger.Info().
			Str("method:", r.Method).
			Str("url:", r.URL.String()).
			Str("from:", r.RemoteAddr).
			Msg("")
		next(w, r, params)
	}
}
