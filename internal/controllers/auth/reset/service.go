package reset

import model "github.com/ArthurMaverick/zap-ai/internal/models"

type Service interface {
	ResetService(input *InputReset) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) ResetService(input *InputReset) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
		Active:   input.Active,
	}
	resetResult, errResult := s.repository.ResetRepository(&users)
	return resetResult, errResult
}
