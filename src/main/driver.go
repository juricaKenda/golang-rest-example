package main

import (
	"github.com/juricaKenda/golang-rest-example/src/infrastructure/logger"
	"github.com/juricaKenda/golang-rest-example/src/infrastructure/repository"
	"github.com/juricaKenda/golang-rest-example/src/server"
)

func main() {
	// injection
	log := logger.New()
	repo := repository.NewArticleRepo(log)

	s := server.New("localhost", 8080, repo, log, nil)
	_ = s.Start()

	/*
		MAKE IT CONFIGURABLE
			1. Build a simple server (github.com/gorilla/mux)
			2. Dependency injection on a repository
			3. Dependency injection on a logger
			6. Extract status codes

			4. Config package
			5. Enable autoload of env variables (github.com/joho/godotenv/autoload)
	*/
}
