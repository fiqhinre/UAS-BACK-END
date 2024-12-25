package controllers

import (
    "github.com/gin-gonic/gin"
)

// Fungsi CRUD User
func GetAllUsers(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Get all users"})
}

func GetUserById(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"message": "Get user by ID", "id": id})
}

func CreateUser(c *gin.Context) {
    c.JSON(201, gin.H{"message": "User created"})
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"message": "User updated", "id": id})
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"message": "User deleted", "id": id})
}

// Fungsi Tambahan untuk Modul User
func GetUserModuls(c *gin.Context) {
    userId := c.Param("id")
    moduls := []string{"Modul 1", "Modul 2", "Modul 3"} // Simulasi data

    c.JSON(200, gin.H{
        "user_id": userId,
        "moduls":  moduls,
    })
}

func ChangeUserType(c *gin.Context) {
    userId := c.Param("id")
    newJenisUser := c.PostForm("jenis_user")

    // Logika untuk menghapus modul lama dan menambahkan modul baru

    c.JSON(200, gin.H{
        "message":       "Jenis user berhasil diubah",
        "user_id":       userId,
        "new_jenis_user": newJenisUser,
    })
}

func AddModulToUser(c *gin.Context) {
    userId := c.Param("id")
    modulId := c.PostForm("modul_id")

    // Logika untuk menambahkan modul

    c.JSON(201, gin.H{
        "message":  "Modul berhasil ditambahkan ke user",
        "user_id":  userId,
        "modul_id": modulId,
    })
}

func RemoveModulFromUser(c *gin.Context) {
    userId := c.Param("id")
    modulId := c.PostForm("modul_id")

    // Logika untuk menghapus modul

    c.JSON(200, gin.H{
        "message":  "Modul berhasil dihapus dari user",
        "user_id":  userId,
        "modul_id": modulId,
    })
}
