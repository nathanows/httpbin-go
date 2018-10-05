package httpbin

func (s *Server) initRoutes() {
	// HTTP Routes
	s.router.HandleFunc("/delete", s.handleDelete()).Methods("DELETE")
	s.router.HandleFunc("/get", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/patch", s.handlePatch()).Methods("PATCH")
	s.router.HandleFunc("/post", s.handlePost()).Methods("POST")
	s.router.HandleFunc("/put", s.handlePut()).Methods("PUT")

	// Anything Routes
	s.router.HandleFunc("/anything", s.handleAnything())
	s.router.HandleFunc("/anything/{anything}", s.handleAnything())

	// Status Code Routes
	s.router.HandleFunc("/status/{codes}", s.handleStatusCodes())

	// Request Inspection Routes
	s.router.HandleFunc("/headers", s.handleHeaders()).Methods("GET")
	s.router.HandleFunc("/ip", s.handleIP()).Methods("GET")
	s.router.HandleFunc("/user-agent", s.handleUserAgent()).Methods("GET")

	// Auth Routes
	s.router.HandleFunc("/basic-auth/{user}/{password}", s.handleBasicAuth()).Methods("GET")
	s.router.HandleFunc("/bearer", s.handleBearer()).Methods("GET")
	s.router.HandleFunc("/hidden-basic-auth/{user}/{password}", s.handleBasicAuth()).Methods("GET")

	// Response Inspection Routes
	s.router.HandleFunc("/cache", s.handleCache()).Methods("GET")
	s.router.HandleFunc("/cache/{value:[0-9]+}", s.handleCacheControl()).Methods("GET")
	s.router.HandleFunc("/etag/{etag}", s.handleETag()).Methods("GET")
	s.router.HandleFunc("/response-headers", s.handleResponseHeaders()).Methods("GET", "POST")
}
