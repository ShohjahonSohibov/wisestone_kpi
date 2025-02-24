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

type KPIProgressStatusRepository struct {
	collection *mongo.Collection
}

func NewKPIProgressStatusRepository(db *mongo.Database) *KPIProgressStatusRepository {
	return &KPIProgressStatusRepository{collection: db.Collection("kpi_progress_statuses")}
}

func (r *KPIProgressStatusRepository) Create(ctx context.Context, status *models.KPIProgressStatus) error {
	status.BeforeCreate()
	result, err := r.collection.InsertOne(ctx, status)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		status.ID = oid.Hex()
	}
	return nil
}

func (r *KPIProgressStatusRepository) Update(ctx context.Context, status *models.UpdateKPIProgressStatus) error {
	objID, err := primitive.ObjectIDFromHex(status.ID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"status":     status.Status,
			"updated_at": time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *KPIProgressStatusRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	fmt.Println("objID:", objID)
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *KPIProgressStatusRepository) List(ctx context.Context, req *models.ListKPIProgressStatusRequest) (*models.ListKPIProgressStatusResponse, error) {
	// Define the base match filter
	matchFilter := bson.M{}
	if req.TeamId != "" {
		matchFilter["team_id"] = req.TeamId
	}
	if req.EmployeeId != "" {
		matchFilter["employee_id"] = req.EmployeeId
	}
	if req.Date != "" {
		matchFilter["date"] = req.Date
	}
	if req.Status != "" {
		matchFilter["status"] = req.Status
	}

	// Define the aggregation pipeline
	pipeline := []bson.M{
		{"$match": matchFilter}, // Apply filters if provided
		{
			"$addFields": bson.M{
				"team_id":     bson.M{"$toString": "$team_id"},
				"employee_id": bson.M{"$toString": "$employee_id"},
			},
		},
		{
			"$lookup": bson.M{
				"from": "teams",
				"let":  bson.M{"teamId": "$team_id"},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$expr": bson.M{"$eq": []interface{}{bson.M{"$toString": "$_id"}, "$$teamId"}},
						},
					},
				},
				"as": "team",
			},
		},
		{
			"$lookup": bson.M{
				"from": "users",
				"let":  bson.M{"employeeId": "$employee_id"},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$expr": bson.M{"$eq": []interface{}{bson.M{"$toString": "$_id"}, "$$employeeId"}},
						},
					},
				},
				"as": "employee",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$team",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$employee",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"_id":         1,
				"team_id":     1,
				"employee_id": 1,
				"status":      1,
				"date":        1,
				"created_at":  1,
				"updated_at":  1,
				"team":        1,
				"employee":    1,
			},
		},
	}

	// Apply pagination if needed
	if req.Limit > 0 {
		pipeline = append(pipeline,
			bson.M{"$skip": req.Offset},
			bson.M{"$limit": req.Limit},
		)
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*models.KPIProgressStatus
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return &models.ListKPIProgressStatusResponse{
		Count: len(items),
		Items: items,
	}, nil
}
