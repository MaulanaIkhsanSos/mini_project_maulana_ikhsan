package main

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/controller"
	"motivation-app/kutipan_motivasi/pkg/model"
	"motivation-app/kutipan_motivasi/pkg/service"
	"time"
)

func main() {
	// Inisialisasi model-model yang diperlukan
	quoteModel := model.Quote{}
	userModel := model.User{}
	motivationSenderModel := model.AutomatedMotivationSender{}

	// Inisialisasi controller-controller yang diperlukan
	quoteController := controller.QuoteController{QuoteModel: quoteModel}
	userController := controller.UserController{UserModel: userModel}
	motivationSenderController := controller.MotivationSenderController{SenderModel: motivationSenderModel}

	// Konfigurasi email sender untuk MotivationService
	emailSender := &service.SimpleEmailSender{
		SMTPServer:   "smtp.example.com",
		SMTPPort:     "587",
		SMTPUsername: "your_username",
		SMTPPassword: "your_password",
	}

	// Inisialisasi MotivationService
	motivationService := service.NewMotivationService(motivationSenderModel, emailSender)

	// Contoh pemanggilan fungsi dari controller
	quoteController.CreateQuote("Ini adalah kutipan motivasi", "2023-10-18")
	userController.CreateUser("username", "email@example.com", "password")

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

	// Implementasi logika lainnya
	fmt.Println("Aplikasi sudah berjalan dengan baik!")

	// Jangan keluar dari aplikasi
	select {}
}
