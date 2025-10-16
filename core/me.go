package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Me struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	NickName    string             `json:"nick_name" bson:"nick_name"`
	Age         int                `json:"age" bson:"age"`
	EducaRecord []Education        `json:"educa_record" bson:"educa_record"`
	WorkExp     []WorkExperience   `json:"work_exp" bson:"work_exp"` //trest
}

type Education struct {
	ID 		primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	School 	string `json:"school" bson:"school"`
	GPA    	float32 `json:"gpa" bson:"gpa"`
	Year   	string    `json:"year" bson:"year"`
}

type WorkExperience struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Company  string `json:"company" bson:"company"`
	Project  []Project`json:"project" bson:"project"`
	Years    string   `json:"years" bson:"years"`
}

type Project struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectName string             `json:"project_name" bson:"project_name"`
	Description string             `json:"description" bson:"description"`
	Link        string             `json:"link" bson:"link"`
}