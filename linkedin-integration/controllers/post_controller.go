package controllers

import (
	"linkedin-integration/models"
	"linkedin-integration/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB              *gorm.DB
	LinkedInService services.LinkedInService
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		DB:              db,
		LinkedInService: services.NewLinkedInService(),
	}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post) // Use 201 Created for successful post creation
}

func (pc *PostController) GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := pc.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (pc *PostController) PublishPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := pc.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	err := pc.LinkedInService.PublishPost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Post published"})
}

func (pc *PostController) UnauthorizedHandler(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
}
