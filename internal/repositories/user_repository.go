package repositories

import (
	"context"
	"errors"
	"go_pro/internal/apperrors"
	"go_pro/internal/database/querys"
	"go_pro/internal/models"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrDuplicateEmail = apperrors.NewRepositoryError(errors.New("duplicate email"))
var ErrInvalidTokenOrUserAlreadyConfirmed = apperrors.NewRepositoryError(errors.New("invalid token or user already confirmed"))

type UserRepository interface {
	Create(ctx context.Context, email, password, hashKey string) (*models.User, string, error)
	ConfirmUserByToken(ctx context.Context, token string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, email, password, hashKey string) (*models.User, string, error) {
	var user models.User

	user.Email = pgtype.Text{String: email, Valid: true}
	user.Password = pgtype.Text{String: password, Valid: true}

	row := ur.db.QueryRow(ctx, querys.CreateUserQuery, user.Email, user.Password)

	if err := row.Scan(&user.Id, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return &user, "", ErrDuplicateEmail
		}
		return &user, "", apperrors.NewRepositoryError(err)
	}

	userToken, err := ur.createConfirmationToken(ctx, &user, hashKey)

	if err != nil {
		return nil, "", err
	}

	return &user, userToken.Token.String, nil
}

func (ur *userRepository) ConfirmUserByToken(ctx context.Context, token string) error {
	var userId, tokenId pgtype.Numeric
	row := ur.db.QueryRow(ctx, querys.GetUserJoinUserTokenQuery, token)

	if err := row.Scan(&userId, &tokenId); err != nil {
		if err == pgx.ErrNoRows {
			return ErrInvalidTokenOrUserAlreadyConfirmed
		}
		return apperrors.NewRepositoryError(err)
	}

	_, err := ur.db.Exec(ctx, querys.UpdateUserConfirmedQuery, userId)

	if err != nil {
		return apperrors.NewRepositoryError(err)
	}

	_, err = ur.db.Exec(ctx, querys.UpdateTokenConfirmedQuery, tokenId)

	if err != nil {
		return apperrors.NewRepositoryError(err)
	}

	return nil
}

func (ur *userRepository) createConfirmationToken(ctx context.Context, user *models.User, hashKey string) (*models.UserConfirmmationToken, error) {
	var userToken models.UserConfirmmationToken

	userToken.Token = pgtype.Text{String: hashKey, Valid: true}
	userToken.UserId = user.Id

	row := ur.db.QueryRow(ctx, querys.CreateTokenQuery, userToken.UserId, userToken.Token)

	if err := row.Scan(&userToken.Id, &userToken.CreatedAt); err != nil {
		return nil, err
	}

	return &userToken, nil
}
