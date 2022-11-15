package service

import (
	"context"
	"task/helper"
	"task/model/domain"
	"task/model/web"
	"task/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CheckUsername(ctx context.Context, username string) web.UserReponse
	Register(ctx context.Context, request web.UserCreateRequest) web.UserReponse
	Login(ctx context.Context, request web.LoginRequest) (string, error)
	CheckID(id int) bool
	Delete(id int) error
}

type AuthRepositoryImpl struct {
	AuthRepository repository.AuthRepo
	Validate       *validator.Validate
}

func NewUserService(auth repository.AuthRepo, validate *validator.Validate) AuthService {
	return &AuthRepositoryImpl{
		AuthRepository: auth,
		Validate:       validate,
	}
}

func (s *AuthRepositoryImpl) CheckUsername(ctx context.Context, username string) web.UserReponse {
	user := s.AuthRepository.CheckUsername(username)

	return web.UserReponse{
		Username: user.Username,
		Id:       int(user.Id),
		Name:     user.Name,
	}
}

func (s *AuthRepositoryImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserReponse {
	err := s.Validate.Struct(request)
	helper.PanicHandling(err)

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicHandling(err)

	userDomain := domain.User{
		Username: request.Username,
		Name:     request.Name,
		Password: string(password),
	}

	so := s.AuthRepository.Register(&userDomain)
	helper.PanicHandling(so)
	return web.UserReponse{
		Username: userDomain.Username,
		Name:     userDomain.Name,
		Id:       int(userDomain.Id),
	}
}

func (s *AuthRepositoryImpl) Login(ctx context.Context, request web.LoginRequest) (string, error) {

	password, err := s.AuthRepository.Login(request.Username)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (s *AuthRepositoryImpl) CheckID(id int) bool {
	return s.AuthRepository.CheckId(id)
}

func (s *AuthRepositoryImpl) Delete(id int) error {
	if err := s.AuthRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
