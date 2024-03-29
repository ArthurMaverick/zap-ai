package resend

import model "github.com/ArthurMaverick/zap-ai/internal/models"

type Service interface {
	ResendService(input *InputResend) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceResend(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) ResendService(input *InputResend) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		Email: input.Email,
	}
	resultRegister, errRegister := s.repository.ResendRepository(&users)
	return resultRegister, errRegister
}
