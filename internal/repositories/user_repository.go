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

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context, filter *models.ListUsersRequest) (*models.ListUsersResponse, error) {
	findOptions := options.Find()
	filterQuery := bson.M{}

	// Add search functionality
	if filter.MultiSearch != "" {
		filterQuery["$or"] = []bson.M{
			{"full_name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"full_name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"email": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"position": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
		}
	}

	// Add sorting functionality
	if filter.SortOrder == "asc" {
		findOptions.SetSort(bson.M{"_id": -1})
	} else if filter.SortOrder == "desc" {
		findOptions.SetSort(bson.M{"_id": 1})
	}

	// Apply pagination
	if filter.Limit > 0 {
		findOptions.SetLimit(int64(filter.Limit))
		findOptions.SetSkip(int64(filter.Offset))
	}

	// Get total count with filter
	total, err := r.collection.CountDocuments(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	// Execute the query with filter and pagination
	cursor, err := r.collection.Find(ctx, filterQuery, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	response := &models.ListUsersResponse{
		Users: users,
		Count: int(total),
	}

	return response, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	// Convert user struct to map to avoid issues with bson tags
	userMap := bson.M{
		"email":        user.Email,
		"password":     user.Password,
		"full_name_en": user.FullNameEn,
		"full_name_kr": user.FullNameKr,
		"role_id":      user.RoleId,
		"updated_at":   time.Now(),
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": userMap},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID: %s", user.ID)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
