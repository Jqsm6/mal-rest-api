package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"mal-rest-api/internal/middleware"
	"mal-rest-api/internal/service"
	"mal-rest-api/pkg/logger"
	"mal-rest-api/pkg/utils"
)

const (
	getCharacterByID = "/api/character/:id"
)

type characterHandler struct {
	ctx context.Context
	s   *service.CharacterService
}

func CharacterHandlerRegister(ctx context.Context, r *httprouter.Router, s *service.CharacterService) {
	logger := logger.GetLogger()
	ch := characterHandler{ctx, s}

	r.GET(getCharacterByID, middleware.Logger(ch.GetCharacterByID, logger))
}

func (ch *characterHandler) GetCharacterByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "strconv error", err)
		return
	}

	cacheKey := "getCharacterByID:" + strconv.Itoa(id)

	data, code, err := ch.s.GetCharacterByID(ch.ctx, id, cacheKey)
	if err != nil {
		utils.ResponseWithJson(w, code, data, err)
		return
	}

	utils.ResponseWithJson(w, code, data, err)
}
