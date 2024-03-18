package createtask

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iegort/go-serverless-testing/internal/common/entities"
)

type CreateTaskInput struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Deadline    time.Time `json:"deadline" validate:"required"`
}

func (dto *CreateTaskInput) Validate() error {
	return validator.New().Struct(dto)
}

type CreateTaskOutput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

func (output *CreateTaskOutput) fromTask(entity entities.Task) CreateTaskOutput {
	output.ID = entity.ID
	output.Title = entity.Title
	output.Description = entity.Description
	output.Deadline = entity.Deadline

	return *output
}
