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

type PermissionRepository struct {
	collection *mongo.Collection
}

func NewPermissionRepository(db *mongo.Database) *PermissionRepository {
	return &PermissionRepository{collection: db.Collection("permissions")}
}

func (r *PermissionRepository) FindByID(ctx context.Context, id string) (*models.Permission, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var permission models.Permission
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepository) FindByAction(ctx context.Context, actionKr string) (*models.Permission, error) {
	var permission models.Permission
	err := r.collection.FindOne(ctx, bson.M{"action_kr": actionKr}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepository) FindAll(ctx context.Context, filter *models.ListPermissionRequest) (*models.ListPermissionResponse, error) {
	findOptions := options.Find()
	filterQuery := bson.M{}

	if filter.MultiSearch != "" {
		filterQuery["$or"] = []bson.M{
			{"action_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"action_ru": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"description_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"description_ru": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
		}
	}

	if filter.SortOrder == "asc" {
		findOptions.SetSort(bson.M{"_id": -1})
	} else if filter.SortOrder == "desc" {
		findOptions.SetSort(bson.M{"_id": 1})
	}

	if filter.Limit > 0 {
		findOptions.SetLimit(int64(filter.Limit))
		findOptions.SetSkip(int64(filter.Offset))
	}

	total, err := r.collection.CountDocuments(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	cursor, err := r.collection.Find(ctx, filterQuery, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var permissions []*models.Permission
	if err = cursor.All(ctx, &permissions); err != nil {
		return nil, err
	}

	response := &models.ListPermissionResponse{
		Permissions: permissions,
		Count:       int(total),
	}

	return response, nil
}

func (r *PermissionRepository) Create(ctx context.Context, permission *models.Permission) error {
	permission.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, permission)
	if err != nil {
		return err
	}
	return nil
}

func (r *PermissionRepository) Update(ctx context.Context, permission *models.Permission) error {
	objectID, err := primitive.ObjectIDFromHex(permission.ID)
	if err != nil {
		return err
	}

	permissionMap := bson.M{
		"action_kr":      permission.ActionKr,
		"action_ru":      permission.ActionRu,
		"description_kr": permission.DescriptionKr,
		"description_ru": permission.DescriptionRu,
		"updated_at":     time.Now(),
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": permissionMap},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", permission.ID)
	}

	return nil
}

func (r *PermissionRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
