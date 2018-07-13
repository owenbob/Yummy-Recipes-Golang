package models

import "gopkg.in/mgo.v2/bson"

// Recipes : model for creating a recipe
type Recipes struct{
	ID         bson.ObjectId  `bson:"_id" json:"id"`
	title       string         `bson:"_title" json:"title"`
	description string         `bson:"_description" json:"description"`
}