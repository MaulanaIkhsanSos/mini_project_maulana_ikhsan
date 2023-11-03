package main

import (
	"fmt"
	"net/http"
	"time"

	"motivation-app/kutipan_motivasi/pkg/controller"
	"motivation-app/kutipan_motivasi/pkg/model"
	"motivation-app/kutipan_motivasi/pkg/service"
	"motivation-app/kutipan_motivasi/pkg/util" // Import util package

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	motivationSenderController := controller.MotivationSenderController{SenderModel: motivationSenderModel}

	// Konfigurasi email sender untuk MotivationService
	emailSender := util.NewSimpleEmailSender("smtp.example.com", "587", "your_username", "your_password")

	// Inisialisasi MotivationService
	motivationService := service.NewMotivationService(motivationSenderModel, emailSender)

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
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, user)
	})

	// Rute HTTP untuk mendapatkan daftar pengguna
	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUsers()
		return c.JSON(http.StatusOK, users)
	})

	// Rute HTTP untuk mengirim email
	e.POST("/send-email", func(c echo.Context) error {
		email := new(util.EmailMessage)
		if err := c.Bind(email); err != nil {
			return err
		}
		err := emailSender.SendEmail(email.Subject, email.Body, email.To)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, "Email sent successfully")
	})

	// Implementasi pengiriman motivasi otomatis
	go func() {
		for {
			now := time.Now()
			if now.Hour() == 7 && now.Minute() == 0 {
				motivationService.SendMotivation()
			}
			time.Sleep(1 * time.Minute)
		}
	}()

	fmt.Println("Aplikasi sudah berjalan dengan baik!")

	// Mulai server Echo pada port 8080
	e.Start(":8080")
}
