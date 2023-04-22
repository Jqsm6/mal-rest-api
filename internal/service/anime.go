package service

import (
	"context"
	"net/http"

	"github.com/rl404/go-malscraper"
	"mal-rest-api/internal/repository"
)

type AnimeService struct {
	mal *malscraper.Malscraper
	rep *repository.Repository
}

func NewAnimeService(rep *repository.Repository, mal *malscraper.Malscraper) *AnimeService {
	return &AnimeService{
		rep: rep,
		mal: mal,
	}
}

func (as *AnimeService) GetByID(ctx context.Context, id int, cacheKey string) (interface{}, int, error) {
	val, err := as.rep.Get(ctx, cacheKey)
	if err == nil {
		return val, http.StatusOK, nil
	}

	data, code, err := as.mal.GetAnime(id)
	if err != nil {
		return data, code, err
	}

	as.rep.Set(ctx, cacheKey, data)

	return data, code, err
}

func (as *AnimeService) GetCharactersByID(ctx context.Context, id int, cacheKey string) (interface{}, int, error) {
	val, err := as.rep.Get(ctx, cacheKey)
	if err == nil {
		return val, http.StatusOK, nil
	}

	data, code, err := as.mal.GetAnimeCharacter(id)
	if err != nil {
		return data, code, err
	}

	as.rep.Set(ctx, cacheKey, data)

	return data, code, err
}
