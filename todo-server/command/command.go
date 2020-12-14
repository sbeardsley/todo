package command

import (
	"beards.ly/todo/repository"
	"github.com/sbeardsley/cqrs"
)

// Handlers a collection of handlers
type Handlers []cqrs.CommandHandler

// NewHandlers constructor for all of the handlers needed by the system
func NewHandlers(userRepo repository.UserRepository) Handlers {

	return Handlers{
		NewCreateUserHandler(userRepo),
	}
}
