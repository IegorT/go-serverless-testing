package createtask

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/iegort/go-serverless-testing/internal/common/entities"
)

type storage interface {
	SaveTask(context.Context, entities.Task) error
}

type CreateTask struct {
	store storage
	log   slog.Logger
}

func NewCreateTask(store storage, log slog.Logger) *CreateTask {
	return &CreateTask{store: store, log: log}
}

func (uc *CreateTask) Exec(ctx context.Context, dto CreateTaskInput) (CreateTaskOutput, error) {
	var output CreateTaskOutput

	if err := dto.Validate(); err != nil {
		return output, err
	}

	task := uc.createTask(dto)
	if err := uc.store.SaveTask(ctx, task); err != nil {
		return output, err
	}

	return output.fromTask(task), nil
}

func (uc *CreateTask) createTask(dto CreateTaskInput) entities.Task {
	return entities.Task{
		ID:          uuid.NewString(),
		Title:       dto.Title,
		Description: dto.Description,
		Deadline:    dto.Deadline,
	}

}
