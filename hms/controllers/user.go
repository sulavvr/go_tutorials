package controllers

import (
	// "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sulavvr/hms/models"
)

type User struct {
	model models.User
}

func (user User) Index(writer http.ResponseWriter, req *http.Request) {
	users := user.model.All(DB)

	data["users"] = users

	err := Templates.ExecuteTemplate(writer, "users.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (user User) Show(writer http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(url[len(url)-1])

	_user, err := user.model.Find(id, DB)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		Templates.ExecuteTemplate(writer, "404.html", err.Error())
		return
	}

	reservation := &models.Reservation{}
	data["user"] = _user
	data["reservations"] = reservation.Reservations(id, DB)
	// log.Print(data["reservations"])

	err = Templates.ExecuteTemplate(writer, "user.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
