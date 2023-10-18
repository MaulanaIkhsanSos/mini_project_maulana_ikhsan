// feedback_controller.go
package controller

import "yourapp/pkg/model"

type FeedbackController struct {
	FeedbackModel model.Feedback
}

func (fc *FeedbackController) AddFeedback(quoteID, userID int, like, dislike bool) {
	// Implement logic to add feedback
}
