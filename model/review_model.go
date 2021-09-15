package model

type Review struct {
	BaseModel
	Review  string `faker:"paragraph"`
	Rate    uint   `faker:"boundary_start=0, boundary_end=10"`
	UserID  uint   `faker:"boundary_start=0, boundary_end=100"`
	MovieID uint   `faker:"boundary_start=0, boundary_end=200"`
}
