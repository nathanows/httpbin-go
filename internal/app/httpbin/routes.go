package httpbin

func (s *Server) initRoutes() {
	// HTTP Methods
	s.router.HandleFunc("/delete", s.handleDelete()).Methods("DELETE")
	s.router.HandleFunc("/get", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/patch", s.handlePatch()).Methods("PATCH")
	s.router.HandleFunc("/post", s.handlePost()).Methods("POST")
	s.router.HandleFunc("/put", s.handlePut()).Methods("PUT")

	// Anything Methods
	s.router.HandleFunc("/anything", s.handleAnything())
	s.router.HandleFunc("/anything/{anything}", s.handleAnything())

	// Status Codes Methods
	s.router.HandleFunc("/status/{codes}", s.handleStatusCodes())
}
