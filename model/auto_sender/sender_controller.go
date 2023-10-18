// motivation_sender_controller.go
package controller

import "yourapp/pkg/model"

type MotivationSenderController struct {
	SenderModel model.AutomatedMotivationSender
}

func (msc *MotivationSenderController) SendMotivation() {
	// Implement logic to send automated motivation every morning
}
