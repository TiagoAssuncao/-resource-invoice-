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
		"TodoIndex",
		"GET",
		"/users",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/invoices",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/users/{userId}",
		TodoShow,
	},
}
