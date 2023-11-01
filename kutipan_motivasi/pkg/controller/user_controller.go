package controller

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/model"
)

type UserController struct {
	Users []model.User
}

func (uc *UserController) CreateUser(namaPengguna, email, kataSandi string) {
	newUser := model.NewUser(len(uc.Users)+1, namaPengguna, email, kataSandi)
	uc.Users = append(uc.Users, *newUser)
	fmt.Println("Pengguna baru berhasil dibuat.")
}

func (uc *UserController) GetUsers() []model.User {
	return uc.Users
}
