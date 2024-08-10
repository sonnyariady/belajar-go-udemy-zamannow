package repository

import "belajar-golang-unit-test/entity"

type BarangRepository interface {
	FindById(id string) *entity.Barang
}
