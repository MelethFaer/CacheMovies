package movie

import (
	"Cache/internal/models"
	"Cache/pkg/database"
	"log"
)

func GetPopularMovies() *[]models.Movie {
	var movies []models.Movie
	if err := database.Connect.Find(&movies).Error; err != nil {
		log.Println(err)
	}
	return &movies
}
