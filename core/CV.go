package core

type CVrepository interface {
	GetAll() ([]Me, error)
	GetById(id int) (*Me, error)
	Create(m *Me) error
	Update(id int, m *Me) error
	Delete(id int) error
}