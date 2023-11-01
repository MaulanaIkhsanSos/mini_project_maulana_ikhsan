package main

import (
	"fmt"
	"time"

	"motivation-app/kutipan_motivasi/pkg/controller"
	"motivation-app/kutipan_motivasi/pkg/model"
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

	// Contoh pemanggilan fungsi dari controller
	quoteController.CreateQuote("Ini adalah kutipan motivasi", "2023-10-18")
	userController.CreateUser("username", "email@example.com", "password")

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

	quoteController = controller.QuoteController{}

	quoteController.CreateQuote("Ini adalah kutipan motivasi 1.", "2023-10-18")
	quoteController.CreateQuote("Ini adalah kutipan motivasi 2.", "2023-10-19")

	quotes := quoteController.GetQuotes()
	fmt.Println("Daftar Kutipan Motivasi:")
	for _, quote := range quotes {
		fmt.Printf("KutipanID: %d, Isi: %s, Tanggal: %s\n", quote.KutipanID, quote.IsiKutipan, quote.Tanggal)
	}

	userController = controller.UserController{}

	userController.CreateUser("Pengguna1", "user1@example.com", "password1")
	userController.CreateUser("Pengguna2", "user2@example.com", "password2")

	users := userController.GetUsers()
	fmt.Println("Daftar Pengguna:")
	for _, user := range users {
		fmt.Printf("UserID: %d, Nama Pengguna: %s, Email: %s\n", user.UserID, user.NamaPengguna, user.Email)
	}

	fmt.Println("Aplikasi sudah berjalan dengan baik!")

	select {}
}
