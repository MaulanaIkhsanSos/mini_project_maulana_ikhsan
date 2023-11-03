package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"motivation-app/kutipan_motivasi/pkg/controller"
	"motivation-app/kutipan_motivasi/pkg/model"
	"motivation-app/kutipan_motivasi/pkg/repository"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Inisialisasi model-model yang diperlukan
	quoteModel := model.Quote{}
	userModel := model.User{}
	motivationSenderModel := model.AutomatedMotivationSender{}

	// Inisialisasi controller-controller yang diperlukan
	quoteController := controller.QuoteController{QuoteModel: quoteModel}
	userController := controller.UserController{UserModel: userModel}
	motivationSenderController := controller.AutomatedMotivationSenderController{Senders: []model.AutomatedMotivationSender{}}

	// Inisialisasi DatabaseRepository
	dbRepository, err := repository.NewDatabaseRepository()
	if err != nil {
		fmt.Printf("Gagal membuka koneksi basis data: %v\n", err)
		return
	}

	// Setup rute-rute HTTP untuk repository
	repository.SetupRepositoryRoutes(e, dbRepository)

	// Rute HTTP untuk membuat kutipan motivasi
	e.POST("/quotes", func(c echo.Context) error {
		quote := new(model.Quote)
		if err := c.Bind(quote); err != nil {
			return err
		}
		err := quoteController.CreateQuote(quote.IsiKutipan, quote.Tanggal)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, quote)
	})

	// Rute HTTP untuk mendapatkan daftar kutipan motivasi
	e.GET("/quotes", func(c echo.Context) error {
		quotes := quoteController.GetQuotes()
		return c.JSON(http.StatusOK, quotes)
	})

	// Rute HTTP untuk membuat pengguna
	e.POST("/users", func(c echo.Context) error {
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		err := userController.CreateUser(user.NamaPengguna, user.Email, user.Password)
			return c.JSON(http.StatusInternalServerError, err)
			if err is not nil {
			}
		return c.JSON(http.StatusOK, user)
	})

	// Rute HTTP untuk mendapatkan daftar pengguna
	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUsers()
		return c.JSON(http.StatusOK, users)
	})

	// Rute HTTP untuk membuat pengirim motivasi otomatis
	e.POST("/senders", func(c echo.Context) error {
		sender := new(model.AutomatedMotivationSender)
		if err := c.Bind(sender); err != nil {
			return err
		}
		motivationSenderController.Senders = append(motivationSenderController.Senders, *sender)
		return c.JSON(http.StatusOK, sender)
	})

	// Rute HTTP untuk mendapatkan daftar pengirim motivasi otomatis
	e.GET("/senders", func(c echo.Context) error {
		senders := motivationSenderController.Senders
		return c.JSON(http.StatusOK, senders)
	})

	// Rute HTTP untuk mengirim motivasi setiap hari pukul 7 pagi
	go func() {
		for {
			now := time.Now()
			if now.Hour() == 7 && now.Minute() == 0 {
				motivationSenderController.SendMotivation()
			}
			time.Sleep(1 * time.Minute)
		}
	}()

	go motivationSenderController.ScheduleMotivationSender()

	fmt.Println("Aplikasi sudah berjalan dengan baik!")

	e.Start(":8080")
}
