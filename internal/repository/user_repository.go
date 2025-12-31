package repository

import (
	"context"
	"database/sql"
	"tasked/internal/database"
	"tasked/internal/domain"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, username string, email string, password string) (*domain.User, error)
	UpdateUser(ctx context.Context, id int64, username string, email string) (*domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	queries *database.Queries
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		queries: database.New(db),
	}
}

func (r *userRepository) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	dbUser, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}, nil
}

func (r *userRepository) CreateUser(ctx context.Context, username string, email string, password string) (*domain.User, error) {
	dbUser, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int64, username string, email string) (*domain.User, error) {
	dbUser, err := r.queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:       id,
		Username: username,
		Email:    email,
	})
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	return r.queries.DeleteUser(ctx, id)
}
