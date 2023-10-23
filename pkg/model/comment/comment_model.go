// comment_controller.go
package controller

type CommentController struct {
	CommentModel model.Comment
}

func (cc *CommentController) AddComment(quoteID, userID int, content, time string) {
	// Implement logic to add a comment
}
