package repositories

import (
	"context"
	"errors"
	"go_pro/internal/apperrors"
	"go_pro/internal/database/querys"
	"go_pro/internal/models"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrDuplicateEmail = apperrors.NewRepositoryError(errors.New("duplicate email"))

type UserRepository interface {
	Create(ctx context.Context, email, password string) (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, email, password string) (*models.User, error) {
	var user models.User

	user.Email = pgtype.Text{String: email, Valid: true}
	user.Password = pgtype.Text{String: password, Valid: true}

	row := ur.db.QueryRow(ctx, querys.CreateUserQuery, user.Email, user.Password)

	if err := row.Scan(&user.Id, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return &user, ErrDuplicateEmail
		}
		return &user, apperrors.NewRepositoryError(err)
	}

	return &user, nil
}
