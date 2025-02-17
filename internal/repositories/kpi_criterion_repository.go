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

type KPICriterionRepository struct {
	collection *mongo.Collection
}

func NewKPICriterionRepository(db *mongo.Database) *KPICriterionRepository {
	return &KPICriterionRepository{collection: db.Collection("kpi_criterions")}
}
func (r *KPICriterionRepository) FindByID(ctx context.Context, id string) (*models.KPICriterion, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var criterion models.KPICriterion
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&criterion)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, err
	}

	return &criterion, nil
}
func (r *KPICriterionRepository) FindAll(ctx context.Context, filter *models.ListKPICriterionRequest) (*models.ListKPICriterionResponse, error) {
	findOptions := options.Find()

	// Set sort order
	sortOrder := 1
	if filter.SortOrder == "desc" {
		sortOrder = -1
	}
	findOptions.SetSort(bson.D{{Key: "_id", Value: sortOrder}})

	// Build filter
	filterQuery := bson.M{}
	if filter.MultiSearch != "" {
		filterQuery["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
		}
	}
	if filter.DivisionID != "" {
		filterQuery["division_id"] = filter.DivisionID
	}

	// Execute find query
	cursor, err := r.collection.Find(ctx, filterQuery, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var criterions []*models.KPICriterion
	if err := cursor.All(ctx, &criterions); err != nil {
		return nil, err
	}

	return &models.ListKPICriterionResponse{
		Items: criterions,
		Count: len(criterions),
	}, nil
}
func (r *KPICriterionRepository) Create(ctx context.Context, criterion *models.KPICriterion) error {
	criterion.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, criterion)
	return err
}

func (r *KPICriterionRepository) Update(ctx context.Context, criterion *models.KPICriterion) error {
	objectID, err := primitive.ObjectIDFromHex(criterion.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if criterion.DivisionID != "" {
		updateFields["division_id"] = criterion.DivisionID
	}
	if criterion.NameEn != "" {
		updateFields["name_en"] = criterion.NameEn
	}
	if criterion.NameKr != "" {
		updateFields["name_kr"] = criterion.NameKr
	}
	if criterion.TotalRatio != 0 {
		updateFields["total_ratio"] = criterion.TotalRatio
	}
	if criterion.DescriptionEn != "" {
		updateFields["description_en"] = criterion.DescriptionEn
	}
	if criterion.DescriptionKr != "" {
		updateFields["description_kr"] = criterion.DescriptionKr
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
		return fmt.Errorf("no document found with ID: %s", criterion.ID)
	}

	return nil
}

func (r *KPICriterionRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
