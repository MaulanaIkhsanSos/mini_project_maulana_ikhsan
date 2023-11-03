package controller

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Users []model.User
}

func (uc *UserController) CreateUser(c echo.Context) error {
	namaPengguna := c.FormValue("namaPengguna")
	email := c.FormValue("email")
	kataSandi := c.FormValue("kataSandi")

	newUser := model.NewUser(len(uc.Users)+1, namaPengguna, email, kataSandi)
	uc.Users = append(uc.Users, *newUser)
	fmt.Println("Pengguna baru berhasil dibuat.")
	return c.JSON(http.StatusOK, newUser)
}

func (uc *UserController) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, uc.Users)
}

func SetupUserRoutes(e *echo.Echo, uc *UserController) {
	e.POST("/users", uc.CreateUser)
	e.GET("/users", uc.GetUsers)
}
