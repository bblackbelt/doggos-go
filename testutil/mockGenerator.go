package testutil

import (
	domainModel "doggos/domain/model"
	repoModel "doggos/repo/model"
	"fmt"
)

func GenerateMockedDogsDto(n int) []repoModel.DoggoDto {
	retVal := make([]repoModel.DoggoDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, repoModel.DoggoDto{
			Height: i,
			ID:     fmt.Sprintf("Id %d", i),
			Width:  i,
			URL:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedDogsBreedDto(n),
		})
	}
	return retVal
}

func GenerateMockedDogsBreeds(n int) []domainModel.Breed {
	retVal := make([]domainModel.Breed, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, domainModel.Breed{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateMockedDogsBreedDto(n int) []repoModel.BreedDto {
	retVal := make([]repoModel.BreedDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, repoModel.BreedDto{
			ID:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateDoggos(n int) []domainModel.Doggo {
	retVal := make([]domainModel.Doggo, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, domainModel.Doggo{
			Height: i,
			ID:     fmt.Sprintf("Id %d", i),
			Width:  i,
			URL:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedDogsBreeds(n),
		})
	}
	return retVal
}
