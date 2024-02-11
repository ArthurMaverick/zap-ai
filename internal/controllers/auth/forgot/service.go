package forgot

import model "github.com/ArthurMaverick/zap-ai/internal/models"

type Service interface {
	ForgotService(input *InputForgot) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceForgot(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ForgotService(input *InputForgot) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		Email: input.Email,
	}
	resultRegister, errRegister := s.repository.ForgotRepository(&users)
	return resultRegister, errRegister
}
