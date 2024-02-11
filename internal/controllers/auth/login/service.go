package login

import (
	model "github.com/ArthurMaverick/zap-ai/internal/models"
)

type Service interface {
	LoginService(input *InputLogin)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&users)
	return resultLogin, errLogin
}
