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
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"healthcheck",
		"GET",
		"/health",
		HealthCheck,
	},
	Route{
		"GetPeople",
		"GET",
		"/people",
		GetPeople,
	},
	Route{
		"GetPerson",
		"GET",
		"/people/{id}",
		GetPerson,
	},
	Route{
		"CreatePerson",
		"POST",
		"/people/{id}",
		CreatePerson,
	},
	Route{
		"DeletePerson",
		"DELETE",
		"/people/{id}",
		DeletePerson,
	},



}
