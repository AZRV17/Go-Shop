package repository

import (
	"context"
	"github.com/AZRV17/Go-Shop/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongo.Collection
}

func NewUserRepo(db *mongo.Collection) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	res, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	user.ID = id

	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) FindById(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	var account domain.User

	if err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *UserRepo) FindAll(ctx context.Context) (*[]domain.User, error) {
	var accounts []domain.User
	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return nil, err
	}

	return &accounts, nil
}
