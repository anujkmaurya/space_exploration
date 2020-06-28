package routes

import (
	"net/http"
)

func Init() {

	// Handle routes
	http.Handle("/", InitHandlers())
}
