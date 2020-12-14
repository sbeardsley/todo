package controller

import (
	"net/http"

	"beards.ly/todo/app"
	"github.com/gin-gonic/gin"
)

// Server represents the web server
type Server struct {
	App    app.Application
	Router *gin.Engine
}

// NewServer constructor
func NewServer(a app.Application) *Server {
	server := &Server{
		App:    a,
		Router: gin.Default(),
	}

	return server
}

// Run server
func (s *Server) Run(addr string) {
	err := s.Router.Run(addr)
	if err != nil {
		panic(err)
	}
}

// APIResult response
func (s *Server) APIResult(c *gin.Context, result interface{}, errors ...error) {
	success := len(errors) == 0
	c.JSON(http.StatusOK, gin.H{"success": success, "result": result, "errors": errors})
}
