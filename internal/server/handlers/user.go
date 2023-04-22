package handlers

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"mal-rest-api/internal/middleware"
	"mal-rest-api/internal/service"
	"mal-rest-api/pkg/logger"
	"mal-rest-api/pkg/utils"
)

const (
	getUserByUsername = "/api/user/:username"
)

type userHandler struct {
	ctx context.Context
	s   *service.UserService
}

func UserHandlerRegister(ctx context.Context, r *httprouter.Router, s *service.UserService) {
	logger := logger.GetLogger()
	uh := userHandler{ctx, s}

	r.GET(getUserByUsername, middleware.Logger(uh.GetUserByUsername, logger))
}

func (uh *userHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")

	cacheKey := "getUser:" + username

	data, code, err := uh.s.GetUserByUsername(uh.ctx, username, cacheKey)
	if err != nil {
		utils.ResponseWithJson(w, code, data, err)
		return
	}

	utils.ResponseWithJson(w, code, data, err)
}
