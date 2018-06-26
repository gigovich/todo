package main

import (
	"github.com/gigovich/fargo/app"
)

func main() {
	if err := app.Register(App); err != nil {
		fmt.Println(err)
	}

	if err := App.Run(); err != nil {
		fmt.Println(err)
	}
}
