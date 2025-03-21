package repositories

import (
	"context"
	"fmt"
	"time"

	"kpi/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type KPIDivisionRepository struct {
	collection *mongo.Collection
}

func NewKPIDivisionRepository(db *mongo.Database) *KPIDivisionRepository {
	return &KPIDivisionRepository{collection: db.Collection("kpi_divisions")}
}

func (r *KPIDivisionRepository) FindByID(ctx context.Context, id string) (*models.KPIDivision, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var division models.KPIDivision
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&division)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, err
	}

	return &division, nil
}

func (r *KPIDivisionRepository) FindAll(ctx context.Context, filter *models.ListKPIDivisionRequest) (*models.ListKPIDivisionResponse, error) {
	pipeline := []bson.M{}

	matchStage := bson.M{}
	if filter.MultiSearch != "" {
		matchStage["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
		}
	}
	if filter.ParentID != "" {
		matchStage["parent_id"] = filter.ParentID
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	pipeline = append(pipeline, []bson.M{
		{
			"$addFields": bson.M{
				"parent_id_obj": bson.M{
					"$convert": bson.M{
						"input":   "$parent_id",
						"to":      "objectId",
						"onError": nil,
					},
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "kpi_parents",
				"localField":   "parent_id_obj",
				"foreignField": "_id",
				"as":           "parent",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$parent",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"parent_id_obj": 0,
			},
		},
	}...)

	sortOrder := 1
	if filter.SortOrder == "desc" {
		sortOrder = -1
	}
	pipeline = append(pipeline, bson.M{
		"$sort": bson.M{"_id": sortOrder},
	})

	if filter.Limit > 0 {
		pipeline = append(pipeline, bson.M{"$skip": filter.Offset})
		pipeline = append(pipeline, bson.M{"$limit": filter.Limit})
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var divisions []*models.KPIDivision
	if err := cursor.All(ctx, &divisions); err != nil {
		return nil, err
	}

	return &models.ListKPIDivisionResponse{
		Items: divisions,
		Count: len(divisions),
	}, nil
}

func (r *KPIDivisionRepository) Create(ctx context.Context, division *models.KPIDivision) error {
	division.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, division)
	return err
}

func (r *KPIDivisionRepository) Update(ctx context.Context, division *models.KPIDivision) error {
	objectID, err := primitive.ObjectIDFromHex(division.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if division.ParentID != "" {
		updateFields["parent_id"] = division.ParentID
	}
	if division.NameEn != "" {
		updateFields["name_en"] = division.NameEn
	}
	if division.NameKr != "" {
		updateFields["name_kr"] = division.NameKr
	}
	if division.DescriptionEn != "" {
		updateFields["description_en"] = division.DescriptionEn
	}
	if division.DescriptionKr != "" {
		updateFields["description_kr"] = division.DescriptionKr
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", division.ID)
	}

	return nil
}

func (r *KPIDivisionRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
