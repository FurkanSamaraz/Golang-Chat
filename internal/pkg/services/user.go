package services

import (
	"fmt"

	api_structure "github.com/FurkanSamaraz/Golang-Chat/internal/pkg/structures"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct{ DB *gorm.DB }

type IUserInstance interface {
	Login(filter api_structure.User) (string, error)
	Register(data api_structure.User) (api_structure.StatusMessage, error)
	VerifyContact(username string) *api_structure.Response
	ChatHistory(username1, username2, fromTS, toTS string) *api_structure.Response
	ContactList(username string) *api_structure.Response
}

func (r *UserService) Login(filter api_structure.User) (string, error) {
	result := api_structure.User{}

	var err error

	if err = r.DB.Table(result.TableName()).Preload(clause.Associations).Model(&api_structure.User{}).Where(filter).Find(&result).Error; err != nil {
		fmt.Printf("not user error")
		return err.Error(), err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   result.Id,
		"name": result.Name,
	})
	tokenString, err := token.SignedString([]byte("gizli-anahtar"))
	if err != nil {
		fmt.Println("hatalı token")
	}
	return tokenString, err
}

func (r *UserService) Register(data api_structure.User) (api_structure.StatusMessage, error) {
	result := api_structure.StatusMessage{}
	var err error
	if err = r.DB.Table(result.TableName()).Create(&data).Error; err != nil {
		result.Message = "Error Register"
		return result, err
	}
	result.Message = "Successfully Register"

	return result, err
}
