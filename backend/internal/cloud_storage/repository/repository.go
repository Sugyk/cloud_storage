package repository

import (
	"context"
	"fmt"

	"github.com/fir1/rest-api/internal/cloud_storage/model"
	"github.com/fir1/rest-api/pkg/db"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

type FileHead struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
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
	query := "SELECT id FROM Files WHERE id = $1;"
	err := r.Db.GetContext(ctx, &entity, query, id)
	fmt.Println(err, "asdasdasd", entity)
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

func (r Repository) ListFiles(ctx context.Context, user_id int, page int, offset int) ([]FileHead, error) {
	var entities []FileHead
	query := "SELECT id, filename FROM Files WHERE user_id = $1 LIMIT $2 OFFSET $3;"
	err := r.Db.SelectContext(ctx, &entities, query, user_id, offset, offset*page)

	return entities, db.HandleError(err)
}
