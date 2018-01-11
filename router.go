package appeng

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter factory
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.Methods(route.Method...).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

//Route struct
type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes array
type Routes []Route

var routes = Routes{
	Route{
		"mail",
		[]string{"POST"},
		"/_ah/mail/anyone@yoyo-dot-strange-mariner-191706.appspotmail.com",
		incomingMail,
	},
	Route{
		"startvm",
		[]string{"GET"},
		"/start/{zone}/{servername}",
		startvmHandler,
	},
	Route{
		"stopvm",
		[]string{"GET", "POST"},
		"/stop/{zone}/{servername}",
		stopvmHandler,
	},
}
