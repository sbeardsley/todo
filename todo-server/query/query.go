package query

import (
	"beards.ly/todo/repository"
	"github.com/sbeardsley/cqrs"
)

// Handlers a collection of query handlers
type Handlers []cqrs.QueryHandler

// NewHandlers constructor for all of the handlers needed by the system
func NewHandlers(userRepo repository.UserRepository) Handlers {

	return Handlers{
		NewUserByUsernameHandler(userRepo),
	}
}
