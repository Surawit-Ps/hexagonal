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

func (s *Service) DeleteEducation(userId string, eduId string) error {
	return s.repo.DeleteEducation(userId, eduId)
}

// func (s *Service) AddEducation(userID string, edu Education) error {
// 	return s.repo.AddEducation(userID, edu)
// }
// func (s *Service) UpdateEducation(userID string, eduID string, edu Education) error {
// 	return s.repo.UpdateEducation(userID, eduID, edu)
// }

