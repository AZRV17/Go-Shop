package service

import (
	"context"
	"github.com/AZRV17/Go-Shop/internal/domain"
	"github.com/AZRV17/Go-Shop/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoodService struct {
	repo repository.Goods
}

func NewGoodService(repo repository.Goods) *GoodService {
	return &GoodService{
		repo: repo,
	}
}

func (s GoodService) Create(ctx context.Context, input *CreateGoodInput) (*domain.Good, error) {
	good := &domain.Good{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
	}

	good, err := s.repo.Create(ctx, good)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (s GoodService) Update(ctx context.Context, input *UpdateGoodInput) (*domain.Good, error) {
	good := &domain.Good{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
	}

	good, err := s.repo.Update(ctx, good)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (s GoodService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}

func (s GoodService) FindById(ctx context.Context, id primitive.ObjectID) (*domain.Good, error) {
	good, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (s GoodService) FindAll(ctx context.Context) (*[]domain.Good, error) {
	goods, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return goods, nil
}
