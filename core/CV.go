package core

type CVrepository interface {
	GetAll() ([]Me, error)
	GetById(id string) (*Me, error)
	Create(m *Me) error
	Update(id string, m *Me) error
	Delete(id string) error
}
