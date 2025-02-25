package repositories

import (
	"context"
	"fmt"
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

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if kpiParent.NameKr != "" {
		updateFields["name_kr"] = kpiParent.NameKr
	}
	if kpiParent.NameEn != "" {
		updateFields["name_en"] = kpiParent.NameEn
	}
	if kpiParent.DescriptionKr != "" {
		updateFields["description_kr"] = kpiParent.DescriptionKr
	}
	if kpiParent.DescriptionEn != "" {
		updateFields["description_en"] = kpiParent.DescriptionEn
	}
	if kpiParent.Year != "" {
		updateFields["year"] = kpiParent.Year
	}
	if kpiParent.Type != "" {
		updateFields["type"] = kpiParent.Type
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateFields},
	)

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", kpiParent.ID)
	}

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

func (r *KpiParentRepository) GetByID(ctx context.Context, id, kpiType string) (*models.KPIParent, error) {
	matchStage := bson.M{}
	if id != "" {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid ID format: %v", err)
		}

		matchStage = bson.M{"_id": objID}
	} else if kpiType != "" {
		currentYear := time.Now().Format("2006")
		matchStage["type"] = kpiType
		matchStage["year"] = currentYear
	}

	pipeline := []bson.M{
		{
			"$match": matchStage,
		},
		{
			"$lookup": bson.M{
				"from": "kpi_divisions",
				"let":  bson.M{"parent_id": bson.M{"$toString": "$_id"}},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$eq": []string{"$parent_id", "$$parent_id"},
							},
						},
					},
					{
						"$lookup": bson.M{
							"from": "kpi_criterions",
							"let":  bson.M{"division_id": bson.M{"$toString": "$_id"}},
							"pipeline": []bson.M{
								{
									"$match": bson.M{
										"$expr": bson.M{
											"$eq": []string{"$division_id", "$$division_id"},
										},
									},
								},
								{
									"$lookup": bson.M{
										"from": "kpi_factors",
										"let":  bson.M{"criterion_id": bson.M{"$toString": "$_id"}},
										"pipeline": []bson.M{
											{
												"$match": bson.M{
													"$expr": bson.M{
														"$eq": []string{"$criterion_id", "$$criterion_id"},
													},
												},
											},
											{
												"$lookup": bson.M{
													"from": "kpi_factor_indicators",
													"let":  bson.M{"factor_id": bson.M{"$toString": "$_id"}},
													"pipeline": []bson.M{
														{
															"$match": bson.M{
																"$expr": bson.M{
																	"$eq": []string{"$factor_id", "$$factor_id"},
																},
															},
														},
													},
													"as": "factor_indicators",
												},
											},
										},
										"as": "factors",
									},
								},
							},
							"as": "criterions",
						},
					},
				},
				"as": "divisions",
			},
		},
		{
			"$project": bson.M{
				"_id":             1,
				"name_en":         1,
				"name_kr":         1,
				"description_en":  1,
				"description_kr":  1,
				"year":            1,
				"type":            1,
				"status":          1,
				"rejection_count": 1,
				"created_at":      1,
				"updated_at":      1,
				"divisions":       1,
			},
		},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error executing aggregation: %v", err)
	}
	defer cursor.Close(ctx)

	// Check if cursor has any data
	if !cursor.Next(ctx) {
		return nil, mongo.ErrNoDocuments
	}

	// Decode results one by one to avoid struct mismatches
	var result []models.KPIParent
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error decoding result: %v", err)
	}

	// Debugging: Print the final result
	fmt.Printf("Fetched KPI Parent: %+v\n", result)

	return &result[0], nil
}

func (r *KpiParentRepository) List(ctx context.Context, req *models.ListKPIParentRequest) (*models.ListKPIParentResponse, error) {
	filter := bson.M{}
	if req.MultiSearch != "" {
		filter["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": req.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": req.MultiSearch, "$options": "i"}},
		}
	}

	if req.Year != "" {
		filter["year"] = req.Year
	}
	if req.Status != "" {
		filter["status"] = req.Status
	}
	if req.Type != "" {
		filter["type"] = req.Type
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

func (r *KpiParentRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"status":     status,
		"updated_at": time.Now(),
	}

	if status == string(models.KPIStatusRejected) {
		updateFields["$inc"] = bson.M{"rejection_count": 1}
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateFields},
	)

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", id)
	}

	return err
}
