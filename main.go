package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	lis:=http.NewServeMux()

	lis.HandleFunc("/",homePage)

	lis.HandleFunc("/about",aboutPage)

	log.Println("The server has started on port 8080")

	err:=http.ListenAndServe(":8080",lis)

	if err!=nil{
		log.Fatal("There was a error starting the server as ",err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"You've landed in the home page")
}

func aboutPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"You've landed to the about me page")
}
