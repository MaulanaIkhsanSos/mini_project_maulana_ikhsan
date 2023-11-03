package controller

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron"
)

type MotivationSenderController struct {
	SenderModel model.AutomatedMotivationSender
}

func (msc *MotivationSenderController) SendMotivation(c echo.Context) error {
	fmt.Println("Mengirimkan kata-kata motivasi kepada pengguna pada pukul 7 pagi.")
	msc.SenderModel.IsiMotivasi = "Selamat pagi! Semangat hari ini!"
	return c.JSON(http.StatusOK, map[string]string{"message": "Motivasi telah dikirim"})
}

func (msc *MotivationSenderController) ScheduleMotivationSender() {
	c := cron.New()
	c.AddFunc("0 7 * * *", func() {
		msc.SendMotivation
	})
	c.Start()
}

// Tambahkan rute-rute HTTP
func SetupMotivationRoutes(e *echo.Echo, msc *MotivationSenderController) {
	e.GET("/send-motivation", msc.SendMotivation)
}
