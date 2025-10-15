package core


type Service struct {
	repo CVrepository
}

func NewService(r CVrepository) *Service {
	return &Service{repo: r}
}
 
func (s *Service) GetAll() ([]Me, error){
	return s.repo.GetAll()
}

func (s *Service) GetById(id int) (*Me, error){
	return s.repo.GetById(id)
}

func (s *Service) Create(m *Me) error {
	return s.repo.Create(m) 
}

func (s *Service) Update(id int, m *Me) error{
	return s.repo.Update(id, m)
}

func (s *Service) Delete(id int) error{
	return s.repo.Delete(id)
}