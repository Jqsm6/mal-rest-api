package service

import (
	"context"
)

type Anime interface {
	GetByID(ctx context.Context, id int, cacheKey string) (interface{}, int, error)
	GetCharactersByID(ctx context.Context, id int, cacheKey string) (interface{}, int, error)
}

type Character interface {
	GetCharacterByID(ctx context.Context, id int, cacheKey string) (interface{}, int, error)
}

type User interface {
	GetUserByUsername(ctx context.Context, username string, cacheKey string) (interface{}, int, error)
}

type Service struct {
	Anime
	Character
	User
}
