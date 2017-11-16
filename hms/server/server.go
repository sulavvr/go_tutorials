package server

import (
	"net/http"

	"github.com/sulavvr/hms/controllers"
)

var (
	HotelController controllers.Hotel
	UserController  controllers.User
)

// func init() {
// 	HotelController = &controllers.Hotel{}
// 	UserController = &controllers.User{}
// }

/**
 * Start handles the routes and also starts the server to serve the application.
 */
func Start() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", HotelController.Index)
	http.HandleFunc("/hotels/", HotelController.Index)
	http.HandleFunc("/hotel/", HotelController.Show)
	http.HandleFunc("/hotel/book/", HotelController.Book)
	http.HandleFunc("/hotel/book/confirm", HotelController.Confirm)
	http.HandleFunc("/users/", UserController.Index)
	http.HandleFunc("/user/", UserController.Show)

	http.ListenAndServe(":3000", nil)
}
