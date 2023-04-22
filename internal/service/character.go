package service

import (
	"context"
	"net/http"

	"github.com/rl404/go-malscraper"
	"mal-rest-api/internal/repository"
)

type CharacterService struct {
	rep *repository.Repository
	mal *malscraper.Malscraper
}

func NewCharacterService(rep *repository.Repository, mal *malscraper.Malscraper) *CharacterService {
	return &CharacterService{
		rep: rep,
		mal: mal,
	}
}

func (cs *CharacterService) GetCharacterByID(ctx context.Context, id int, cacheKey string) (interface{}, int, error) {
	val, err := cs.rep.Get(ctx, cacheKey)
	if err == nil {
		return val, http.StatusOK, nil
	}

	data, code, err := cs.mal.GetCharacterAnime(id)
	if err != nil {
		return data, code, err
	}

	cs.rep.Set(ctx, cacheKey, data)

	return data, code, err
}
