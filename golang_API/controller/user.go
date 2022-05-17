package controller

import (
	"golang_api/model"
	"golang_api/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserByID(c *gin.Context)
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{
		userUsecase: userUsecase,
	}
}

func (u *userController) GetUserByID(c *gin.Context) {
	var (
		response model.Response
	)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.BindQuery(&id); err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(err.Error()))
		return
	}

	if id == 0 {
		c.AbortWithStatusJSON(response.ErrorBadRequest("user_id cannot 0"))
		return
	}

	found, err := u.userUsecase.GetUserByID(c, uint32(id))
	if err != nil {
		c.AbortWithStatusJSON(response.ErrorInternalServer(err))
		return
	}
	if found == nil {
		c.AbortWithStatusJSON(response.ErrorDataNotFound())
		return
	}

	c.JSON(response.SuccessData(found))
}

func (u *userController) GetAllUser(c *gin.Context) {
	var (
		response model.Response
	)

	found, err := u.userUsecase.GetAllUser(c)
	if err != nil {
		c.AbortWithStatusJSON(response.ErrorInternalServer(err))
		return
	}
	if found == nil {
		c.AbortWithStatusJSON(response.ErrorDataNotFound())
		return
	}

	c.JSON(response.SuccessData(found))
}

func (u *userController) CreateUser(c *gin.Context) {
	var (
		request struct {
			Username *string `json:"username"`
			Password *string `json:"password"`
		}
		response model.Response
	)

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(err.Error()))
		return
	}

	newUser := new(model.User)
	newUser.Username = *request.Username
	newUser.Password = *request.Password

	err := u.userUsecase.CreateUser(c, newUser)
	if err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(string(err.Error())))
		return
	}

	c.JSON(response.SuccessCreated(err))
}

func (u *userController) DeleteUser(c *gin.Context) {
	var (
		response model.Response
	)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.BindQuery(&id); err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(err.Error()))
		return
	}

	if id == 0 {
		c.AbortWithStatusJSON(response.ErrorBadRequest("user_id cannot 0"))
		return
	}

	err := u.userUsecase.DeleteUser(c, uint32(id))
	if err != nil {
		c.AbortWithStatusJSON(response.ErrorInternalServer(err))
		return
	}

	c.JSON(response.SuccessData(nil))
}

func (u *userController) UpdateUser(c *gin.Context) {
	var (
		request struct {
			Username *string `json:"username"`
			Password *string `json:"password"`
		}
		response model.Response
	)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(err.Error()))
		return
	}

	updatedUser := new(model.User)
	updatedUser.Username = *request.Username
	updatedUser.Password = *request.Password

	err := u.userUsecase.UpdateUser(c, updatedUser, uint32(id))
	if err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(string(err.Error())))
		return
	}

	c.JSON(response.SuccessUpdated(err))
}

func (u *userController) Login(c *gin.Context) {
	var (
		request struct {
			Username *string `json:"username"`
			Password *string `json:"password"`
		}
		response model.Response
	)

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(err.Error()))
		return
	}

	newUser := new(model.User)
	newUser.Username = *request.Username
	newUser.Password = *request.Password

	token, err := u.userUsecase.Login(c, newUser)
	if err != nil {
		c.AbortWithStatusJSON(response.ErrorBadRequest(string(err.Error())))
		return
	}

	c.JSON(response.SuccessData(token))
}
