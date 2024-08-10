package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type BarangRepositoryMock struct {
	Mock mock.Mock
}

func (repository *BarangRepositoryMock) FindById(id string) *entity.Barang {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		barang := arguments.Get(0).(entity.Barang)
		return &barang
	}
}
