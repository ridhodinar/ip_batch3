package seed

import (
	"review_movie/model"

	"github.com/bxcodec/faker/v3"
)

func (s Seeder) MovieGenreSeed() {
	for i := 0; i < 100; i++ {
		movieGenre := model.MovieGenre{}
		_ = faker.FakeData(&movieGenre)
		movie := model.Movie{}
		genre := model.Genre{}
		//s.Db.Take(&movie)
		//s.Db.Take(&genre)
		s.Db.First(&movie, movieGenre.MovieID)
		s.Db.First(&genre, movieGenre.GenreID)
		movieGenre.Movie = movie.Title
		movieGenre.Genre = genre.Name
		s.Db.Create(&movieGenre)
	}
}
