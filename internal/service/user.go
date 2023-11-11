package service

import (
	"context"
	"errors"
	"github.com/AZRV17/Go-Shop/internal/domain"
	"github.com/AZRV17/Go-Shop/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) SignInByLogin(ctx context.Context, login, password string) (*domain.User, error) {
	user, err := s.repo.FindByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	if user == nil {
		err = errors.New("user not found")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) SignInByEmail(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		err = errors.New("user not found")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) SignUp(ctx context.Context, input *CreateUserInput) (*domain.User, error) {
	hashPassword, err := s.hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Login:    input.Login,
		Email:    input.Email,
		Password: hashPassword,
	}

	user, err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id primitive.ObjectID, password string) error {
	hashPassword, err := s.hashPassword(password)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(ctx, id, hashPassword)
}

func (s *UserService) ChangeEmail(ctx context.Context, id primitive.ObjectID, email string) error {
	return s.repo.UpdateEmail(ctx, id, email)
}

func (s *UserService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) FindById(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindAll(ctx context.Context) (*[]domain.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}
