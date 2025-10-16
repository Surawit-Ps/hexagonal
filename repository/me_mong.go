package repository

import (
	"context"
	"hexagonal/core"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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


	m.ID = primitive.NewObjectID()


	for i := range m.EducaRecord {
		m.EducaRecord[i].ID = primitive.NewObjectID()
	}

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

func (r *MongoRepo) DeleteEducation(userId string, eduId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	eduObjectID, err := primitive.ObjectIDFromHex(eduId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userObjectID}
	update := bson.M{
		"$pull": bson.M{
			"educa_record": bson.M{"_id": eduObjectID},
		},
	}

	_, err = r.col.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoRepo) AddEducation(userId string, edu *core.Education) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	edu.ID = primitive.NewObjectID()

	update := bson.M{
		"$push": bson.M{"educa_record": edu},
	}
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": userObjID}, update)
	return err
}

func (r *MongoRepo) UpdateEducation(userId string, eduId string, edu *core.Education) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	eduObjID, _ := primitive.ObjectIDFromHex(eduId)

	update := bson.M{
		"$set": bson.M{
			"educa_record.$.school": edu.School,
			"educa_record.$.gpa":    edu.GPA,
			"educa_record.$.year":   edu.Year,
		},
	}

	_, err := r.col.UpdateOne(ctx, bson.M{"_id": userObjID, "educa_record._id": eduObjID}, update)
	return err
}

// ---------------- Work Experience ----------------
func (r *MongoRepo) AddWorkExp(userId string, work *core.WorkExperience) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	work.ID = primitive.NewObjectID()

	update := bson.M{
		"$push": bson.M{"work_exp": work},
	}
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": userObjID}, update)
	return err
}

func (r *MongoRepo) UpdateWorkExp(userId string, workId string, work *core.WorkExperience) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	workObjID, _ := primitive.ObjectIDFromHex(workId)

	update := bson.M{
		"$set": bson.M{
			"work_exp.$.company": work.Company,
			"work_exp.$.years":   work.Years,
			"work_exp.$.project": work.Project,
		},
	}
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": userObjID, "work_exp._id": workObjID}, update)
	return err
}

func (r *MongoRepo) DeleteWorkExp(userId string, workId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	workObjID, _ := primitive.ObjectIDFromHex(workId)

	update := bson.M{
		"$pull": bson.M{"work_exp": bson.M{"_id": workObjID}},
	}
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": userObjID}, update)
	return err
}

// ---------------- Project ----------------
func (r *MongoRepo) AddProject(userId string, workId string, proj *core.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	workObjID, _ := primitive.ObjectIDFromHex(workId)
	proj.ID = primitive.NewObjectID()

	update := bson.M{
		"$push": bson.M{"work_exp.$.project": proj},
	}

	_, err := r.col.UpdateOne(ctx, bson.M{"_id": userObjID, "work_exp._id": workObjID}, update)
	return err
}

func (r *MongoRepo) UpdateProject(userId string, workId string, projId string, proj *core.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	workObjID, _ := primitive.ObjectIDFromHex(workId)
	projObjID, _ := primitive.ObjectIDFromHex(projId)

	// ใช้ positional filter ซ้อนสำหรับ array project
	filter := bson.M{
		"_id":                    userObjID,
		"work_exp._id":           workObjID,
		"work_exp.project._id":   projObjID,
	}
	update := bson.M{
		"$set": bson.M{
			"work_exp.$[w].project.$[p].project_name": proj.ProjectName,
			"work_exp.$[w].project.$[p].description":  proj.Description,
			"work_exp.$[w].project.$[p].link":        proj.Link,
		},
	}
	opts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"w._id": workObjID},
			bson.M{"p._id": projObjID},
		},
	})
	_, err := r.col.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *MongoRepo) DeleteProject(userId string, workId string, projId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, _ := primitive.ObjectIDFromHex(userId)
	workObjID, _ := primitive.ObjectIDFromHex(workId)
	projObjID, _ := primitive.ObjectIDFromHex(projId)

	filter := bson.M{"_id": userObjID, "work_exp._id": workObjID}
	update := bson.M{
		"$pull": bson.M{"work_exp.$.project": bson.M{"_id": projObjID}},
	}
	_, err := r.col.UpdateOne(ctx, filter, update)
	return err
}



