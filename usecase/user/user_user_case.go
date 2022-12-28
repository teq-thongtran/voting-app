package user

import (
	"context"
	"errors"
	"fmt"
	"myapp/auth"
	"myapp/customError"
	"myapp/model"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/user"
	"strings"

	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUserCase interface {
	Create(ctx context.Context, req *payload.CreateUserRequest) (*presenter.SignUpResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateUserRequest) (*presenter.UserResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.UserResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest) (*presenter.ListUserResponseWrapper, error)
	SignIn(ctx context.Context, req *payload.SignInRequest) (*presenter.SignUpResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	UserRepo user.Repository
}

func New(repo *repository.Repository) UserUserCase {
	return &UseCase{
		UserRepo: repo.User,
	}
}

func (u *UseCase) validateCreate(req *payload.CreateUserRequest) error {
	if req.Name == "" {
		return customError.ErrRequestInvalidParam("name")
	}

	if req.Username == "" {
		return customError.ErrRequestInvalidParam("UserName")
	}

	if req.Password == "" {
		return customError.ErrRequestInvalidParam("Password")
	}

	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) == 0 {
		req.Name = ""
		return customError.ErrRequestInvalidParam("name")
	}

	if len(req.Username) == 0 {
		return customError.ErrRequestInvalidParam("username")
	}

	if len(req.Password) == 0 {
		return customError.ErrRequestInvalidParam("password")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserRequest,
) (*presenter.SignUpResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	Password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	fmt.Println("PASSWORD: ", Password, string(Password))
	myUser := &model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: string(Password),
	}

	err := u.UserRepo.Create(ctx, myUser)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	jwtToken, err := auth.Generate_JWT_Token(req.Username)

	if err != nil {
		return nil, customError.ErrRequestInvalidParam("JWT TOKEN")
	}

	return &presenter.SignUpResponseWrapper{User: myUser, Token: jwtToken}, nil
}

func (u *UseCase) SignIn(
	ctx context.Context,
	req *payload.SignInRequest,
) (*presenter.SignUpResponseWrapper, error) {
	myUser, err := u.UserRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}
	res := bcrypt.CompareHashAndPassword([]byte(myUser.Password), []byte(req.Password))

	if res != nil {
		return nil, customError.ErrInvalidParams(res)
	}

	jwtToken, err := auth.Generate_JWT_Token(req.Username)

	if err != nil {
		return nil, customError.ErrRequestInvalidParam("JWT TOKEN")
	}

	return &presenter.SignUpResponseWrapper{User: myUser, Token: jwtToken}, nil

}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateUserRequest) (*model.User, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrModelNotFound()
		}

		return nil, customError.ErrModelGet(err, "User")
	}

	if req.Name != nil {
		*req.Name = strings.TrimSpace(*req.Name)
		if len(*req.Name) == 0 {
			return nil, customError.ErrRequestInvalidParam("name")
		}

		myUser.Name = *req.Name
	}

	if req.Username != nil {
		*req.Username = strings.TrimSpace(*req.Username)
		if len(*req.Username) == 0 {
			return nil, customError.ErrRequestInvalidParam("Username")
		}

		myUser.Username = *req.Username
	}

	return myUser, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.UserRepo.Update(ctx, myUser)
	if err != nil {
		return nil, customError.ErrModelUpdate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrModelNotFound()
		}

		return customError.ErrModelGet(err, "User")
	}

	err = u.UserRepo.Delete(ctx, myUser, false)
	if err != nil {
		return customError.ErrModelDelete(err)
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
) (*presenter.ListUserResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s", req.OrderBy))
	}

	myUsers, total, err := u.UserRepo.GetList(ctx, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListUserResponseWrapper{
		Users: myUsers,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrModelNotFound()
		}

		return nil, customError.ErrModelGet(err, "User")
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}
