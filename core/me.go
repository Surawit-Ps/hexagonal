package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Me struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	NickName    string             `json:"nick_name" bson:"nick_name"`
	Age         int                `json:"age" bson:"age"`
	EducaRecord []string           `json:"educa_record" bson:"educa_record"`
	WorkExp     []string           `json:"work_exp" bson:"work_exp"` //trest
}