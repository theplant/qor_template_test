package app

import (
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/qor"
)

func App() *http.ServeMux {
	admin := admin.New(&qor.Config{DB: nil})

	mux := http.NewServeMux()
	admin.MountTo("/", mux)
	return mux
}
