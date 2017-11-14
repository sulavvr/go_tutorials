package controllers

import (
	// "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sulavvr/hms/models"
)

type Hotel struct {
	model models.Hotel
}

func (hotel Hotel) Index(writer http.ResponseWriter, req *http.Request) {
	hotels := hotel.model.All(DB)

	data["hotels"] = hotels

	err := Templates.ExecuteTemplate(writer, "hotels.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (hotel Hotel) Show(writer http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(url[len(url)-1])

	room := &models.Room{}             // get room instance
	_hotel := hotel.model.Find(id, DB) // get hotel
	rooms := room.Rooms(id, DB)        // get all rooms for hotel

	data["hotel"] = _hotel
	data["rooms"] = rooms

	err := Templates.ExecuteTemplate(writer, "hotel.html", data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
