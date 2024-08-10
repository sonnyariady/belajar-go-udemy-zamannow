package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var barangRepository = &repository.BarangRepositoryMock{Mock: mock.Mock{}}
var barangService = BarangService{Repository: barangRepository}

func TestBarangService_GetNotFound(t *testing.T) {
	//program mock
	barangRepository.Mock.On("FindById", "1").Return(nil)

	barang, err := barangService.Get("1")
	assert.Nil(t, barang)
	assert.NotNil(t, err)
}

func TestBarangService_GetFound(t *testing.T) {
	//program mock

	barang := entity.Barang{
		Kode:  "2",
		Nama:  "Meja",
		Harga: 2000,
	}

	barangRepository.Mock.On("FindById", "2").Return(barang)

	result, err := barangService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, barang.Kode, result.Kode)
	assert.Equal(t, barang.Nama, result.Nama)
	assert.Equal(t, barang.Harga, result.Harga)
}
