package main 

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type movie struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director`
}

type Director struct {
	FirstName string `json:"firstname`
	LastName string `json:"lastname`
}

var movies []movie

func getmovies (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index,item:= range movies {
		if item.Id == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _,item:= range movies {
		if item.Id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createmovie (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemovie (w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for index,item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]... )
			var movie movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main(){
	r:= mux.NewRouter()

	movies = append(movies, movie{Id: "1",Isbn: "43227",Title: "Movie one",Director: &Director{FirstName: "John",LastName: "Doe"}})
	movies = append(movies, movie{Id: "2",Isbn: "45469",Title: "Movie Two",Director: &Director{FirstName: "Sam",LastName: "Mendes"}})
	movies = append(movies, movie{Id: "3",Isbn: "56883",Title: "Movie Three",Director: &Director{FirstName: "Ron",LastName: "Howard"}})
	movies = append(movies, movie{Id: "4",Isbn: "97655",Title: "Movie Four",Director: &Director{FirstName: "Tim",LastName: "Burton"}})
	movies = append(movies, movie{Id: "5",Isbn: "73182",Title: "Movie Five",Director: &Director{FirstName: "Ang",LastName: "Lee"}})
	movies = append(movies, movie{Id: "6",Isbn: "45927",Title: "Movie Six",Director: &Director{FirstName: "James",LastName: "Cameron"}})
	movies = append(movies, movie{Id: "7",Isbn: "28374",Title: "Movie Seven",Director: &Director{FirstName: "Peter",LastName: "Jackson"}})
	movies = append(movies, movie{Id: "8",Isbn: "60419",Title: "Movie Eight",Director: &Director{FirstName: "Spike",LastName: "Lee"}})
	movies = append(movies, movie{Id: "9",Isbn: "15738",Title: "Movie Nine",Director: &Director{FirstName: "Clint",LastName: "Eastwood"}})
	movies = append(movies, movie{Id: "10",Isbn: "91056",Title: "Movie Ten",Director: &Director{FirstName: "David",LastName: "Flincher"}})

	r.HandleFunc("/movies",getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getmovie).Methods("GET")
	r.HandleFunc("/movies",createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deletemovie).Methods("DELETE")

	fmt.Printf("starting the server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}
