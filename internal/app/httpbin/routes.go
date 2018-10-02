package httpbin

func (s *Server) initRoutes() {
	// HTTP Methods
	s.router.HandleFunc("/delete", s.HandleDelete()).Methods("DELETE")
}
