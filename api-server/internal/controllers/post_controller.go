package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"api-server/internal/models"

	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PostController struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewPostController(DB *gorm.DB, Cache *redis.Client) PostController {
	return PostController{DB, Cache}
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
	cacheKey := fmt.Sprintf("%s:%s:%s", postCategory, fromDate, toDate)
	val, err := pc.Cache.Get(ctx, cacheKey).Result()
	if err != nil {
		log.Println(err)
		err = pc.DB.Where("category = ? AND  created_at >= ? AND created_at <= ?", postCategory, fromDateTime.Format(time.RFC3339Nano), toDateTime.Format(time.RFC3339Nano)).Find(&posts).Error
		if err != nil {
			log.Println(err)
		}
	} else {
		err = json.Unmarshal([]byte(val), &posts)
		if err != nil {
			log.Println(err)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": posts})

	b, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
	}
	err = pc.Cache.Set(ctx, cacheKey, b, 0).Err()
	if err != nil {
		log.Println(err)
	}

}

func (pc *PostController) Kek(ctx *gin.Context) {
	log.Println("kekec")
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": "kek"})
}
