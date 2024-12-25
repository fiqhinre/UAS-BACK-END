package controllers

import (
	"context"
	"net/http"
	"unairsatu_v2/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk menambahkan modul ke template_modul
func AddTemplateModul(c *gin.Context) {
	var template bson.M
	if err := c.BindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tambahkan data ke collection template_moduls
	result, err := config.DB.Collection("template_moduls").InsertOne(context.Background(), template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Template modul added",
		"id":      result.InsertedID,
	})
}

// Fungsi untuk menghapus modul dari template_modul
func RemoveTemplateModul(c *gin.Context) {
	id := c.Param("id")

	// Ubah string ID ke ObjectID MongoDB
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Hapus data dari collection template_moduls
	result, err := config.DB.Collection("template_moduls").DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Template modul removed",
		"deleted": result.DeletedCount,
	})
}
