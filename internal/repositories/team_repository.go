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

type TeamRepository struct {
	collection *mongo.Collection
}

func NewTeamRepository(db *mongo.Database) *TeamRepository {
	return &TeamRepository{collection: db.Collection("teams")}
}

func (r *TeamRepository) FindByID(ctx context.Context, id string) (*models.Team, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var team models.Team
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&team)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepository) FindAll(ctx context.Context, filter *models.ListTeamsRequest) (*models.ListTeamsResponse, error) {
	findOptions := options.Find()
	filterQuery := bson.M{}

	// Add search functionality for names
	if filter.MultiSearch != "" {
		filterQuery["$or"] = []bson.M{
			{"name_en": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
			{"name_kr": bson.M{"$regex": filter.MultiSearch, "$options": "i"}},
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

	// Get total count
	total, err := r.collection.CountDocuments(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	// Execute the query with pagination
	cursor, err := r.collection.Find(ctx, filterQuery, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teams []*models.Team
	if err = cursor.All(ctx, &teams); err != nil {
		return nil, err
	}

	response := &models.ListTeamsResponse{
		Items: teams,
		Count: int(total),
	}

	return response, nil
}

func (r *TeamRepository) Create(ctx context.Context, team *models.Team) error {
	team.BeforeCreate()
	_, err := r.collection.InsertOne(ctx, team)
	return err
}

func (r *TeamRepository) Update(ctx context.Context, team *models.Team) error {
	objectID, err := primitive.ObjectIDFromHex(team.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if team.NameEn != "" {
		updateFields["name_en"] = team.NameEn
	}
	if team.NameKr != "" {
		updateFields["name_kr"] = team.NameKr
	}
	if team.DescriptionEn != "" {
		updateFields["description_en"] = team.DescriptionEn
	}
	if team.DescriptionKr != "" {
		updateFields["description_kr"] = team.DescriptionKr
	}
	if team.LeaderId != "" {
		updateFields["leader_id"] = team.LeaderId
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
		return fmt.Errorf("no document found with ID: %s", team.ID)
	}

	return nil
}

func (r *TeamRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
