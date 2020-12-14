package query

import (
	"beards.ly/todo/repository"
)

// UserByUsername query
type UserByUsername struct {
	Username string
}

// User model
type User struct {
	ID       uint
	Username string
	Email    string
}

// UserByUsernameHandler handles UserByUsername
type UserByUsernameHandler struct {
	repo repository.UserRepository
}

// NewUserByUsernameHandler query handler constructor
func NewUserByUsernameHandler(repo repository.UserRepository) UserByUsernameHandler {
	return UserByUsernameHandler{
		repo: repo,
	}
}

// NewQuery used during registration of handler
func (h UserByUsernameHandler) NewQuery() interface{} {
	return &UserByUsername{}
}

// Handle UserByUsername
func (h UserByUsernameHandler) Handle(message interface{}) (interface{}, error) {

	qry := message.(*UserByUsername)
	user, err := h.repo.FindByUsername(qry.Username)
	if err != nil {
		return &User{}, err
	}

	result := &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return result, nil
}
