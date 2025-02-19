package repositories

import (
	"context"
	"fmt"
	"time"

	"kpi/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KPIFactorRepository struct {
	collection *mongo.Collection
}

func NewKPIFactorRepository(db *mongo.Database) *KPIFactorRepository {
	return &KPIFactorRepository{collection: db.Collection("kpi_factors")}
}

func (r *KPIFactorRepository) FindByID(ctx context.Context, id string) (*models.KPIFactor, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var factor models.KPIFactor
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&factor)
	if err != nil {
		return nil, err
	}

	return &factor, nil
}

func (r *KPIFactorRepository) FindAll(ctx context.Context, filter *models.ListKPIFactorRequest) (*models.ListKPIFactorResponse, error) {
	findOptions := options.Find()
	filterQuery := bson.M{}

	if filter.MultiSearch != "" {
		filterQuery["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
		}
	}

	if filter.CriterionID != "" {
		filterQuery["criterion_id"] = filter.CriterionID
	}

	if filter.SortOrder == "asc" {
		findOptions.SetSort(bson.M{"_id": 1})
	} else {
		findOptions.SetSort(bson.M{"_id": -1})
	}

	findOptions.SetSkip(int64(filter.Offset))
	findOptions.SetLimit(int64(filter.Limit))

	total, err := r.collection.CountDocuments(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	cursor, err := r.collection.Find(ctx, filterQuery, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var factors []*models.KPIFactor
	if err = cursor.All(ctx, &factors); err != nil {
		return nil, err
	}

	return &models.ListKPIFactorResponse{
		Items: factors,
		Count: int(total),
	}, nil
}

func (r *KPIFactorRepository) Create(ctx context.Context, factor *models.KPIFactor) error {
	factor.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, factor)
	return err
}

func (r *KPIFactorRepository) Update(ctx context.Context, factor *models.KPIFactor) error {
	objectID, err := primitive.ObjectIDFromHex(factor.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if factor.CriterionID != "" {
		updateFields["criterion_id"] = factor.CriterionID
	}
	if factor.NameEn != "" {
		updateFields["name_en"] = factor.NameEn
	}
	if factor.NameKr != "" {
		updateFields["name_kr"] = factor.NameKr
	}
	if factor.Ratio != 0 {
		updateFields["ratio"] = factor.Ratio
	}
	if factor.DescriptionEn != "" {
		updateFields["description_en"] = factor.DescriptionEn
	}
	if factor.DescriptionKr != "" {
		updateFields["description_kr"] = factor.DescriptionKr
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
		return fmt.Errorf("no document found with ID: %s", factor.ID)
	}

	return nil
}

func (r *KPIFactorRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *KPIFactorRepository) GetSumRatioByCriterionID(ctx context.Context, criterionID string) (float64, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"criterion_id": criterionID,
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"total": bson.M{"$sum": "$ratio"},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, nil
	}

	return result[0]["total"].(float64), nil
}