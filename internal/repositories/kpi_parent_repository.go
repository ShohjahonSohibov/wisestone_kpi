package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"kpi/internal/models"
)

type KpiParentRepository struct {
	collection *mongo.Collection
}

func NewKPIParentRepository(db *mongo.Database) *KpiParentRepository {
	return &KpiParentRepository{collection: db.Collection("kpi_parents")}
}

func (r *KpiParentRepository) Create(ctx context.Context, kpiParent *models.KPIParent) error {
	kpiParent.BeforeCreate()

	result, err := r.collection.InsertOne(ctx, kpiParent)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		kpiParent.ID = oid.Hex()
	}

	return nil
}

func (r *KpiParentRepository) Update(ctx context.Context, kpiParent *models.KPIParent) error {
	objID, err := primitive.ObjectIDFromHex(kpiParent.ID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name_en":        kpiParent.NameEn,
			"name_kr":        kpiParent.NameKr,
			"description_en": kpiParent.DescriptionEn,
			"description_kr": kpiParent.DescriptionKr,
			"updated_at":     time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *KpiParentRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *KpiParentRepository) GetByID(ctx context.Context, id string) (*models.KPIParent, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var kpiParent models.KPIParent
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&kpiParent)
	if err != nil {
		return nil, err
	}

	return &kpiParent, nil
}

func (r *KpiParentRepository) List(ctx context.Context, req *models.ListKPIParentRequest) (*models.ListKPIParentResponse, error) {
	filter := bson.M{}
	if req.MultiSearch != "" {
		filter["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": req.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": req.MultiSearch, "$options": "i"}},
		}
	}

	opts := options.Find()

	opts.SetSkip(int64(req.Offset))
	opts.SetLimit(int64(req.Limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*models.KPIParent
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListKPIParentResponse{
		Count: int(count),
		Items: items,
	}, nil
}
