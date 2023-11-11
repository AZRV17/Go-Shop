package repository

import (
	"context"
	"github.com/AZRV17/Go-Shop/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoodRepo struct {
	db *mongo.Collection
}

func NewGoodRepo(db *mongo.Collection) *GoodRepo {
	return &GoodRepo{
		db: db,
	}
}

func (r *GoodRepo) Create(ctx context.Context, good *domain.Good) (*domain.Good, error) {
	res, err := r.db.InsertOne(ctx, good)
	if err != nil {
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	good.ID = id

	return good, nil
}

func (r *GoodRepo) Update(ctx context.Context, good *domain.Good) (*domain.Good, error) {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": good.ID}, bson.M{"$set": good})
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (r *GoodRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (r *GoodRepo) FindById(ctx context.Context, id primitive.ObjectID) (*domain.Good, error) {
	var good domain.Good

	if err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&good); err != nil {
		return nil, err
	}

	return &good, nil
}

func (r *GoodRepo) FindAll(ctx context.Context) (*[]domain.Good, error) {
	var goods []domain.Good

	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &goods); err != nil {
		return nil, err
	}

	return &goods, nil
}
