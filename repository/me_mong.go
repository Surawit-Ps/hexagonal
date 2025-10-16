package repository

import (
	"context"
	"hexagonal/core"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) core.CVrepository {
	return &MongoRepo{col: db.Collection("prismo")}
}

func (r *MongoRepo) GetAll() ([]core.Me, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []core.Me
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *MongoRepo) GetById(id string) (*core.Me, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var me core.Me
	err = r.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&me)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &me, err
}

func (r *MongoRepo) Create(m *core.Me) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ✅ ใส่ ID ให้ Me
	m.ID = primitive.NewObjectID()

	// ✅ ใส่ ID ให้ Education
	for i := range m.EducaRecord {
		m.EducaRecord[i].ID = primitive.NewObjectID()
	}

	// ✅ ใส่ ID ให้ WorkExperience และ Project
	for i := range m.WorkExp {
		m.WorkExp[i].ID = primitive.NewObjectID()
		for j := range m.WorkExp[i].Project {
			m.WorkExp[i].Project[j].ID = primitive.NewObjectID()
		}
	}

	_, err := r.col.InsertOne(ctx, m)
	return err
}


func (r *MongoRepo) Update(id string, m *core.Me) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"name":         m.Name,
		"nick_name":    m.NickName,
		"age":          m.Age,
		"educa_record": m.EducaRecord,
		"work_exp":     m.WorkExp,
	}

	_, err = r.col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	return err
}

func (r *MongoRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.col.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

