package model

import (
	"log"

	api_structure "github.com/FurkanSamaraz/Golang-Chat/internal/pkg/structures"
)

type UserModel struct {
	Scv RedisService
}

func (redisModel *UserModel) Register(user *api_structure.User) *api_structure.Response {
	// kullanıcı setinde kullanıcı adı olup olmadığını kontrol edin
	// varsa hata döndür
	// yeni kullanıcı oluştur
	// hata için yanıt oluştur
	res := &api_structure.Response{Status: true}

	status, err := redisModel.Scv.IsUserExist(user.Username)
	if err != nil {
		log.Fatal(err)
	}
	if status {
		res.Status = false
		res.Message = "username already taken. try something else."
		return res
	}

	err = redisModel.Scv.RegisterNewUser(user)
	if err != nil {
		res.Status = false
		res.Message = "something went wrong while registering the user. please try again after sometime."
		return res
	}

	return res
}

func (redisModel *UserModel) Login(user *api_structure.User) *api_structure.Response {
	// geçersiz kullanıcı adı ve şifre hatası verirse
	// geçerli kullanıcı ise yeni oturum oluştur
	res := &api_structure.Response{Status: true}

	err := redisModel.Scv.IsUserAuthentic(user)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res
	}

	return res
}
