package todo

import (
	"github.com/gigovich/fargo/contrib/auth"
	"github.com/gigovich/fargo/router"
)

var Router = router.New(
	router.URL("/", MainView),
	router.Include("/auth/", auth.Router),
)
