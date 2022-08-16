package cache

import (
	"Cache/internal/models"
	"Cache/internal/movie"
	"context"
	"log"
	"sync"
	"time"
)

type CachedMovies struct {
	lock   sync.Mutex
	Movies *[]models.Movie
}

func (c *CachedMovies) StartCaching() {

	ctx := context.Background()

	go func() {
		timer := time.NewTicker(30 * time.Second)
		defer timer.Stop()

		for {
			select {
			// Актуализировать перечнь фильмов
			case <-timer.C:
				movies := movie.GetPopularMovies()

				c.lock.Lock()
				c.Movies = movies
				c.lock.Unlock()

				log.Println("Список фильмов обновлен")

			// Завершить фоновый процесс кеширования
			case <-ctx.Done():
				break
			}
		}
	}()
}
