package main

import (
	"sample-app/pkg"
)

func main() {

	r := pkg.SetupRouter()
	r.Run(":3000")

}
