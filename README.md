# Building up a RESTapi

This is a sample program which stores data in the SQLite database. Access to this data is done via http requests. 

- Post Request = Query the user inside database
- Put Request = Add a new user to the database

 
# How to Install dependencies
you may need first need to install glide!
```
glide up

```
# How to Run 

You can change database location or name of the database via flags 
```
go run main.go -location="./Database/"  -dbname="database"
```
Now Server is listening on port: 60234

#How to test 
Open Postman to send GET / POST / PUT requests! 

## Sample
### GET request

Returns an empty json format 
```
URL : http://localhost:60234/
HEADER : accept application/json
```
### POST request

Return the complete information about the user if it's stored in the database
```
URL : http://localhost:60234/
HEADER : accept application/json
BODY :
{
		"Email": "niloofar.gheibi@gmail.com",
		"Phone": "",
		"Name": "",
		"Surname": ""
}
```


### PUT request
Inserts new user inside the database

```
URL : http://localhost:3000/ 
HEADER : accept application/json
BODY :
{
	"Information": [{
		"Email": "new.user@sample.com",
		"Phone": "123456789",
		"Name": "new",
		"Surname": "user"
	}]
}
```
