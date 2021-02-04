package domain

import (
	"doggos/domain/adapter"
	"doggos/domain/model"
	"doggos/repo"
)

type DoggoUseCase interface {
	GetDoggos(page int, limit int, breedID string) ([]model.Doggo, error)
}

type doggoUseCase struct {
	repository repo.DoggoRepository
	mapper     adapter.DoggosMapper
}

func NewDoggoUseCase(doggoRepo repo.DoggoRepository, doggosMapper adapter.DoggosMapper) DoggoUseCase {
	return &doggoUseCase{
		repository: doggoRepo,
		mapper:     doggosMapper,
	}
}

func (d *doggoUseCase) GetDoggos(page int, limit int, breedID string) ([]model.Doggo, error) {
	type doggosOrError struct {
		doggos []model.Doggo
		err    error
	}
	doggosChannel := make(chan doggosOrError, 1)
	defer close(doggosChannel)
	go func() {
		r, e := d.repository.GetDoggos(page, limit, breedID)
		if e != nil {
			doggosChannel <- doggosOrError{doggos: nil, err: e}
		} else {
			doggosChannel <- doggosOrError{doggos: d.mapper.Map(r), err: nil}
		}
	}()
	r := <-doggosChannel
	return r.doggos, r.err
}
