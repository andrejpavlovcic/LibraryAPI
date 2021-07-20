package book

import "github.com/Andre711/LibraryAPI/router"

var Routes = router.RoutePrefix{
	"/books",
	[]router.Route{
		router.Route{
			"AllBooks",
			"GET",
			"/all",
			allBooks,
		},
		router.Route{
			"FindBook",
			"GET",
			"/avaible",
			avaibleBooks,
		},
		router.Route{
			"FindBook",
			"GET",
			"/{ID}",
			getBook,
		},
		router.Route{
			"CreateBook",
			"POST",
			"/{Title}/{Stock}",
			newBook,
		},
		router.Route{
			"DeleteBook",
			"DELETE",
			"/{ID}",
			deleteBook,
		},
		router.Route{
			"UpdateBookStock",
			"PUT",
			"/{ID}/{Stock}",
			updateBookStock,
		},
	},
}
