package model

type BreedDto struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BreedGroup  string `json:"breed_group"`
	LifeSpan    string `json:"life_span"`
	Temperament string `json:"temperament"`
}
