package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/sulavvr/hms/database"
)

type Controller interface {
	Index(writer http.ResponseWriter, req *http.Request)
	Show(writer http.ResponseWriter, req *http.Request)
}

var (
	Templates *template.Template
	DB        *sql.DB
	data      map[string]interface{}
)

func init() {
	Templates = template.Must(template.ParseGlob("views/*"))
	database.Setup()

	DB = database.Get()
	data = make(map[string]interface{})
}
