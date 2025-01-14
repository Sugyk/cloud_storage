package repository

import (
	"context"

	"github.com/fir1/rest-api/internal/cloud_storage/model"
	"github.com/fir1/rest-api/pkg/db"
	"github.com/jmoiron/sqlx"
)

const OFFSET = 10

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Create(ctx context.Context, entity *model.Users) error {
	query := `INSERT INTO Users (telegram_id, username)
				VALUES (:telegram_id, :username)
				ON CONFLICT (telegram_id)
				DO UPDATE SET username = EXCLUDED.username;`
	_, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}
	return db.HandleError(err)
}

func (r Repository) FindFile(ctx context.Context, id int) (model.Files, error) {
	entity := model.Files{}
	query := "SELECT * FROM Files WHERE id = $1;"
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, db.HandleError(err)
}

func (r Repository) DeleteFile(ctx context.Context, id int) error {
	query := "DELETE FROM Files WHERE id = $1;"
	_, err := r.Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) ListFiles(ctx context.Context, user_id int, page int) ([]model.Files, error) {
	entities := []model.Files{}
	query := "SELECT * FROM Files WHERE user_id = $1 LIMIT 10 OFFSET $2;"
	err := r.Db.GetContext(ctx, &entities, query, user_id, OFFSET*page)
	return entities, db.HandleError(err)
}
