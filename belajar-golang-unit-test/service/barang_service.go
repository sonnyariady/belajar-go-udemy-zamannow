package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type BarangService struct {
	Repository repository.BarangRepository
}

func (service BarangService) Get(id string) (*entity.Barang, error) {
	barang := service.Repository.FindById(id)
	if barang == nil {
		return barang, errors.New("Barang tidak ketemu")
	} else {
		return barang, nil
	}

}
