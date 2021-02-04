package model

type Doggo struct {
	Breeds []Breed `json:"breeds"`
	Height int     `json:"height"`
	ID     string  `json:"id"`
	URL    string  `json:"url"`
	Width  int     `json:"width"`
}
