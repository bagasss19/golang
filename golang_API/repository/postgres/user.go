package postgresrepository

import (
	"context"
	"fmt"
	"golang_api/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, UserID uint32) (data *model.User, err error)
	GetUserByUsername(ctx context.Context, UserName string) (data *model.User, err error)
	GetAllUser(ctx context.Context) (record []*model.User, err error)
	UpdateUser(ctx context.Context, updatedData *model.User, UserID uint32) (result *model.User, err error)
	CreateUser(ctx context.Context, User *model.User) (result *model.User, err error)
	DeleteUser(ctx context.Context, userID uint32) (bool, error)
}

func NewUserRepository(sqlConnection *gorm.DB) UserRepository {
	return &UserDB{
		repoPrefix: "userRepo",
		db:         sqlConnection,
	}
}

func (u *UserDB) GetUserByID(ctx context.Context, UserID uint32) (data *model.User, err error) {
	prefix := u.repoPrefix + ".GetUserByID"
	query := u.db.WithContext(ctx)

	if err := query.Where("id = ?", UserID).First(&data).Error; err != nil {
		return nil, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return data, nil
}

func (u *UserDB) GetUserByUsername(ctx context.Context, UserName string) (data *model.User, err error) {
	prefix := u.repoPrefix + ".GetUserByID"
	query := u.db.WithContext(ctx)

	if err := query.Where("username = ?", UserName).First(&data).Error; err != nil {
		return nil, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return data, nil
}

func (u *UserDB) GetAllUser(ctx context.Context) (record []*model.User, err error) {
	prefix := u.repoPrefix + ".GetAllUser"
	query := u.db.WithContext(ctx)

	if err := query.Find(&record).Error; err != nil {
		return nil, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return record, nil
}

func (u *UserDB) UpdateUser(ctx context.Context, updatedData *model.User, UserID uint32) (result *model.User, err error) {
	prefix := u.repoPrefix + ".UpdateUser"

	query := u.db.WithContext(ctx).
		Where("id = ?", UserID).
		Select(
			"username",
			"password",
		)

	if err := query.Updates(&updatedData).Error; err != nil {
		return nil, errors.Wrapf(err, "[%s] error while update", prefix)
	}

	return result, err
}

func (u *UserDB) CreateUser(ctx context.Context, User *model.User) (result *model.User, err error) {
	prefix := u.repoPrefix + ".CreateUser"
	query := u.db.WithContext(ctx)

	db := query.Save(User)
	if err = db.Error; err != nil {
		return nil, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return result, nil
}

func (u *UserDB) DeleteUser(ctx context.Context, userID uint32) (bool, error) {
	prefix := u.repoPrefix + ".DeleteUser"
	query := u.db.WithContext(ctx)

	if err := query.Delete(&model.User{}, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return true, nil
}
