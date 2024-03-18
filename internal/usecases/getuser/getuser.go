package getuser

import (
	"context"
	"log/slog"

	"github.com/iegort/go-serverless-testing/internal/common/entities"
)

type storage interface {
	getUser(context.Context, GetUserInput) (entities.User, error)
}

type GetUser struct {
	store storage
	log   slog.Logger
}

func NewGetUser(store storage, log slog.Logger) *GetUser {
	return &GetUser{store: store, log: log}
}

func (uc *GetUser) Exec(ctx context.Context, input GetUserInput) (entities.User, error) {
	return uc.store.getUser(ctx, input)
}
