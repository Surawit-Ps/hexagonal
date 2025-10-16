package core

type Service struct {
	repo CVrepository
}

func NewService(r CVrepository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAll() ([]Me, error) {
	return s.repo.GetAll()
}

func (s *Service) GetById(id string) (*Me, error) {
	return s.repo.GetById(id)
}

func (s *Service) Create(m *Me) error {
	return s.repo.Create(m)
}

func (s *Service) Update(id string, m *Me) error {
	return s.repo.Update(id, m)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
