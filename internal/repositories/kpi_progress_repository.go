package repositories

import (
	"context"

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

func (r *KPIProgressRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *KPIProgressRepository) TeamProgress(ctx context.Context, req *models.KPIProgressTeamFilter) (*models.ListKPIProgressResponse, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"team_id": req.TeamId,
				"date":    req.Date,
			},
		},
		{
			"$lookup": bson.M{
				"from": "kpi_parents",
				"let":  bson.M{},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$and": []bson.M{
								{"type": "team"},
								{"year": req.Date[:4]},
							},
						},
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
				},
				"as": "kpi",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$kpi",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"_id":        1,
				"date":       1,
				"team_id":    1,
				"created_at": 1,
				"updated_at": 1,
				"kpi":        1,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []models.KPIProgress
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	res := &models.ListKPIProgressResponse{}

	// Get all progress documents for the given date and team
	progressPipeline := []bson.M{
		{
			"$match": bson.M{
				"team_id": req.TeamId,
				"date":    req.Date,
			},
		},
	}

	progressCursor, err := r.collection.Aggregate(ctx, progressPipeline)
	if err != nil {
		return nil, err
	}
	defer progressCursor.Close(ctx)

	var progresses []models.KPIProgress
	if err = progressCursor.All(ctx, &progresses); err != nil {
		return nil, err
	}

	if len(items) > 0 {
		progress := &items[0]

		// Calculate totals for KPI structure
		for i, division := range progress.Kpi.Divisions {
			divisionTotal := 0.0

			for j, criterion := range division.Criterions {
				criterionTotal := 0.0

				for k, factor := range criterion.Factors {
					factorTotal := 0.0

					// Sum up factor indicators progress_range
					for _, indicator := range factor.FactorIndicators {

						// Filter and update factor indicators
						matchedIndicators := make([]*models.ShortKPIPrgFactorIndicator, 0)
						for _, p := range progresses {
							if factor.ID == p.FactorId {
								if indicator.ID == p.FactorIndicatorId {
									indicator.Ratio = p.Ratio
									matchedIndicators = append(matchedIndicators, indicator)
									factorTotal += float64(p.Ratio)
								}
							}
						}
						progress.Kpi.Divisions[i].Criterions[j].Factors[k].FactorIndicators = matchedIndicators
					}

					// Update factor's total_ratio
					progress.Kpi.Divisions[i].Criterions[j].Factors[k].TotalRatio = factorTotal
					criterionTotal += factorTotal
				}

				// Update criterion's total_ratio
				progress.Kpi.Divisions[i].Criterions[j].TotalRatio = criterionTotal
				divisionTotal += criterionTotal
			}

			// Update division's total_ratio
			progress.Kpi.Divisions[i].TotalRatio = divisionTotal
			progress.Kpi.TotalRatio += divisionTotal
		}

		res.Progress = progress
	}
	return res, nil
}
