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

type RolePermissionRepository struct {
	collection *mongo.Collection
}

func NewRolePermissionRepository(db *mongo.Database) *RolePermissionRepository {
	return &RolePermissionRepository{collection: db.Collection("role_permissions")}
}

func (r *RolePermissionRepository) FindByID(ctx context.Context, id string) (*models.RolePermission, error) {
	// objectID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, err
	// }

	// pipeline := mongo.Pipeline{
	// 	{
	// 		{"$match": bson.M{"_id": objectID}},
	// 	},
	// 	{
	// 		{"$lookup": bson.M{
	// 			"from":         "roles",
	// 			"localField":   "role_id",
	// 			"foreignField": "_id",
	// 			"as":          "role",
	// 		}},
	// 	},
	// 	{
	// 		{"$lookup": bson.M{
	// 			"from":         "permissions",
	// 			"localField":   "permission_id",
	// 			"foreignField": "_id",
	// 			"as":          "permission",
	// 		}},
	// 	},
	// 	{
	// 		{"$unwind": "$role"},
	// 	},
	// 	{
	// 		{"$unwind": "$permission"},
	// 	},
	// 	{
	// 		{"$project": bson.M{
	// 			"_id":            1,
	// 			"role_id":        1,
	// 			"permission_id":   1,
	// 			"created_at":     1,
	// 			"updated_at":     1,
	// 			"role_name_uz":   "$role.name_uz",
	// 			"role_name_en":   "$role.name_en",
	// 			"role_name_kr":   "$role.name_kr",
	// 			"perm_name_uz":   "$permission.name_uz",
	// 			"perm_name_en":   "$permission.name_en",
	// 			"perm_name_kr":   "$permission.name_kr",
	// 		}},
	// 	},
	// }

	// cursor, err := r.collection.Aggregate(ctx, pipeline)
	// if err != nil {
	// 	return nil, err
	// }
	// defer cursor.Close(ctx)

	// var results []models.RolePermission
	// if err = cursor.All(ctx, &results); err != nil {
	// 	return nil, err
	// }

	// if len(results) == 0 {
	// 	return nil, nil
	// }

	// return &results[0], nil
	return nil, nil
}

func (r *RolePermissionRepository) FindAll(ctx context.Context, filter *models.ListRolePermissionRequest) (*models.ListRolePermissionResponse, error) {
	findOptions := options.Find()
	filterQuery := bson.M{}

	// Add filters for role_id and permission_id if they are provided
	if filter.RoleId != "" {
		filterQuery["role_id"] = filter.RoleId
	}
	if filter.PermissionId != "" {
		filterQuery["permission_id"] = filter.PermissionId
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

func (r *RolePermissionRepository) Update(ctx context.Context, rolePermission *models.RolePermission) error {
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}