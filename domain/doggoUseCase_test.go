package domain_test

import (
	adapterReal "doggos/domain/adapter"
	adapter "doggos/domain/adapter/mocks"

	usecase "doggos/domain"
	repo "doggos/repo/mocks"
	util "doggos/testutil"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	mockRepo := new(repo.DoggoRepository)
	mockRepo.On("GetDoggos", 1, 25, "").Return(nil, errors.New(""))

	mockMapper := new(adapter.DoggosMapper)

	useCase := usecase.NewDoggoUseCase(mockRepo, mockMapper)
	res, err := useCase.GetDoggos(1, 25, "")

	assert.NotNil(t, err)
	mockMapper.AssertNotCalled(t, "Map", res)
}

func Test_successful(t *testing.T) {
	limit := 25
	dto := util.GenerateMockedDogsDto(limit)
	mockRepo := new(repo.DoggoRepository)
	mockRepo.On("GetDoggos", 1, limit, "").Return(dto, nil)

	mapper := adapterReal.NewDoggoMapper()

	useCase := usecase.NewDoggoUseCase(mockRepo, mapper)
	res, err := useCase.GetDoggos(1, limit, "")

	assert.Nil(t, err)
	assert.Equal(t, res, mapper.Map(dto))
}
