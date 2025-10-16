package core

type CVrepository interface {
	GetAll() ([]Me, error)
	GetById(id string) (*Me, error)
	Create(m *Me) error
	Update(id string, m *Me) error
	Delete(id string) error
	DeleteEducation(userId string, eduId string) error

	// AddEducation(userID string, edu Education) error
    // UpdateEducation(userID string, eduID string, edu Education) error


    // CRUD ย่อย Work Experience
    // AddWork(userID string, work WorkExperience) error
    // UpdateWork(userID string, workID string, work WorkExperience) error
    // DeleteWork(userID string, workID string) error

    // CRUD ย่อย Project (ซ้อนใน WorkExp)
    // AddProject(userID string, workID string, project Project) error
    // UpdateProject(userID string, workID string, projectID string, project Project) error
    // DeleteProject(userID string, workID string, projectID string) error
}
