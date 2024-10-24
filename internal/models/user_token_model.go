package models

import "github.com/jackc/pgx/v5/pgtype"

type UserConfirmmationToken struct {
	Id        pgtype.Numeric
	UserId    pgtype.Numeric
	Token     pgtype.Text
	Confirmed pgtype.Bool
	CreatedAt pgtype.Date
	UpdatedAt pgtype.Date
}
