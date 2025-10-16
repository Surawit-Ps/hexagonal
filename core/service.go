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

func (s *Service) AddEducation(userId string, edu *Education) error {
	return s.repo.AddEducation(userId, edu)
}

func (s *Service) UpdateEducation(userId string, eduId string, edu *Education) error {
	return s.repo.UpdateEducation(userId, eduId, edu)
}


func (s *Service) AddWorkExp(userId string, work *WorkExperience) error {
	return s.repo.AddWorkExp(userId, work)
}

func (s *Service) UpdateWorkExp(userId string, workId string, work *WorkExperience) error {
	return s.repo.UpdateWorkExp(userId, workId, work)
}

func (s *Service) DeleteWorkExp(userId string, workId string) error {
	return s.repo.DeleteWorkExp(userId, workId)
}

func (s *Service) AddProject(userId string, workId string, proj *Project) error {
	return s.repo.AddProject(userId, workId, proj)
}

func (s *Service) UpdateProject(userId string, workId string, projId string, proj *Project) error {
	return s.repo.UpdateProject(userId, workId, projId, proj)
}

func (s *Service) DeleteProject(userId string, workId string, projId string) error {
	return s.repo.DeleteProject(userId, workId, projId)
}

