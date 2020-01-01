package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Language struct {
	ID string `json:"id,omitempty"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Paradigm string `json:"paradigm"`
}

var Lang []Language

func GetLanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range Lang {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Language{})
}

func GetLanguages(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Lang)
}

func AddLanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var lang Language
	_ = json.NewDecoder(req.Body).Decode(&lang)
	lang.ID = params["id"]
	Lang = append(Lang, lang)
	json.NewEncoder(w).Encode(Lang)
}

func DelLanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range Lang {
		if item.ID == params["id"] {
			Lang = append(Lang[:index], Lang[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Lang)
}

func main() {
	router := mux.NewRouter()
	Lang = append(Lang, Language{ID: "1", Creator: "Guido Van Rossum", Name: "Python", Paradigm: "Multi-paradigm: functional, imperative, object-oriented, reflective"})
	Lang = append(Lang, Language{ID: "2", Creator: "Brendan Eich", Name: "JavaScript", Paradigm: "Multi-paradigm: event-driven, functional, imperative, object-oriented (prototype-based)"})
	router.HandleFunc("/language", GetLanguages).Methods("GET")
	router.HandleFunc("/language/{id}", GetLanguage).Methods("GET")
	router.HandleFunc("/language/{id}", AddLanguage).Methods("POST")
	router.HandleFunc("/language/{id}", DelLanguage).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
