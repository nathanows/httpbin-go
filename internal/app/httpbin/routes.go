package httpbin

func (s *Server) initRoutes() {
	s.router.HandleFunc("/something", s.handleSomething())
}
