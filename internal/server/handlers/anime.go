package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"mal-rest-api/internal/middleware"
	"mal-rest-api/pkg/logger"
	"mal-rest-api/pkg/utils"

	"mal-rest-api/internal/service"
)

const (
	getByID           = "/api/anime/:id"
	getCharactersByID = "/api/anime/:id/characters"
)

type animeHandler struct {
	ctx context.Context
	s   *service.AnimeService
}

func AnimeHandlerRegister(ctx context.Context, r *httprouter.Router, s *service.AnimeService) {
	logger := logger.GetLogger()
	ah := animeHandler{ctx, s}

	r.GET(getByID, middleware.Logger(ah.GetByID, logger))
	r.GET(getCharactersByID, middleware.Logger(ah.GetCharactersByID, logger))
}

func (ah *animeHandler) GetByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "strconv error", err)
		return
	}

	cacheKey := "getByID:" + strconv.Itoa(id)

	data, code, err := ah.s.GetByID(ah.ctx, id, cacheKey)
	if err != nil {
		utils.ResponseWithJson(w, code, data, err)
		return
	}

	utils.ResponseWithJson(w, code, data, err)
}

func (ah *animeHandler) GetCharactersByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "strconv.Atoi error", err)
		return
	}

	cacheKey := "getCharactersByID:" + strconv.Itoa(id)

	data, code, err := ah.s.GetCharactersByID(ah.ctx, id, cacheKey)
	if err != nil {
		utils.ResponseWithJson(w, code, data, err)
		return
	}

	utils.ResponseWithJson(w, code, data, err)
}
