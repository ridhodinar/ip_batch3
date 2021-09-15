package seed

import (
	"review_movie/model"

	"github.com/bxcodec/faker/v3"
)

func (s Seeder) UserSeed() {
	for i := 0; i < 100; i++ {
		user := model.User{}
		_ = faker.FakeData(&user)
		s.Db.Create(&user)
	}
}
