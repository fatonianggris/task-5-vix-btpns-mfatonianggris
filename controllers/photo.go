package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"net/http"
	"rakaminbtpn/database"
	"rakaminbtpn/helpers"
	"rakaminbtpn/models"
)

func List_photo(c *gin.Context) {
	var (
		photos []models.Photo
	)

	db := database.Get_db()
	err := db.Preload("User").Find(&photos).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &photos,
	})
}

func Create_photo(c *gin.Context) {
	db := database.Get_db()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.Get_content_type(c)
	_, _ = db, contentType

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func Update_photo(c *gin.Context) {
	db := database.Get_db()
	contentType := helpers.Get_content_type(c)
	_, _ = db, contentType

	Photo := models.Photo{}
	NewPhoto := models.Photo{}

	id := c.Param("userId")

	err := db.First(&Photo, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&NewPhoto)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"updated_at": Photo.UpdatedAt,
	})
}

func Delete_photo(c *gin.Context) {
	db := database.Get_db()
	contentType := helpers.Get_content_type(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	id := c.Param("photoId")

	err := db.First(&Photo, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
