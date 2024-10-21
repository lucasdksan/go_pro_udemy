package repositories

import (
	"context"
	"go_pro/internal/apperrors"
	"go_pro/internal/database/querys"
	"go_pro/internal/models"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List(ctx context.Context, userId int) ([]models.Note, error)
	GetById(ctx context.Context, userId int, id int) (*models.Note, error)
	Create(ctx context.Context, userId int, title, content, color string) (*models.Note, error)
	Update(ctx context.Context, userId int, id int, title, content, color string) (*models.Note, error)
	Delete(ctx context.Context, userId int, id int) error
}

type noteRepository struct {
	db *pgxpool.Pool
}

func (nr *noteRepository) List(ctx context.Context, userId int) ([]models.Note, error) {
	var list []models.Note

	rows, err := nr.db.Query(ctx, querys.ListNoteQuery, userId)

	if err != nil {
		return nil, apperrors.NewRepositoryError(err)
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Note

		if err = rows.Scan(
			&row.Id, &row.Title,
			&row.Content, &row.Color,
			&row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, apperrors.NewRepositoryError(err)
		}

		list = append(list, row)
	}

	return list, nil
}

func (nr *noteRepository) GetById(ctx context.Context, userId int, id int) (*models.Note, error) {
	var note models.Note

	row := nr.db.QueryRow(ctx, querys.GetByIdNoteQuery, id, userId)

	if err := row.Scan(&note.Id, &note.Title,
		&note.Content, &note.Color,
		&note.CreatedAt, &note.UpdatedAt); err != nil {
		return &note, apperrors.NewRepositoryError(err)
	}

	return &note, nil
}

func (nr *noteRepository) Create(ctx context.Context, userId int, title, content, color string) (*models.Note, error) {
	var note models.Note

	note.Title = pgtype.Text{String: title, Valid: true}
	note.Content = pgtype.Text{String: content, Valid: true}
	note.Color = pgtype.Text{String: color, Valid: true}

	row := nr.db.QueryRow(ctx, querys.CreateNoteQuery, note.Title, note.Content, note.Color, userId)

	if err := row.Scan(&note.Id, &note.CreatedAt); err != nil {
		return &models.Note{}, apperrors.NewRepositoryError(err)
	}

	return &note, nil
}

func (nr *noteRepository) Update(ctx context.Context, userId int, id int, title, content, color string) (*models.Note, error) {
	var note models.Note

	var titleValue, contentValue, colorValue, updatedAtValue interface{}

	if len(title) > 0 {
		titleValue = title
	} else {
		titleValue = nil
	}

	if len(content) > 0 {
		contentValue = content
	} else {
		contentValue = nil
	}

	if len(color) > 0 {
		colorValue = color
	} else {
		colorValue = nil
	}

	updatedAtValue = time.Now()

	_, err := nr.db.Exec(ctx, querys.UpdateNoteQuery, titleValue, contentValue, colorValue, updatedAtValue, id, userId)

	if err != nil {
		return &models.Note{}, apperrors.NewRepositoryError(err)
	}

	note.Title = pgtype.Text{String: title, Valid: true}
	note.Content = pgtype.Text{String: content, Valid: true}
	note.Color = pgtype.Text{String: color, Valid: true}
	note.UpdatedAt = pgtype.Date{Time: time.Now(), Valid: true}
	note.Id = pgtype.Numeric{Int: big.NewInt(int64(id))}

	return &note, nil
}

func (nr *noteRepository) Delete(ctx context.Context, userId int, id int) error {
	_, err := nr.db.Exec(ctx, querys.DeleteNoteQuery, id, userId)

	if err != nil {
		return apperrors.NewRepositoryError(err)
	}

	return nil
}

func NewNoteRepository(db *pgxpool.Pool) NoteRepository {
	return &noteRepository{
		db: db,
	}
}
