package main

import (
	"net/http"

	"github.com/Andre711/LibraryAPI/book"
	"github.com/Andre711/LibraryAPI/reservation"
	customRouter "github.com/Andre711/LibraryAPI/router"
	"github.com/Andre711/LibraryAPI/user"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//Init Router
	router := mux.NewRouter()

	//Append User Routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, user.Routes)

	//Append Book Routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, book.Routes)

	//Append Reservation Routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, reservation.Routes)

	for _, route := range customRouter.AppRoutes {

		//Create Subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//Loop Through Each Sub Route
		for _, r := range route.SubRoutes {

			var handler http.Handler = r.HandlerFunc

			// Attach Sub Route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	return router
}
