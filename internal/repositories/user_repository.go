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

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
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
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users *models.ListUsersResponse
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	user.BeforeUpdate()
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	// Convert user struct to map to avoid issues with bson tags
	userMap := bson.M{
		"email":     user.Email,
		"password":  user.Password,
		"full_name": user.FullName,
		"role_id":   user.RoleId,
		"updated_at": time.Now(),
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
	fmt.Println("here")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
