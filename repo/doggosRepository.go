package repo

import (
	"doggos/executor"
	"doggos/repo/model"
	"encoding/json"
)

type DoggoRepository interface {
	GetDoggos(page int, limit int, breedID string) ([]model.DoggoDto, error)
}

type doggoRepository struct {
	Executor       executor.Executor
	RequestFactory executor.RequestFactory
}

func NewDoggoRepository(Executor executor.Executor, factory executor.RequestFactory) DoggoRepository {
	return &doggoRepository{
		Executor:       Executor,
		RequestFactory: factory,
	}
}

func (doggoRepository *doggoRepository) GetDoggos(page int, limit int, breedID string) ([]model.DoggoDto, error) {
	req, e := doggoRepository.RequestFactory.Build(page, limit, breedID)
	if e != nil {
		return nil, e
	}
	response, eww := doggoRepository.Executor.Execute(req)
	if eww != nil {
		return nil, eww
	}
	defer response.Body.Close()
	var data []model.DoggoDto
	json.NewDecoder(response.Body).Decode(&data)
	return data, eww
}
