package model

type MovieGenre struct {
	BaseModel
	Movie   string `faker:"-"`
	Genre   string `faker:"-"`
	MovieID uint   `faker:"boundary_start=1, boundary_end=200"`
	GenreID uint   `faker:"boundary_start=1, boundary_end=5"`
}
