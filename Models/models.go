package Models

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"RESTapi/Util"
)

type User struct {
	Id      	int    `json:"id"`
	Email     	string `json:"Email"`
	Phone   	string `json:"Phone"`
	Name    	string `json:"Name"`
	Surname     string `json:"Surname"`
}

func (U User) Get(dblocation string, table string, email string) []User{

	var data []User

	log.Println(dblocation+table+".db")
	db, err := sql.Open("sqlite3", dblocation+table+".db")
	Util.Must(err)

	var query string
	query = fmt.Sprintf("select id as 'id',Email as 'Email',Phone as 'Phone',Name as 'Name',Surname as 'Surname' FROM User WHERE Email==\"%s\" ", email)
	rows, err := db.Query(query)
	Util.Must(err)

	for rows.Next() {
		err = rows.Scan(&U.Id, &U.Email, &U.Phone, &U.Name,&U.Surname)
		Util.Must(err)
		data = append(data, U)
	}
	rows.Close()

	return data
}

func (U User) Add(dblocation string, table string, user User) bool {

	log.Println(user)
	log.Println(dblocation+table+".db")
	db, err := sql.Open("sqlite3", dblocation+table+".db")
	Util.Must(err)

	query := fmt.Sprintf("INSERT into User(Email,Phone,Name,Surname) VALUES(\"%s\",\"%s\",\"%s\",\"%s\")", user.Email, user.Phone, user.Name, user.Surname)
	res, err := db.Exec(query)
	Util.Must(err)
	log.Println(res)

	affected, _ := res.RowsAffected()
	if affected != 1 {
		log.Fatalf("Expected %d for affected rows, but %d:", 1, affected)
		return false
	}

	return true
}
