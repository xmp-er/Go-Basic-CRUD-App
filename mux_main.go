package main

import (
	"strconv"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
        "github.com/gorilla/mux"
)

const port string = ":8080"


type article struct{
	Title string `json:"title"`
	Body string `json:"body"`
}

var blogs = make(map[int]article)

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
			fmt.Fprintln(w,"There was a error reading the data",err)
		}

		body_as_string := string(body) //making the byte slice as string
		

		//next step is unmarshalling the string into json

		var data article

		err = json.Unmarshal([]byte(body_as_string),&data)

		if err!=nil{
			fmt.Fprintln(w,"There was a error unmarshalling the data",err)
		}

		id,_:= strconv.Atoi(mux.Vars(r)["id"])

		for k,_:= range blogs{
			if k == id{
				blogs[id]=data
				fmt.Fprintln(w,"The article with id ",id, " has been update and now has the title", blogs[id].Title," and the content has been updated to ",blogs[id].Body)
				return
			}
		}


        }
        if r.Method == http.MethodPut {
                fmt.Fprintln(w,"You've hit a PUT method inside the handler",blogs)

		data,err:=ioutil.ReadAll(r.Body)

		if err!=nil{
			fmt.Fprintln(w,"There was a error reading the data",err)
		}
		data_as_string := string(data)

		var data_final article

		json.Unmarshal([]byte(data_as_string),&data_final)

		id,_:=strconv.Atoi(mux.Vars(r)["id"])

		if len(blogs)==0{
			blogs[id] = data_final
			fmt.Fprintln(w,"The first data has been inserted")
		}else{
			for k,_:= range blogs{
			     if k==id{
				fmt.Fprintln(w,"The id already has data,please put a POST request instead")
				return
			     }
		        }
		        blogs[id] = data_final
		        fmt.Fprintln(w,"The body with id = ",id," has been updated to",blogs[id].Title," and ",blogs[id].Body)
		}

		fmt.Fprintln(w,blogs)
        }
        if r.Method == http.MethodDelete {
                fmt.Fprintln(w,"You've hit a DELETE method inside the handler",blogs)

		data,_:=ioutil.ReadAll(r.Body)

		data_as_string:=string(data)

		var data_final article

		json.Unmarshal([]byte(data_as_string),&data_final)

		id,_:=strconv.Atoi(mux.Vars(r)["id"])

		for i,_:=range blogs{
			if id==i{
				delete(blogs,id)
				fmt.Fprintln(w,"The data for the given id has been deleted")
				return
			}
		}
		fmt.Fprintln(w,"There was no article in the database matching this id.")

        }


}
