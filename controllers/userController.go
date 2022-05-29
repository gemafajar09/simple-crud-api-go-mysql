package controllers

import (
	"net/http"

	"api-go-mysql/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUsers struct {
	Nama         string `json:"nama"`
	Telpon       string `json:"telpon"`
	Jeniskelamin string `json:"jeniskelamin"`
	Alamat       string `json:"alamat"`
}

type UpdateUsers struct {
	Nama         string `json:"nama"`
	Telpon       string `json:"telpon"`
	Jeniskelamin string `json:"jeniskelamin"`
	Alamat       string `json:"alamat"`
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func TambahUser(c *gin.Context) {
	var input CreateUsers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Nama: input.Nama, Telpon: input.Telpon, Jeniskelamin: input.Jeniskelamin, Alamat: input.Alamat}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CariUser(c *gin.Context) {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	var input UpdateUsers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.Jeniskelamin = input.Jeniskelamin
	updatedInput.Nama = input.Nama
	updatedInput.Telpon = input.Telpon
	updatedInput.Alamat = input.Alamat

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
