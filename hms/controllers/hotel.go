package controllers

import (
	"log"
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

func (hotel Hotel) Book(writer http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(url[len(url)-1])

	rate := &models.Rate{}
	_rate := rate.Find(id, DB)
	data["rate"] = _rate

	err := Templates.ExecuteTemplate(writer, "book.html", data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (hotel Hotel) Confirm(writer http.ResponseWriter, req *http.Request) {
	userDetails := map[string]string{
		"name":   req.FormValue("name"),
		"email":  req.FormValue("email"),
		"phone":  req.FormValue("phone"),
		"street": req.FormValue("street"),
		"state":  req.FormValue("state"),
		"zip":    req.FormValue("zip"),
	}
	user := &models.User{}
	uid, err := user.SetInfo(userDetails).Insert(DB)
	if err != nil {
		log.Fatal(err.Error())
	}

	r, _ := strconv.Atoi(req.FormValue("rate_id"))
	rate := &models.Rate{}
	_rate := rate.Find(r, DB)

	noGuests, _ := strconv.Atoi(req.FormValue("guests"))
	reservationDetails := map[string]interface{}{
		"user_id":   int(uid),
		"rate":      _rate,
		"check_in":  req.FormValue("check_in"),
		"check_out": req.FormValue("check_out"),
		"guests":    noGuests,
	}

	reservation := &models.Reservation{}
	_, err = reservation.SetInfo(reservationDetails).Insert(DB)
	if err != nil {
		log.Fatal(err.Error())
	}

	data["name"] = userDetails["name"]
	err = Templates.ExecuteTemplate(writer, "confirm.html", data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
