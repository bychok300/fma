package controllers

import "awesomeProject/api/middlewares"

func (s *Server) initializeRoutes() {

	//TODO move urls to configs

	// Root Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Root)).Methods("GET")

	// Posts Route
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.AllPosts)).Methods("POST")
	s.Router.HandleFunc("/createPost", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/ratePost", middlewares.SetMiddlewareJSON(s.RatePost)).Methods("POST")

}
