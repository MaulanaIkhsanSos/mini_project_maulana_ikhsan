package main

import (
	"fmt"
	"motivation-app/pkg/controller"
	"motivation-app/pkg/model"
)

func main() {
	// Inisialisasi model-model yang diperlukan
	quoteModel := model.Quote{}
	userModel := model.User{}
	feedbackModel := model.Feedback{}
	commentModel := model.Comment{}
	motivationSenderModel := model.AutomatedMotivationSender{}

	// Inisialisasi controller-controller yang diperlukan
	quoteController := controller.QuoteController{QuoteModel: quoteModel}
	userController := controller.UserController{UserModel: userModel}
	feedbackController := controller.FeedbackController{FeedbackModel: feedbackModel}
	commentController := controller.CommentController{CommentModel: commentModel}
	motivationSenderController := controller.MotivationSenderController{SenderModel: motivationSenderModel}

	// Contoh pemanggilan fungsi dari controller
	quoteController.CreateQuote("Ini adalah kutipan motivasi", "2023-10-18")
	userController.CreateUser("username", "email@example.com", "password")

	// Implementasi logika lainnya
	fmt.Println("Aplikasi sudah berjalan dengan baik!")
}
