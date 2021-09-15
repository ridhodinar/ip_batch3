package model

type Genre struct {
	BaseModel
	Name string `faker:"word,unique"`
}
