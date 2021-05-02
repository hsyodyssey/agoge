package main

import (
	"fmt"
	"os"

	"github.com/hsyodyssey/agoge/router"
)

func main() {

	switch os.Args[1] {
	case "p":
		r := router.SetupRouter()
		r.Run()
	case "h":
		fmt.Println("[Agoge] Please type p for starting...")
		fmt.Println("[Agoge] Program stop...")
	default:
		fmt.Println("[Agoge] Wrong input...")
		fmt.Println("[Agoge] Program stop...")

	}

}
