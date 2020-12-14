package command

import (
	"beards.ly/todo/model"
	"beards.ly/todo/repository"
)

// CreateUser command
type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserHandler handles CreateUser
type CreateUserHandler struct {
	repo repository.UserRepository
}

// NewCreateUserHandler command handler constructor
func NewCreateUserHandler(repo repository.UserRepository) CreateUserHandler {
	return CreateUserHandler{
		repo: repo,
	}
}

// NewCommand used during registration of handler
func (h CreateUserHandler) NewCommand() interface{} {
	return &CreateUser{}
}

// Handle CreateUser
func (h CreateUserHandler) Handle(message interface{}) error {

	cmd := message.(*CreateUser)
	user := model.User{
		Username: cmd.Username,
		Email:    cmd.Email,
		Password: cmd.Password,
	}
	h.repo.Save(user)

	return nil
}
