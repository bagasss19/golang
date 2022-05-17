package usecase

import (
	"context"
	"fmt"
	"golang_api/helper"
	"golang_api/middleware"
	"golang_api/model"
	repository "golang_api/repository/postgres"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, UserID uint32) (data *model.User, err error)
	GetAllUser(ctx context.Context) (data []*model.User, err error)
	CreateUser(ctx context.Context, data *model.User) error
	DeleteUser(ctx context.Context, UserID uint32) error
	UpdateUser(ctx context.Context, data *model.User, UserID uint32) error
	Login(ctx context.Context, data *model.User) (token string, err error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}
func (u *userUsecase) GetUserByID(ctx context.Context, UserID uint32) (data *model.User, err error) {
	prefix := ".surveyUsecase.GetUserByID"

	user, err := u.userRepository.GetUserByID(ctx, UserID)
	if err != nil {
		if err == helper.ErrorDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("[%s] error while get one user: %+v", prefix, err)
	}

	return user, nil
}

func (u *userUsecase) GetAllUser(ctx context.Context) (data []*model.User, err error) {
	prefix := ".surveyUsecase.GetAllUser"

	users, err := u.userRepository.GetAllUser(ctx)
	if err != nil {
		if err == helper.ErrorDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("[%s] error while get all users: %+v", prefix, err)
	}

	return users, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, data *model.User) error {
	prefix := ".surveyUsecase.CreateUser"

	found, _ := u.userRepository.GetUserByUsername(ctx, data.Username)
	if found != nil {
		return fmt.Errorf("[%s] user already exist", prefix)
	}

	if data.Username == "" {
		return fmt.Errorf("[%s] username cannot be empty", prefix)
	}

	if data.Password == "" {
		return fmt.Errorf("[%s] password cannot be empty", prefix)
	}
	_, err := u.userRepository.CreateUser(ctx, data)
	if err != nil {
		if err == helper.ErrorDataNotFound {
			return nil
		}
		return fmt.Errorf("[%s] error while create users: %+v", prefix, err)
	}

	return nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, UserID uint32) error {
	prefix := ".surveyUsecase.CreateUser"

	_, err := u.userRepository.DeleteUser(ctx, UserID)
	if err != nil {
		if err == helper.ErrorDataNotFound {
			return nil
		}
		return fmt.Errorf("[%s] error while create users: %+v", prefix, err)
	}

	return nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, data *model.User, UserID uint32) error {
	prefix := ".surveyUsecase.UpdateUser"
	_, err := u.userRepository.UpdateUser(ctx, data, UserID)
	if err != nil {
		if err == helper.ErrorDataNotFound {
			return nil
		}
		return fmt.Errorf("[%s] error while update users: %+v", prefix, err)
	}

	return nil
}

func (u *userUsecase) Login(ctx context.Context, data *model.User) (token string, err error) {
	prefix := ".surveyUsecase.Login"

	found, _ := u.userRepository.GetUserByUsername(ctx, data.Username)
	if found == nil {
		return "", fmt.Errorf("[%s] wrong username or password", prefix)
	}

	if data.Username == "" {
		return "", fmt.Errorf("[%s] username cannot be empty", prefix)
	}

	if data.Password == "" {
		return "", fmt.Errorf("[%s] password cannot be empty", prefix)
	}

	checkPass, _ := helper.CompareHashAndPassword(data.Password, found.Password)

	if !checkPass {
		return "", fmt.Errorf("[%s] wrong username or password: %+v", prefix, err)
	}

	token, err = middleware.CreateToken(uint32(found.ID), found.Username)

	if err != nil {
		return "", fmt.Errorf("[%s] error while getting token: %+v", prefix, err)
	}

	return token, nil
}
