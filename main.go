package main

import (
	"net/http"

	"github.com/MasterChanMonkey/go-mega-code/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
