package controllers

import (
	"unairsatu_v2/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllModuls(c *gin.Context) {
	moduls := []bson.M{}
	cursor, _ := config.DB.Collection("modul").Find(c, bson.M{})
	cursor.All(c, &moduls)
	c.JSON(200, moduls)
}

func CreateModul(c *gin.Context) {
	var modul bson.M
	if err := c.BindJSON(&modul); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, _ := config.DB.Collection("modul").InsertOne(c, modul)
	c.JSON(201, result)
}

func UpdateModul(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var modul bson.M
	if err := c.BindJSON(&modul); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, _ := config.DB.Collection("modul").UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": modul})
	c.JSON(200, result)
}

func DeleteModul(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	result, _ := config.DB.Collection("modul").DeleteOne(c, bson.M{"_id": id})
	c.JSON(200, result)
}
