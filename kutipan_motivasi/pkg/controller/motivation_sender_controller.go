package controller

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/model"

	"github.com/robfig/cron"
)

type MotivationSenderController struct {
	SenderModel model.AutomatedMotivationSender
}

func (msc *MotivationSenderController) SendMotivation() {

	fmt.Println("Mengirimkan kata-kata motivasi kepada pengguna pada pukul 7 pagi.")

	msc.SenderModel.IsiMotivasi = "Selamat pagi! Semangat hari ini!"
}

func (msc *MotivationSenderController) ScheduleMotivationSender() {
	c := cron.New()
	c.AddFunc("0 7 * * *", msc.SendMotivation)
	c.Start()
}
