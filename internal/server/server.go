package server

import (
	cacheMov "Cache/internal/cache"
	"Cache/internal/movie"
	"Cache/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

var cacheObject = cacheMov.CachedMovies{}

func initCaching() {
	cacheObject.Movies = movie.GetPopularMovies()
	cacheObject.StartCaching()
}

func Start() {
	
	// Подключение у базе данных
	database.Start()

	// Инциализация процесса кеширования
	initCaching()

	route := gin.New()
	route.GET("/", getMovies)

	// Запуск сервера
	if err := route.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
