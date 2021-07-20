package reservation

import "github.com/Andre711/LibraryAPI/router"

var Routes = router.RoutePrefix{
	"/reservations",
	[]router.Route{
		router.Route{
			"AllReservations",
			"GET",
			"",
			allReservations,
		},
		router.Route{
			"CreateReservation",
			"POST",
			"/{UserID}/{BookID}",
			newReservation,
		},
		router.Route{
			"DeleteReservation",
			"DELETE",
			"/{UserID}/{BookID}",
			deleteReservation,
		},
	},
}
