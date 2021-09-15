package seed

import (
	"review_movie/model"

	"github.com/bxcodec/faker/v3"
)

func (s Seeder) GenreSeed() {
	for i := 0; i < 5; i++ {
		genre := model.Genre{}
		_ = faker.FakeData(&genre)
		s.Db.Create(&genre)
	}
}
