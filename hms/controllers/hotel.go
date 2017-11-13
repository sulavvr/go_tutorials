package controllers

import (
	// "log"
	"net/http"
	"strconv"
	"strings"

	// "github.com/sulavvr/hms/database"
	"github.com/sulavvr/hms/models"
)

type Hotel struct{}

func (hotel Hotel) Index(writer http.ResponseWriter, req *http.Request) {
	h := &models.Hotel{} // get hotel instance

	hotels := h.All(DB)
	data := make(map[string]interface{})
	data["hotels"] = hotels

	err := Templates.ExecuteTemplate(writer, "hotels.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (hotel Hotel) Show(writer http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(url[len(url)-1])

	h := &models.Hotel{}     // get hotel instance
	r := &models.Room{}      // get room instance
	_hotel := h.Find(id, DB) // get hotel
	rooms := r.Rooms(id, DB) // get all rooms for hotel

	data := make(map[string]interface{})
	data["hotel"] = _hotel
	data["rooms"] = rooms

	err := Templates.ExecuteTemplate(writer, "show.html", data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
