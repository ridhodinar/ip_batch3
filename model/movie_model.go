package model

type Movie struct {
	BaseModel
	Ratings uint   `faker:"boundary_start=0, boundary_end=10"`
	Title   string `faker:"sentence"`
	Year    uint   `faker:"boundary_start=1900, boundary_end=2021"`
}
