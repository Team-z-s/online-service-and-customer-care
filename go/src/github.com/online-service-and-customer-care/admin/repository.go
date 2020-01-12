package comment

import (
	"gitlab.com/username/online-service-and-customer-care/entity"
)

// AdminRepository specifies admin related related database operations
type CommentRepository interface {
	Comments() ([]entity.Comment, []error)
	Comment(id uint) (*entity.Comment, []error)
	UpdateComment(comment *entity.Comment) (*entity.Comment, []error)
	DeleteComment(id uint) (*entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
}
