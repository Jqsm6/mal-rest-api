package service

import (
	"context"
	"net/http"

	"github.com/rl404/go-malscraper"
	"mal-rest-api/internal/repository"
)

type UserService struct {
	rep *repository.Repository
	mal *malscraper.Malscraper
}

func NewUserService(rep *repository.Repository, mal *malscraper.Malscraper) *UserService {
	return &UserService{
		mal: mal,
		rep: rep,
	}
}

func (us *UserService) GetUserByUsername(ctx context.Context, username string, cacheKey string) (interface{}, int, error) {
	val, err := us.rep.Get(ctx, cacheKey)
	if err == nil {
		return val, http.StatusOK, nil
	}

	data, code, err := us.mal.GetUser(username)
	if err != nil {
		return data, code, err
	}

	us.rep.Set(ctx, cacheKey, data)

	return data, code, err
}
