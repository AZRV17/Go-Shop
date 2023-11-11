package service

import (
	"context"
	"github.com/AZRV17/Go-Shop/internal/domain"
	"github.com/AZRV17/Go-Shop/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserInput struct {
	Login    string `json:"login" bson:"login"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Users interface {
	SignInByLogin(ctx context.Context, login, password string) (*domain.User, error)
	SignInByEmail(ctx context.Context, email, password string) (*domain.User, error)
	SignUp(ctx context.Context, input *CreateUserInput) (*domain.User, error)
	ChangePassword(ctx context.Context, id primitive.ObjectID, password string) error
	ChangeEmail(ctx context.Context, id primitive.ObjectID, email string) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindById(ctx context.Context, id primitive.ObjectID) (*domain.User, error)
	FindAll(ctx context.Context) (*[]domain.User, error)
}

type CreateGoodInput struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Quantity    int64   `json:"quantity" bson:"quantity"`
}

type UpdateGoodInput struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
	Quantity    int64              `json:"quantity" bson:"quantity"`
}

type Goods interface {
	Create(ctx context.Context, input *CreateGoodInput) (*domain.Good, error)
	Update(ctx context.Context, input *UpdateGoodInput) (*domain.Good, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindById(ctx context.Context, id primitive.ObjectID) (*domain.Good, error)
	FindAll(ctx context.Context) (*[]domain.Good, error)
}

type Service struct {
	User Users
	Good Goods
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.Users),
		Good: NewGoodService(repo.Goods),
	}
}
