package repositories

import (
	"context"
	"go_pro/internal/database"
	"go_pro/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
}

type noteRepository struct {
	db *pgxpool.Pool
}

func (nr *noteRepository) List() ([]models.Note, error) {
	var list []models.Note

	rows, err := nr.db.Query(context.Background(), database.ListNoteQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Note

		if err = rows.Scan(
			&row.Id, &row.Title,
			&row.Content, &row.Color,
			&row.CreatedAt); err != nil {
			return nil, err
		}

		list = append(list, row)
	}

	return list, nil
}

func NewNoteRepository(db *pgxpool.Pool) NoteRepository {
	return &noteRepository{
		db: db,
	}
}
