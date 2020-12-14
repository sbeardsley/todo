package controller

import (
	"beards.ly/todo/command"
	"beards.ly/todo/query"
	"github.com/gin-gonic/gin"
)

// CreateUser action
func (s *Server) CreateUser(c *gin.Context) {
	var createUser command.CreateUser
	err := c.BindJSON(&createUser)
	if err != nil {
		s.APIResult(c, nil, err)
		return
	}

	err = s.App.Dispatcher.Send(createUser)
	if err != nil {
		s.APIResult(c, nil, err)
		return
	}

	user, e := s.App.Dispatcher.SendQ(&query.UserByUsername{Username: createUser.Username})
	if e != nil {
		s.APIResult(c, user, e)
	} else {
		s.APIResult(c, user)
	}
}
