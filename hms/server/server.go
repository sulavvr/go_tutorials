package server

import (
	"html/template"
	// "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sulavvr/hms/database"
	"github.com/sulavvr/hms/models"
	// "github.com/sulavvr/hms/models/room"
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseGlob("views/*"))
	database.Setup()
}

func Start() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hotels/", show)

	http.ListenAndServe(":3000", nil)
}

func index(writer http.ResponseWriter, req *http.Request) {
	db := database.Get()
	// defer db.Close()

	h := &models.Hotel{} // get hotel instance

	hotels := h.All(db)
	data := make(map[string]interface{})
	data["hotels"] = hotels

	err := templates.ExecuteTemplate(writer, "hotels.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func show(writer http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(url[len(url)-1])
	db := database.Get()
	// defer db.Close()

	h := &models.Hotel{}     // get hotel instance
	r := &models.Room{}      // get room instance
	hotel := h.Find(id, db)  // get hotel
	rooms := r.Rooms(id, db) // get all rooms for hotel

	data := make(map[string]interface{})
	data["hotel"] = hotel
	data["rooms"] = rooms

	err := templates.ExecuteTemplate(writer, "show.html", data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
