package services

import (
	"context"
	"errors"
	"tasked/internal/domain"
	apperrors "tasked/internal/errors"
	"tasked/internal/repository"
	"tasked/internal/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, apperrors.ErrNotFound
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if !utils.ValidateEmail(email) {
		return nil, errors.New("Invalid email format")
	}
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) CreateUser(ctx context.Context, username, email, password string) (*domain.User, error) {
	if !utils.ValidateEmail(email) {
		return nil, apperrors.ErrBadRequest
	}
	passwordHashed, err := utils.HashedPassword(password)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}
	return s.repo.CreateUser(ctx, username, email, passwordHashed)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, username, email string) (*domain.User, error) {
	if !utils.ValidateEmail(email) {
		return nil, errors.New("Invalid email format")
	}
	return s.repo.UpdateUser(ctx, id, username, email)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
