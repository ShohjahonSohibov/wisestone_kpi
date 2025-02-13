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

type RoleRepository struct {
	collection *mongo.Collection
}

func NewRoleRepository(db *mongo.Database) *RoleRepository {
	return &RoleRepository{collection: db.Collection("roles")}
}

func (r *RoleRepository) FindByID(ctx context.Context, id string) (*models.Role, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var role models.Role
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) FindAll(ctx context.Context, filter *models.ListRoleRequest) (*models.ListRoleResponse, error) {
    findOptions := options.Find()
    filterQuery := bson.M{}

    if filter.MultiSearch != "" {
        filterQuery["$or"] = []bson.M{
            {"name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
            {"name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
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

	var roles []*models.Role
	if err = cursor.All(ctx, &roles); err != nil {
		return nil, err
	}

	response := &models.ListRoleResponse{
		Items: roles,
		Count: int(total),
	}

	return response, nil
}

func (r *RoleRepository) Create(ctx context.Context, role *models.Role) error {
	role.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, role)
	return err
}

func (r *RoleRepository) Update(ctx context.Context, role *models.Role) error {
    objectID, err := primitive.ObjectIDFromHex(role.ID)
    if err != nil {
        return err
    }

    roleMap := bson.M{
        "name_en":    role.NameEn,
        "name_kr":    role.NameKr,
        "updated_at": time.Now(),
    }

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": roleMap},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", role.ID)
	}

	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}