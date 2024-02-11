package activation

import model "github.com/ArthurMaverick/zap-ai/internal/models"

type Service interface {
	ActivationService(input *InputActivation) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceActivation(respository Repository) *service {
	return &service{
		repository: respository,
	}
}

func (s *service) ActivationService(input *InputActivation) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		Email:  input.Email,
		Active: input.Active,
	}

	activationResult, activationError := s.repository.ActivationRepository(&users)
	return activationResult, activationError
}
