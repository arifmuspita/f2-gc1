package usecase

import "errors"

var (
	ErrNotFound      = errors.New("error not found")
	ErrCustomerEmpty = errors.New("customer tidak ditemukan")
	ErrIdNotValid    = errors.New("ID tidak valid untuk dihapus")
)
