package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"api-server/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) FindPostsByCategory(ctx *gin.Context) {
	postCategory := ctx.Query("postCategory")
	fromDate := ctx.Query("fromDate")
	toDate := ctx.Query("toDate")

	fromDateTime, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		fmt.Println(err)
	}

	toDateTime, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		fmt.Println(err)
	}

	var posts []models.Post
	err = pc.DB.Where("category = ? AND  created_at >= ? AND created_at <= ?", postCategory, fromDateTime.Format(time.RFC3339Nano), toDateTime.Format(time.RFC3339Nano)).Find(&posts).Error
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": posts})
}

func (pc *PostController) Kek(ctx *gin.Context) {
	log.Println("kekec")
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": "kek"})
}
