package server

import (
	"net/http"

	"github.com/sulavvr/hms/controllers"
)

var (
	hotelController controllers.Controller
)

func init() {
	hotelController = &controllers.Hotel{}
}

/**
 * Start handles the routes and also starts the server to serve the application.
 */
func Start() {
	http.HandleFunc("/", hotelController.Index)
	http.HandleFunc("/hotels/", hotelController.Show)

	http.ListenAndServe(":3000", nil)
}
