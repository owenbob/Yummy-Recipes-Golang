package dao
import "../models"

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


// RecipesDAO : Recipes   Data Acess Object
type RecipesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

//Recipes collection
const (
	COLLECTION = "recipes"
)

// Connect : Database connection
func (m *RecipesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll : list of recipes
func (m *RecipesDAO) FindAll() ([]models.Recipes, error) {
	var recipes []models.Recipes
	err := db.C(COLLECTION).Find(bson.M{}).All(&recipes)
	return recipes, err
}

// FindById : a recipe by its id
func (m *RecipesDAO) FindById(id string) (models.Recipes, error) {
	var recipe models.Recipes
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&recipe)
	return recipe, err
}

// Insert : a recipe into database
func (m *RecipesDAO) Insert(recipe models.Recipes) error {
	err := db.C(COLLECTION).Insert(&recipe)
	return err
}

// Delete : an existing recipe
func (m *RecipesDAO) Delete(recipe models.Recipes) error {
	err := db.C(COLLECTION).Remove(&recipe)
	return err
}

// // Update an existing recipe
// func (m *RecipesDAO) Update(recipe models.Recipes) error {
// 	err := db.C(COLLECTION).UpdateId(recipe.id, &recipe)
// 	return err
// }