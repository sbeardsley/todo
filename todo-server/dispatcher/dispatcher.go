package dispatcher

import (
	"beards.ly/todo/command"
	"beards.ly/todo/query"
	"beards.ly/todo/repository"
	"github.com/sbeardsley/cqrs"
)

// NewSender constructs cqrs.Sender
func NewSender(userRepo repository.UserRepository) cqrs.Sender {
	sender := cqrs.NewInMemorySender()
	handlers := command.NewHandlers(userRepo)
	qhandlers := query.NewHandlers(userRepo)
	sender.Register(handlers...)
	sender.RegisterQ(qhandlers...)
	return sender
}
