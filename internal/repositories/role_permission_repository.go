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

type RolePermissionRepository struct {
	collection *mongo.Collection
}

func NewRolePermissionRepository(db *mongo.Database) *RolePermissionRepository {
	return &RolePermissionRepository{collection: db.Collection("role_permissions")}
}
func (r *RolePermissionRepository) FindByID(ctx context.Context, id string) (*models.RolePermission, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": objectID,
			},
		},
		{
			"$addFields": bson.M{
				"role_id_obj":       bson.M{"$toObjectId": "$role_id"},
				"permission_id_obj": bson.M{"$toObjectId": "$permission_id"},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "roles",
				"localField":   "role_id_obj",
				"foreignField": "_id",
				"as":           "role",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "permissions",
				"localField":   "permission_id_obj",
				"foreignField": "_id",
				"as":           "permission",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$role",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$permission",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"_id":           1,
				"role_id":       1,
				"permission_id": 1,
				"role":          1,
				"permission":    1,
				"created_at":    1,
				"updated_at":    1,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.RolePermission
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func (r *RolePermissionRepository) FindAll(ctx context.Context, filter *models.ListRolePermissionRequest) (*models.ListRolePermissionResponse, error) {
	matchStage := bson.M{}
	if filter.RoleId != "" {
		matchStage["role_id"] = filter.RoleId
	}
	if filter.PermissionId != "" {
		matchStage["permission_id"] = filter.PermissionId
	}

	pipeline := []bson.M{
		{
			"$match": matchStage,
		},
		{
			"$addFields": bson.M{
				"role_id_obj":       bson.M{"$toObjectId": "$role_id"},
				"permission_id_obj": bson.M{"$toObjectId": "$permission_id"},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "roles",
				"localField":   "role_id_obj",
				"foreignField": "_id",
				"as":           "role",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "permissions",
				"localField":   "permission_id_obj",
				"foreignField": "_id",
				"as":           "permission",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$role",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$permission",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}

	// Add pagination
	if filter.Limit > 0 {
		pipeline = append(pipeline, bson.M{"$skip": filter.Offset})
		pipeline = append(pipeline, bson.M{"$limit": filter.Limit})
	}

	// Get total count
	total, err := r.collection.CountDocuments(ctx, matchStage)
	if err != nil {
		return nil, err
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rolePermissions []*models.RolePermission
	if err = cursor.All(ctx, &rolePermissions); err != nil {
		return nil, err
	}

	response := &models.ListRolePermissionResponse{
		Data:  rolePermissions,
		Count: int(total),
	}

	return response, nil
}
func (r *RolePermissionRepository) Create(ctx context.Context, rolePermission *models.RolePermission) error {
	rolePermission.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, rolePermission)
	return err
}

func (r *RolePermissionRepository) Update(ctx context.Context, rolePermission *models.UpdateRolePermission) error {
	objectID, err := primitive.ObjectIDFromHex(rolePermission.ID)
	if err != nil {
		return err
	}

	rolePermissionMap := bson.M{
		"role_id":       rolePermission.RoleId,
		"permission_id": rolePermission.PermissionId,
		"updated_at":    time.Now(),
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": rolePermissionMap},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", rolePermission.ID)
	}

	return nil
}

func (r *RolePermissionRepository) Delete(ctx context.Context, id string) error {
    // Validate ID length
    if len(id) != 24 {
        return fmt.Errorf("invalid ID format: length should be 24 characters")
    }

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid ID format: %v", err)
    }

    result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return fmt.Errorf("no role-permission found with ID: %s", id)
    }

    return nil
}
