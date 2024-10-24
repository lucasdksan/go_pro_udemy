package repositories

import (
	"context"
	"errors"
	"go_pro/internal/apperrors"
	"go_pro/internal/database/querys"
	"go_pro/internal/models"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrDuplicateEmail = apperrors.NewRepositoryError(errors.New("duplicate email"))
var ErrInvalidTokenOrUserAlreadyConfirmed = apperrors.NewRepositoryError(errors.New("invalid token or user already confirmed"))
var ErrEmailNotFound = apperrors.NewRepositoryError(errors.New("email not found"))
var fail = func(err error) error {
	slog.Error(err.Error())
	return apperrors.NewRepositoryError(err)
}

type UserRepository interface {
	Create(ctx context.Context, email, password, hashKey string) (*models.User, string, error)
	CreateResetPasswordToken(ctx context.Context, email, hashToken string) (string, error)
	ConfirmUserByToken(ctx context.Context, token string) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserConfirmationByToken(ctx context.Context, token string) (*models.UserConfirmmationToken, error)
	UpdatePasswordByToken(ctx context.Context, pass, token string) (string, error)
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

	tx, err := ur.db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return &user, "", fail(err)
	}

	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, querys.CreateUserQuery, user.Email, user.Password)

	if err := row.Scan(&user.Id, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return &user, "", fail(ErrDuplicateEmail)
		}
		return &user, "", fail(err)
	}

	userToken, err := ur.createConfirmationToken(ctx, tx, &user, hashKey)

	if err != nil {
		return &user, "", fail(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return &user, "", fail(err)
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

	tx, err := ur.db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return fail(err)
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, querys.UpdateUserConfirmedQuery, userId)

	if err != nil {
		return fail(err)
	}

	_, err = tx.Exec(ctx, querys.UpdateTokenConfirmedQuery, tokenId)

	if err != nil {
		return fail(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fail(err)
	}

	return nil
}

func (ur *userRepository) createConfirmationToken(ctx context.Context, tx pgx.Tx, user *models.User, hashKey string) (*models.UserConfirmmationToken, error) {
	var userToken models.UserConfirmmationToken

	userToken.Token = pgtype.Text{String: hashKey, Valid: true}
	userToken.UserId = user.Id

	row := tx.QueryRow(ctx, querys.CreateTokenQuery, userToken.UserId, userToken.Token)

	if err := row.Scan(&userToken.Id, &userToken.CreatedAt); err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return &userToken, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	row := ur.db.QueryRow(ctx, querys.FindByEmailQuery, email)

	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Active); err != nil {
		return nil, apperrors.NewRepositoryError(err)
	}

	return &user, nil
}

func (ur *userRepository) CreateResetPasswordToken(ctx context.Context, email, hashToken string) (string, error) {
	user, err := ur.FindByEmail(ctx, email)

	if err != nil || !user.Active.Bool {
		return "", ErrEmailNotFound
	}

	tx, err := ur.db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return "", fail(err)
	}

	defer tx.Rollback(ctx)

	ctoken, err := ur.createConfirmationToken(ctx, tx, user, hashToken)

	if err != nil {
		return "", ErrEmailNotFound
	}

	if err = tx.Commit(ctx); err != nil {
		return "", fail(err)
	}

	return ctoken.Token.String, nil
}

func (ur *userRepository) GetUserConfirmationByToken(ctx context.Context, token string) (*models.UserConfirmmationToken, error) {
	var userToken models.UserConfirmmationToken

	row := ur.db.QueryRow(ctx, querys.GetUserConfirmationByTokenQuery, token)

	if err := row.Scan(&userToken.Id, &userToken.UserId, &userToken.Token, &userToken.Confirmed, &userToken.CreatedAt, &userToken.UpdatedAt); err != nil {
		return nil, apperrors.NewRepositoryError(err)
	}

	return &userToken, nil
}

func (ur *userRepository) UpdatePasswordByToken(ctx context.Context, pass, token string) (string, error) {
	var userId, tokenId pgtype.Numeric
	var email pgtype.Text
	row := ur.db.QueryRow(ctx, querys.SelectPasswordByTokenQuery, token)

	if err := row.Scan(&userId, &email, &tokenId); err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrInvalidTokenOrUserAlreadyConfirmed
		}

		return "", apperrors.NewRepositoryError(err)
	}

	tx, err := ur.db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return "", fail(err)
	}

	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, querys.UpdateTokenConfirmedQuery, tokenId); err != nil {
		return "", fail(err)
	}

	if _, err := tx.Exec(ctx, querys.UpdatePasswordQuery, pass, userId); err != nil {
		return "", fail(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return "", fail(err)
	}

	return email.String, nil
}
