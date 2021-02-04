package model

type DoggoDto struct {
	Breeds []BreedDto `json:"breeds"`
	ID     string     `json:"id"`
	URL    string     `json:"url"`
	Width  int        `json:"width"`
	Height int        `json:"height"`
}
