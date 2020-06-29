package usecase

import (
	"github.com/personal-work/space_exploration/internal/models"
)

func createAppError(msg string, code int) *models.AppError {
	return &models.AppError{
		ErrorMessage:    msg,
		HTTPResposeCode: code,
	}
}
