package service

import (
	"crypto/sha512"
	"encoding/base64"
	"intikom_test/entity"
	"intikom_test/model"
	"intikom_test/repository"
	"log"
	"time"
)

type userService struct {
	repoUser repository.UserRepository
}

func NewUserService(repoUser repository.UserRepository) UserService {
	return &userService{repoUser}
}

func (svc *userService) CreateUser(req *model.CreateUserRequest) error {
	dataExist, err := svc.repoUser.GetByEmail(req.Email)
	if err != nil {
		log.Println("Error while check email, cause: ", err)
		return model.NewError("500", "Internal server error.")
	} else if dataExist != nil {
		return model.NewError("400", "Email already exist.")
	}

	var password = base64.StdEncoding.EncodeToString([]byte(req.Password))
	var hash = sha512.New()
	hash.Write([]byte(password))
	var hashPassword = hash.Sum(nil)

	timeNow := time.Now()

	var newData = &entity.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashPassword),
		CreatedAt: &timeNow,
	}

	_, err = svc.repoUser.Create(*newData)
	if err != nil {
		log.Println("Error while create new data, cause: ", err)
		return model.NewError("500", "Internal server error.")
	}

	return nil
}

func (svc *userService) DeleteUser(id string) error {
	dataUser, err := svc.repoUser.Get(id)
	if err != nil {
		log.Println("Error while get data, cause: ", err)
		return model.NewError("500", "Internal server error.")
	} else if dataUser == nil {
		return model.NewError("404", "Data not found.")
	}

	err = svc.repoUser.Delete(dataUser)
	if err != nil {
		log.Println("Error while delete data user, cause: ", err)
		return model.NewError("500", "Internal server error.")
	}

	return nil
}

func (svc *userService) GetUser(id string) (*model.DataUserResponse, error) {
	dataUser, err := svc.repoUser.Get(id)
	if err != nil {
		log.Println("Error while get data, cause: ", err)
		return nil, model.NewError("500", "Internal server error.")
	} else if dataUser == nil {
		return nil, model.NewError("404", "Data not found.")
	}

	var updatedAt string
	if dataUser.UpdatedAt != nil {
		updatedAt = dataUser.UpdatedAt.Format("02-Jan-2006 15:04")
	}

	return &model.DataUserResponse{
		Name:      dataUser.Name,
		Email:     dataUser.Email,
		CreatedAt: dataUser.CreatedAt.Format("02-Jan-2006 15:04"),
		UpdatedAt: updatedAt,
	}, nil
}

func (svc *userService) ListUser() ([]model.DataUserResponse, int64, error) {
	dataUsers, total, err := svc.repoUser.List(0, 0)
	if err != nil {
		log.Println("Error while get data, cause: ", err)
		return make([]model.DataUserResponse, 0), 0, model.NewError("500", "Internal server error.")
	} else if len(dataUsers) == 0 {
		return make([]model.DataUserResponse, 0), 0, nil
	}

	var respData []model.DataUserResponse
	for _, data := range dataUsers {
		var updatedAt string
		if data.UpdatedAt != nil {
			updatedAt = data.UpdatedAt.Format("02-Jan-2006 15:04")
		}
		respData = append(respData, model.DataUserResponse{
			Name:      data.Name,
			Email:     data.Email,
			CreatedAt: data.CreatedAt.Format("02-Jan-2006 15:04"),
			UpdatedAt: updatedAt,
		})
	}

	return respData, total, nil
}

func (svc *userService) UpdateUser(req *model.UpdateUserRequest) error {
	dataExist, err := svc.repoUser.Get(req.ID)
	if err != nil {
		log.Println("Error while get data, cause: ", err)
		return model.NewError("500", "Internal server error.")
	} else if dataExist == nil {
		return model.NewError("400", "Data not found.")
	}

	emailExist, err := svc.repoUser.GetByEmail(req.Email)
	if err != nil {
		log.Println("Error while check email, cause: ", err)
		return model.NewError("500", "Internal server error.")
	} else if emailExist != nil {
		return model.NewError("400", "Email already exist.")
	}

	timeNow := time.Now()

	var password = base64.StdEncoding.EncodeToString([]byte(req.Password))
	var hash = sha512.New()
	hash.Write([]byte(password))
	var hashPassword = hash.Sum(nil)

	var newData = &entity.User{
		ID:        dataExist.ID,
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashPassword),
		CreatedAt: dataExist.CreatedAt,
		UpdatedAt: &timeNow,
	}

	err = svc.repoUser.Update(newData)
	if err != nil {
		log.Println("Error while update data, cause: ", err)
		return model.NewError("500", "Internal server error.")
	}

	return nil
}
