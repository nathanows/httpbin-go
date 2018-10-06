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

	// Response Formats
	s.router.HandleFunc("/deny", s.handleDeny()).Methods("GET")
	s.router.HandleFunc("/encoding/utf8", s.handleEncodingUTF8()).Methods("GET")
	s.router.HandleFunc("/html", s.handleHTML()).Methods("GET")
	s.router.HandleFunc("/json", s.handleJSON()).Methods("GET")
	s.router.HandleFunc("/robots.txt", s.handleRobotsTxt()).Methods("GET")
	s.router.HandleFunc("/xml", s.handleXML()).Methods("GET")

	// Dynamic Data
	s.router.HandleFunc("/base64/{value}", s.handleBase64Decode()).Methods("GET")
	s.router.HandleFunc("/bytes/{n:[0-9]+}", s.handleBytes()).Methods("GET")
	s.router.HandleFunc("/delay/{delay:[0-9]+}", s.handleDelay())
	s.router.HandleFunc("/drip", s.handleDrip())
	s.router.HandleFunc("/links/{n:[0-9]+}/{offset:[0-9]+}", s.handleLinks()).Methods("GET")
	s.router.HandleFunc("/range/{numbytes:[0-9]+}", s.handleRange()).Methods("GET")
	s.router.HandleFunc("/stream-bytes/{n:[0-9]+}", s.handleStreamBytes()).Methods("GET")
	s.router.HandleFunc("/stream/{n:[0-9]+}", s.handleStream()).Methods("GET")
	s.router.HandleFunc("/uuid", s.handleUUID()).Methods("GET")

	// Cookies
	s.router.HandleFunc("/cookies", s.handleCookies()).Methods("GET")
	s.router.HandleFunc("/cookies/delete", s.handleCookiesDelete()).Methods("GET")
	s.router.HandleFunc("/cookies/set", s.handleCookiesSet()).Methods("GET")
	s.router.HandleFunc("/cookies/set/{name}/{value}", s.handleCookiesSet()).Methods("GET")
}
