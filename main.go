package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id            int    `json : "id"`
	Name_Of_Movie string `json : "nameofmovie"`
	Released_Year int    `json : "realeaseyear"`
	Coloured      string `json : "coloured"`
	Watched       string `json : "watched"`
}

var Movies []Movie

func All_Movies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
	return
}

func Movie_By_ID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func Create_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func Update_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func Delete_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func main() {

	m1 := Movie{Id: 1, Name_Of_Movie: "IronMan", Released_Year: 2008, Coloured: "Yes", Watched: "Yes"}
	m2 := Movie{Id: 2, Name_Of_Movie: "IronMan_2", Released_Year: 2010, Coloured: "Yes", Watched: "Yes"}
	m3 := Movie{Id: 3, Name_Of_Movie: "IronMan_3", Released_Year: 2013, Coloured: "Yes", Watched: "Yes"}
	m4 := Movie{Id: 4, Name_Of_Movie: "Avengers", Released_Year: 2012, Coloured: "Yes", Watched: "Yes"}
	m5 := Movie{Id: 5, Name_Of_Movie: "Avengers : Age Of Altron", Released_Year: 2015, Coloured: "Yes", Watched: "Yes"}
	m6 := Movie{Id: 6, Name_Of_Movie: "Avengers : Infinity War", Released_Year: 2018, Coloured: "Yes", Watched: "Yes"}
	m7 := Movie{Id: 7, Name_Of_Movie: "Avengers : End Game", Released_Year: 2019, Coloured: "Yes", Watched: "Yes"}

	Movies = []Movie{m1, m2, m3, m4, m5, m6, m7}

	r := mux.NewRouter()
	r.HandleFunc("/api/movies", All_Movies).Methods("GET")
	r.HandleFunc("/api/movies/id", Movie_By_ID).Methods("GET")
	r.HandleFunc("/api/movies", Create_Movie).Methods("CREATE")
	r.HandleFunc("/api/movies/id", Update_Movie).Methods("UPDATE")
	r.HandleFunc("/api/movies/id", Delete_Movie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
