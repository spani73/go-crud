package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
    ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter , r *http.Request){
	
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1" , Isbn: "438277" , Title: "Race 2" , Director: &Director{Firstname: "John" , Lastname: "Doe"}} )
	movies = append(movies, Movie{ID: "2" , Isbn: "418277" , Title: "Race 3" , Director: &Director{Firstname: "Sanjay" , Lastname: "Bhansali"}} )
	


	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}