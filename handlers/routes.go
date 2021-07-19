package handlers

import "github.com/Andre711/LibraryAPI/router"

var Routes = router.RoutePrefix{
	"/users",
	[]router.Route{
		router.Route{
			"AllUsers",
			"GET",
			"",
			allUsers,
		},
		router.Route{
			"FindUser",
			"GET",
			"/{ID}",
			getUser,
		},
		router.Route{
			"CreateUser",
			"POST",
			"/{Name}/{Surname}",
			newUser,
		},
		router.Route{
			"DeleteUser",
			"DELETE",
			"/{ID}",
			deleteUser,
		},
		router.Route{
			"UpdateUser",
			"PUT",
			"/{ID}/{NewName}/{NewSurname}",
			updateUser,
		},
	},
}
