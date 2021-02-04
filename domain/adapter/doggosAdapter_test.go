package adapter_test

import (
	"doggos/domain/adapter"
	dtoModel "doggos/repo/model"
	util "doggos/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMap(t *testing.T) {
	assert.Empty(t, adapter.NewDoggoMapper().Map(make([]dtoModel.DoggoDto, 0)))
}

func TestMapLen(t *testing.T) {
	n := 5
	values := util.GenerateMockedDogsDto(n)
	assert.Len(t, adapter.NewDoggoMapper().Map(values), n)
}

func TestMapContent(t *testing.T) {
	n := 5
	values := util.GenerateMockedDogsDto(n)
	assert.Equal(t, adapter.NewDoggoMapper().Map(values), util.GenerateDoggos(n))
}
