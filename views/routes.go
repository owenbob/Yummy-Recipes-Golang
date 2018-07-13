package views


import (
	"encoding/json"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"../models"
	"../dao"
	"../config"
)

var configAccess = config.Config{}
var dataAccess = dao.RecipesDAO{}
var sucessMessage string = "Succes"


// AllRecipesEndPoint : endpoint for viewing all recipes
func AllRecipesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// FindRecipeEndpoint : endpoint for viewing one recipe
func FindRecipeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// CreateRecipeEndPoint : endpoint for creating a recipe
func CreateRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe models.Recipes
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	recipe.ID = bson.NewObjectId()
	if err := dataAccess.Insert(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, recipe )
}

// UpdateRecipeEndPoint : endpoint for editing a recipe
func UpdateRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
// DeleteRecipeEndPoint : endpoint for deleting a recipe
func DeleteRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// RespondWithError : Function to return with error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// RespondWithJson : Function to return with json
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	configAccess.Read()

	dataAccess.Server = configAccess.Server
	dataAccess.Database = configAccess.Database
	dataAccess.Connect()
}
