package main

import "github.com/hsyodyssey/agoge/router"

func main() {

	r := router.SetupRouter()
	r.Run()

}
