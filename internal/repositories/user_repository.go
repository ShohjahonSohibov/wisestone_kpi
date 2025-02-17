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

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": objectID,
			},
		},
		{
			"$addFields": bson.M{
				"role_id_obj": bson.M{
					"$convert": bson.M{
						"input":   "$role_id",
						"to":      "objectId",
						"onError": nil,
					},
				},
				"team_id_obj": bson.M{
					"$convert": bson.M{
						"input":   "$team_id",
						"to":      "objectId",
						"onError": nil,
					},
				},
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
			"$unwind": bson.M{
				"path":                       "$role",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "teams",
				"localField":   "team_id_obj",
				"foreignField": "_id",
				"as":           "team",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$team",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"role_id":     0,
				"team_id":     0,
				"role_id_obj": 0,
				"team_id_obj": 0,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.User
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func (r *UserRepository) FindAll(ctx context.Context, filter *models.ListUsersRequest) (*models.ListUsersResponse, error) {
	pipeline := []bson.M{}

	// Add search functionality
	if filter.MultiSearch != "" {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				"$or": []bson.M{
					{"full_name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
					{"full_name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
					{"email": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
					{"position": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
				},
			},
		})
	}

	if filter.TeamId != "" {
		fmt.Println("filter:", filter)
		teamObjectID, err := primitive.ObjectIDFromHex(filter.TeamId)
		if err != nil {
			return nil, err
		}
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				"team_id": teamObjectID.Hex(),
			},
		})
	}
	// Convert role_id to ObjectId (if applicable)
	pipeline = append(pipeline, bson.M{
		"$addFields": bson.M{
			"role_id_obj": bson.M{
				"$convert": bson.M{
					"input":   "$role_id",
					"to":      "objectId",
					"onError": nil, // Prevents failure if role_id is not valid
				},
			},
			"team_id_obj": bson.M{
				"$convert": bson.M{
					"input":   "$team_id",
					"to":      "objectId",
					"onError": nil, // Prevents failure if role_id is not valid
				},
			},
		},
	})

	// Lookup from roles collection
	pipeline = append(pipeline, bson.M{
		"$lookup": bson.M{
			"from":         "roles",
			"localField":   "role_id_obj",
			"foreignField": "_id",
			"as":           "role",
		},
	})

	// Unwind role array
	pipeline = append(pipeline, bson.M{
		"$unwind": bson.M{
			"path":                       "$role",
			"preserveNullAndEmptyArrays": true,
		},
	})

	// Lookup from roles collection
	pipeline = append(pipeline, bson.M{
		"$lookup": bson.M{
			"from":         "teams",
			"localField":   "team_id_obj",
			"foreignField": "_id",
			"as":           "team",
		},
	})

	// Unwind role array
	pipeline = append(pipeline, bson.M{
		"$unwind": bson.M{
			"path":                       "$team",
			"preserveNullAndEmptyArrays": true,
		},
	})

	// Project stage to exclude role_id and team_id
	pipeline = append(pipeline, bson.M{
		"$project": bson.M{
			"role_id":     0,
			"team_id":     0,
			"role_id_obj": 0,
			"team_id_obj": 0,
		},
	})

	// Add sorting
	sortOrder := 1
	if filter.SortOrder == "desc" {
		sortOrder = -1
	}
	pipeline = append(pipeline, bson.M{
		"$sort": bson.M{"_id": sortOrder},
	})

	// Count total documents with the filter
	countStage := bson.M{
		"$count": "total",
	}

	// Apply pagination
	paginationStage := bson.M{
		"$facet": bson.M{
			"metadata": []bson.M{countStage},
			"data": []bson.M{
				{"$skip": filter.Offset},
				{"$limit": filter.Limit},
			},
		},
	}
	pipeline = append(pipeline, paginationStage)
	fmt.Println(pipeline)
	// Execute aggregation
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	// Extract data and count
	var users []*models.User
	var total int
	if len(result) > 0 {
		if metadata, ok := result[0]["metadata"].([]interface{}); ok && len(metadata) > 0 {
			if metadataMap, ok := metadata[0].(primitive.M); ok {
				total = int(metadataMap["total"].(int32))
			}
		}
		if data, ok := result[0]["data"].(primitive.A); ok {
			for _, item := range data {
				if userDoc, ok := item.(primitive.M); ok {
					var user models.User
					userBytes, err := bson.Marshal(userDoc)
					if err == nil {
						if err := bson.Unmarshal(userBytes, &user); err == nil {
							users = append(users, &user)
						}
					}
				}
			}
		}
	}

	fmt.Println("len(result):", users, total)

	response := &models.ListUsersResponse{
		Items: users,
		Count: len(users),
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

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if user.Email != "" {
		updateFields["email"] = user.Email
	}
	if user.Password != "" {
		updateFields["password"] = user.Password
	}
	if user.FullNameEn != "" {
		updateFields["full_name_en"] = user.FullNameEn
	}
	if user.FullNameKr != "" {
		updateFields["full_name_kr"] = user.FullNameKr
	}
	if user.RoleId != "" {
		updateFields["role_id"] = user.RoleId
	}

	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
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

func (r *UserRepository) AssignTeam(ctx context.Context, userID, teamID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"team_id":    teamID,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found with ID: %s", userID)
	}
	return nil
}

func (r *UserRepository) RemoveFromTeam(ctx context.Context, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$unset": bson.M{"team_id": ""},
		"$set":   bson.M{"updated_at": time.Now()},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found with ID: %s", userID)
	}
	return nil
}
