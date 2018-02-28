package main

import (
	"flag"
	"github.com/gotschmarcel/goserv"
	"log"
	"net/http"
	"webserver/Models"
)

type DataBase struct {
	Information []Models.User
}

var dblocation string
var table string

func init() {

	_dblocation := flag.String("location", "./Database/", "Address to the location of Database")
	_Table := flag.String("dbname", "database", "Target Table")

	flag.Parse()

	dblocation = *_dblocation
	table = *_Table
}

func main() {

	server := goserv.NewServer()

	// Handle Get Request
	server.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := &Models.User{}
		if err := goserv.WriteJSON(w, data); err != nil {
			goserv.Context(r).Error(err, http.StatusInternalServerError)
			return
		}
	})

	/**
	Handle Post Request
	Getting Data from the Database
	 */
	server.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var req Models.User

		// Read and decode the request's body
		if err := goserv.ReadJSONBody(r, &req); err != nil {
			log.Println("ERROR ... JSON READING")
			goserv.Context(r).Error(err, http.StatusBadRequest)
			return
		}

		// Quering the data
		var list DataBase
		var user Models.User

		list.Information = user.Get(dblocation, table, req.Email)

		for _,info :=range list.Information{
			log.Println("Email -> ",info.Email, "Name -> ",info.Name, "Surname -> ",info.Surname, "Phone -> ",info.Phone)
		}

		if err := goserv.WriteJSON(w, list); err != nil {
			goserv.Context(r).Error(err, http.StatusInternalServerError)
			return
		}

	})

	/**
	Handle Put Request
	Storing data in the database
	 */
	server.Put("/", func(w http.ResponseWriter, r *http.Request) {
		var batch DataBase
		var data Models.User

		/**
		It might you want to store couple of datas together, so that's why we are sending array of Users
		 */
		if err := goserv.ReadJSONBody(r, &batch); err != nil {
			log.Println("ERROR ... JSON READING")
			goserv.Context(r).Error(err, http.StatusBadRequest)
			return
		}

		// Adding to the Database
		for _,info := range batch.Information {
			if data.Add(dblocation, table, info) {
				data = info
				if err := goserv.WriteJSON(w, data); err != nil {
					goserv.Context(r).Error(err, http.StatusInternalServerError)
					return
				}

			}
		}
	})

	log.Fatalln(server.Listen(":60234"))
}

