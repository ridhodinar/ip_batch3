package seed

import (
	"review_movie/model"

	"github.com/bxcodec/faker/v3"
)

func (s Seeder) MovieSeed() {
	for i := 0; i < 200; i++ {
		movie := model.Movie{}
		_ = faker.FakeData(&movie)
		s.Db.Create(&movie)
	}
}
