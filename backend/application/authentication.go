package application

import (
	"context"
	"react-go/backend/domain/model"
	"react-go/backend/domain/repository"
	"react-go/backend/domain/service"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, param *model.User) (*model.User, error)
}

type AuthenticationServiceDIInput struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewAuthenticationServiceDIInput(uRepo repository.UserRepository, uService service.UserService) *AuthenticationServiceDIInput {
	return &AuthenticationServiceDIInput{
		userRepository: uRepo,
		userService:    uService,
	}
}

type authenticationService struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewAuthenticationService(diInput *AuthenticationServiceDIInput) AuthenticationService {
	return &authenticationService{
		userRepository: diInput.userRepository,
		userService:    diInput.userService,
	}
}

func (s *authenticationService) SignUp(ctx context.Context, params *model.User) (user *model.User, err error) {
	user, err = s.userService.NewUser(params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	user, err = s.createUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authenticationService) createUser(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := s.userRepository.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
