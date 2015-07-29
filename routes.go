package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"NewsIndex",
		"GET",
		"/news",
		NewsIndex,
	},
	Route{
		"NewsId",
		"GET",
		"/news/{newsId}",
		NewsShow,
	},
	Route{
		"NewsCreate",
		"POST",
		"/news",
		NewsCreate,
	},
}
