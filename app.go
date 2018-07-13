package main

import "./views"
// import "./models"

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", views.AllRecipesEndPoint).Methods("GET")
	r.HandleFunc("/recipes", views.CreateRecipeEndPoint).Methods("POST")
	r.HandleFunc("/recipes", views.UpdateRecipeEndPoint).Methods("PUT")
	r.HandleFunc("/recipes", views.DeleteRecipeEndPoint).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", views.FindRecipeEndpoint).Methods("GET")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}