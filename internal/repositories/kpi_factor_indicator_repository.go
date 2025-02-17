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

type KPIFactorIndicatorRepository struct {
	collection *mongo.Collection
}

func NewKPIFactorIndicatorRepository(db *mongo.Database) *KPIFactorIndicatorRepository {
	return &KPIFactorIndicatorRepository{collection: db.Collection("kpi_factor_indicators")}
}

func (r *KPIFactorIndicatorRepository) FindByID(ctx context.Context, id string) (*models.KPIFactorIndicator, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var indicator models.KPIFactorIndicator
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&indicator)
	if err != nil {
		return nil, err
	}

	return &indicator, nil
}

func (r *KPIFactorIndicatorRepository) FindAll(ctx context.Context, filter *models.ListKPIFactorIndicatorRequest) (*models.ListKPIFactorIndicatorResponse, error) {
	findOptions := options.Find()
	filterQuery := bson.M{}

	if filter.MultiSearch != "" {
		filterQuery["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
		}
	}

	if filter.FactorID != "" {
		filterQuery["factor_id"] = filter.FactorID
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

	var indicators []*models.KPIFactorIndicator
	if err = cursor.All(ctx, &indicators); err != nil {
		return nil, err
	}

	return &models.ListKPIFactorIndicatorResponse{
		Items: indicators,
		Count: int(total),
	}, nil
}

func (r *KPIFactorIndicatorRepository) Create(ctx context.Context, indicator *models.KPIFactorIndicator) error {
	indicator.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, indicator)
	return err
}

func (r *KPIFactorIndicatorRepository) Update(ctx context.Context, indicator *models.KPIFactorIndicator) error {
	objectID, err := primitive.ObjectIDFromHex(indicator.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if indicator.FactorID != "" {
		updateFields["factor_id"] = indicator.FactorID
	}
	if indicator.NameEn != "" {
		updateFields["name_en"] = indicator.NameEn
	}
	if indicator.NameKr != "" {
		updateFields["name_kr"] = indicator.NameKr
	}
	if indicator.ProgressRange != "" {
		updateFields["progress_range"] = indicator.ProgressRange
	}
	if indicator.DescriptionEn != "" {
		updateFields["description_en"] = indicator.DescriptionEn
	}
	if indicator.DescriptionKr != "" {
		updateFields["description_kr"] = indicator.DescriptionKr
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
		return fmt.Errorf("no document found with ID: %s", indicator.ID)
	}

	return nil
}

func (r *KPIFactorIndicatorRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}