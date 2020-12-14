package controller

// InitializeRoutes configures routes
func (s *Server) InitializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{
		// User routes
		v1.POST("/user", s.CreateUser)
	}
}
