package core

type CVrepository interface {
	GetAll() ([]Me, error)
	GetById(id string) (*Me, error)
	Create(m *Me) error
	Update(id string, m *Me) error
	Delete(id string) error
	DeleteEducation(userId string, eduId string) error
	AddEducation(userId string, edu *Education) error
	UpdateEducation(userId string, eduId string, edu *Education) error

	AddWorkExp(userId string, work *WorkExperience) error
	UpdateWorkExp(userId string, workId string, work *WorkExperience) error
	DeleteWorkExp(userId string, workId string) error

	AddProject(userId string, workId string, proj *Project) error
	UpdateProject(userId string, workId string, projId string, proj *Project) error
	DeleteProject(userId string, workId string, projId string) error

	PatchEducation(userId string, eduId string, eduData map[string]interface{}) error
	PatchWorkExp(userId string, workId string, workData map[string]interface{}) error
	PatchProject(userId string, workId string, projId string, projData map[string]interface{}) error
}
