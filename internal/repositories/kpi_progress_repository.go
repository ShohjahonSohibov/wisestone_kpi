package repositories

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"kpi/internal/models"
)

type KPIProgressRepository struct {
	collection *mongo.Collection
}

func NewKPIProgressRepository(db *mongo.Database) *KPIProgressRepository {
	return &KPIProgressRepository{collection: db.Collection("kpi_progresses")}
}

func (r *KPIProgressRepository) Create(ctx context.Context, progress *models.KPIProgress) error {
	progress.BeforeCreate()
	result, err := r.collection.InsertOne(ctx, progress)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		progress.ID = oid.Hex()
	}
	return nil
}

func (r *KPIProgressRepository) Update(ctx context.Context, progress *models.KPIProgress) error {
	objID, err := primitive.ObjectIDFromHex(progress.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if progress.FactorId != "" {
		updateFields["factor_id"] = progress.FactorId
	}
	if progress.FactorIndicatorId != "" {
		updateFields["factor_indicator_id"] = progress.FactorIndicatorId
	}
	if progress.Ratio != 0 {
		updateFields["ratio"] = progress.Ratio
	}
	if !progress.Date.IsZero() {
		updateFields["date"] = progress.Date
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateFields},
	)

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", progress.ID)
	}

	return err
}

func (r *KPIProgressRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *KPIProgressRepository) GetByID(ctx context.Context, id string) (*models.KPIProgress, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{"_id": objID},
		},
		{
			"$lookup": bson.M{
				"from": "kpi_factors",
				"let":  bson.M{"factor_id": bson.M{"$toString": "$factor_id"}},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$eq": []string{"$_id", "$$factor_id"},
							},
						},
					},
				},
				"as": "factor",
			},
		},
		{
			"$lookup": bson.M{
				"from": "kpi_factor_indicators",
				"let":  bson.M{"indicator_id": bson.M{"$toString": "$factor_indicator_id"}},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$eq": []string{"$_id", "$$indicator_id"},
							},
						},
					},
				},
				"as": "factor_indicator",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$factor",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$factor_indicator",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error executing aggregation: %v", err)
	}
	defer cursor.Close(ctx)

	var result []models.KPIProgress
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error decoding result: %v", err)
	}

	if len(result) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &result[0], nil
}

func (r *KPIProgressRepository) TeamProgress(ctx context.Context, req *models.KPIProgressTeamFilter) (*models.ListKPIProgressResponse, error) {

	pipeline := []bson.M{
		{
			"$match": bson.M{"team_id": req.TeamId},
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
			"$group": bson.M{
				"_id":             "$_id",
				"name_en":         bson.M{"$first": "$name_en"},
				"name_kr":         bson.M{"$first": "$name_kr"},
				"description_en":  bson.M{"$first": "$description_en"},
				"description_kr":  bson.M{"$first": "$description_kr"},
				"year":            bson.M{"$first": "$year"},
				"type":            bson.M{"$first": "$type"},
				"status":          bson.M{"$first": "$status"},
				"rejection_count": bson.M{"$first": "$rejection_count"},
				"created_at":      bson.M{"$first": "$created_at"},
				"updated_at":      bson.M{"$first": "$updated_at"},
				"kpi":             bson.M{"$first": "$divisions"},
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
				"kpi":             1,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*models.KPIProgress
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &models.ListKPIProgressResponse{
		Count: len(items),
		Items: items,
	}, nil
}
