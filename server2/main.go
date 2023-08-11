package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func SetUpHandlers(mux *mux.Router) {
	fmt.Println("Starting Server2...")
	parentURL := "/server2"
	mux.HandleFunc(parentURL + "/movies", getMovies).Methods("GET")
	mux.HandleFunc(parentURL + "/movies/{id}", getMovie).Methods("GET")
	mux.HandleFunc(parentURL + "/movies", createMovie).Methods("POST")
	mux.HandleFunc(parentURL + "/movies/{id}", updateMovie).Methods("PUT")
	mux.HandleFunc(parentURL + "/movies/{id}", deleteMovie).Methods("DELETE")
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for index, movie := range movies {
		if movie.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for _, movie := range movies {
		if movie.ID == movieId {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Host, r.URL.Scheme)
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for index, item := range movies {
		if item.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			item.ID = movieId
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main(){
	router2 := mux.NewRouter()
	SetUpHandlers(router2)
	log.Fatal(http.ListenAndServe(":8082", router2))
}
