package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
        "github.com/gorilla/mux"
)

const port string = ":8080"

var blogs = make(map[int]string)

type article struct{
	Title string `json:"title"`
	Body string `json:"body"`
}

func main(){

	lis:=mux.NewRouter() //using the gorilla mux to initialize the router

	lis.HandleFunc("/",homePage)

	s:= lis.PathPrefix("/data/").Subrouter()

	s.HandleFunc("/{id}",RESTHandler).Methods("GET","POST","PUT","DELETE") //a single method to handle all the requests coming to this function

	log.Println("Starting the server at port",port)

	err:=http.ListenAndServe(port,lis)

	if err!=nil{
		log.Fatal("Error has occured while starting the server")
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"This is the about me page, use /data/ endpoints")
}

func RESTHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		fmt.Fprintln(w,"You've hit a GET method inside the handler",blogs)
	}
	if r.Method == http.MethodPost {
	        body,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			fmt.Println("There is a error with the request, ",err)
		}

		body_as_string := string(body) //converting the byte slice to string
		var data article

		err = json.Unmarshal([]byte(body_as_string),&data)

		if err!=nil{
			fmt.Println("There was a error unmarshalling the json data",err)
		}

		fmt.Fprintln(w,data," has the title ",data.Title," and the body ",data.Body)
        }
        if r.Method == http.MethodPut {
                fmt.Fprintln(w,"You've hit a PUT method inside the handler",blogs)
        }
        if r.Method == http.MethodDelete {
                fmt.Fprintln(w,"You've hit a DELETE method inside the handler",blogs)
        }


}
