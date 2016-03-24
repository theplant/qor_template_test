package main

import (
	"net/http"

	"github.com/theplant/qor_template_test/app"
)

func main() {
	mux := app.App()
	panic(http.ListenAndServe(":9999", mux))
}
