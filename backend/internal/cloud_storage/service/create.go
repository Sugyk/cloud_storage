package service

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/fir1/rest-api/internal/cloud_storage/model"
	"github.com/fir1/rest-api/pkg/erru"
)

type CreateParams struct {
	Telegram_id int    `valid:"required"`
	Username    string `valid:"required"`
}

func (s Service) CreateUser(ctx context.Context, params CreateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return erru.ErrArgument{Wrapped: err}
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Users{
		Telegram_id: params.Telegram_id,
		Username:    params.Username,
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
