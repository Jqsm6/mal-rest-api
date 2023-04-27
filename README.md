
# MyAnimeList RESTful API

The Anime API is a RESTful API that provides information about anime, characters, and users. It uses the MyAnimeList API to fetch data and Redis for caching.

## Getting Started

After installing Go and Docker, run the following commands to start experiencing this starter kit:
```bash
# Clone the repository
git clone https://github.com/Jqsm6/mal-rest-api.git

# Navigate to the project directory
cd mal-rest-api

# Copy the example configuration file and edit it to your needs
cp config.example.yml config.yml

# Start the application with Docker Compose
docker-compose up
```

At this time, you have a RESTful API server running at `http://localhost:8000`. It provides the following endpoints:

* `GET /api/anime/:id` - Get anime information by ID
* `GET /api/anime/:id/characters` - Get anime characters by ID
* `GET /api/character/:id` - Get character information by ID
* `GET /api/user/:username` - Get user information by username

## Full list what has been used

* [httprouter](https://github.com/julienschmidt/httprouter) - lightweight HTTP router
* [go-malscrapper](https://github.com/rl404/go-malscraper) - Another unofficial MyAnimeList API using Go.
* [Zerolog](https://github.com/rs/zerolog) - Zero Allocation JSON Logger
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Clean environment configuration reader for Go
* [go-redis](https://github.com/redis/go-redis) - Type-safe Redis client for Golang
* [Docker](https://www.docker.com/) & [docker-compose](https://docs.docker.com/compose/) - Docker

## Documentation

[Documentation](https://github.com/Jqsm6/mal-rest-api/wiki/REST-API-Documentation)

## TODO

* Add more methods.
