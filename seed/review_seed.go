package seed

import (
	"review_movie/model"

	"github.com/bxcodec/faker/v3"
)

func (s Seeder) ReviewSeed() {
	for i := 0; i < 200; i++ {
		review := model.Review{}
		_ = faker.FakeData(&review)
		s.Db.Create(&review)
	}
}
