package httpbin

func (s *Server) initRoutes() {
	// HTTP Methods
	s.router.HandleFunc("/delete", s.handleDelete()).Methods("DELETE")
}
