package register

import model "github.com/ArthurMaverick/zap-ai/internal/models"

type Service interface {
	RegisterService(input *InputRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterService(input *InputRegister) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}
	resultRegister, errRegister := s.repository.RegisterRepository(&users)
	return resultRegister, errRegister
}
