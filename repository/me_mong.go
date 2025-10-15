package repository

import (
	"context"
	"hexagonal/core"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) core.CVrepository {
	return &MongoRepo{col: db.Collection("Me")}
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
	if err != nil {
		return nil, err
	}
	return &me, nil
}

func (r *MongoRepo) Create(m *core.Me) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m.ID = primitive.NewObjectID() // ✅ สร้าง _id ใหม่อัตโนมัติ
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

	_, err = r.col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": m})
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
