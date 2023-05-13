package model

import (
	"log"
	api_structure "main/internal/pkg/structures"
)

func Register(user *api_structure.User) *api_structure.Response {
	// kullanıcı setinde kullanıcı adı olup olmadığını kontrol edin
	// varsa hata döndür
	// yeni kullanıcı oluştur
	// hata için yanıt oluştur
	res := &api_structure.Response{Status: true}

	status, err := IsUserExist(user.Username)
	if err != nil {
		log.Fatal(err)
	}
	if status {
		res.Status = false
		res.Message = "username already taken. try something else."
		return res
	}

	err = RegisterNewUser(user)
	if err != nil {
		res.Status = false
		res.Message = "something went wrong while registering the user. please try again after sometime."
		return res
	}

	return res
}

func Login(user *api_structure.User) *api_structure.Response {
	// geçersiz kullanıcı adı ve şifre hatası verirse
	// geçerli kullanıcı ise yeni oturum oluştur
	res := &api_structure.Response{Status: true}

	err := IsUserAuthentic(user)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res
	}

	return res
}
