package repository

import (
	"context"
	"github.com/AZRV17/Go-Shop/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdatePassword(ctx context.Context, id primitive.ObjectID, password string) error
	UpdateEmail(ctx context.Context, id primitive.ObjectID, email string) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindById(ctx context.Context, id primitive.ObjectID) (*domain.User, error)
	FindByLogin(ctx context.Context, login string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindAll(ctx context.Context) (*[]domain.User, error)
}

type Goods interface {
	Create(ctx context.Context, good *domain.Good) (*domain.Good, error)
	Update(ctx context.Context, good *domain.Good) (*domain.Good, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindById(ctx context.Context, id primitive.ObjectID) (*domain.Good, error)
	FindAll(ctx context.Context) (*[]domain.Good, error)
}

type Repository struct {
	Users Users
	Goods Goods
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Users: NewUserRepo(db.Collection("users")),
		Goods: NewGoodRepo(db.Collection("goods")),
	}
}
