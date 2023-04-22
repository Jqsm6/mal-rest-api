package main

import (
	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/rl404/go-malscraper"
	"mal-rest-api/internal/config"
	"mal-rest-api/internal/repository"
	"mal-rest-api/internal/server"
	"mal-rest-api/internal/server/handlers"
	"mal-rest-api/internal/service"
	"mal-rest-api/pkg/db/redis"
	"mal-rest-api/pkg/logger"
)

var ctx = context.Background()

func main() {
	logger := logger.GetLogger()
	cfg := config.GetConfig()
	router := httprouter.New()
	httpServer := server.NewServer(router, logger)

	mal, err := malscraper.NewDefault()
	if err != nil {
		logger.Fatal().Err(err)
	}

	rdb := redis.NewRedisClient(cfg)
	rep := repository.NewRepository(ctx, rdb)

	animeService := service.NewAnimeService(rep, mal)
	characterService := service.NewCharacterService(rep, mal)
	userService := service.NewUserService(rep, mal)

	handlers.AnimeHandlerRegister(ctx, router, animeService)
	handlers.CharacterHandlerRegister(ctx, router, characterService)
	handlers.UserHandlerRegister(ctx, router, userService)

	httpServer.Run()
}
