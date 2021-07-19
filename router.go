package main

import (
	"net/http"

	"github.com/Andre711/LibraryAPI/handlers"
	customRouter "github.com/Andre711/LibraryAPI/router"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	/* Init Router */
	router := mux.NewRouter()

	/* Append User Routes */
	customRouter.AppRoutes = append(customRouter.AppRoutes, handlers.Routes)

	for _, route := range customRouter.AppRoutes {

		/* Create Subroute */
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		/* Loop Through Each Sub Route */
		for _, r := range route.SubRoutes {

			var handler http.Handler = r.HandlerFunc

			/* Attach Sub Route */
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	return router
}
